package main

import (
	"blockchain-newsfeed-server/config"
	"blockchain-newsfeed-server/middleware"
	"blockchain-newsfeed-server/module/blockchain"
	"blockchain-newsfeed-server/module/core/business"
	"blockchain-newsfeed-server/module/core/repository"
	"blockchain-newsfeed-server/module/core/transport"
	"blockchain-newsfeed-server/module/telebot"
	"blockchain-newsfeed-server/pkg/cache"
	"blockchain-newsfeed-server/pkg/database"
	"blockchain-newsfeed-server/pkg/tracing"
	"blockchain-newsfeed-server/route"
	"blockchain-newsfeed-server/token"
	"context"
	"errors"
	"fmt"
	"net/http"
	"time"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"

	_ "blockchain-newsfeed-server/docs" // Import generated docs

	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/fx"
)

func handleConnection(lc fx.Lifecycle, redisClient cache.IRedisClient) {
	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			log.Info().Ctx(ctx).Msg("Closing connection...")
			if redisClient != nil {
				redisClient.CloseConnection()
			}
			return nil
		},
	})
}

var ConnectionModule = fx.Module(
	"connection",
	fx.Provide(
		database.PostgresqlDatabaseProvider,
		cache.RedisClientProvider,
		func(cnf *config.Config) (*tgbotapi.BotAPI, error) {
			bot, err := tgbotapi.NewBotAPI(cnf.TelegramBot.Token)
			if err != nil {
				log.Error().Err(err).Msg("init telegram bot error")
				return nil, err
			}
			bot.Debug = cnf.TelegramBot.Debug
			return bot, nil
		},
	),
	fx.Invoke(handleConnection),
)

var BusinessModule = fx.Module(
	"business",
	fx.Provide(
		business.NewAuthenticateBiz,
		business.NewMovieBiz,
		business.NewCustomerBiz,
		business.NewPostBiz,
		telebot.NewTelegramClient,
	),
)

var RepositoryModule = fx.Module(
	"repository",
	fx.Provide(repository.NewUserRepository, repository.NewMovieRepository, repository.NewPostRepository),
)

func NewGinEngine(trpt *transport.Transport, jwtMaker token.IJWTMaker) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	engine := gin.New()
	engine.Use(
		middleware.LogRequestInfo(),
	)

	// Register the Swagger handler
	engine.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	route.RegisterHealthCheckRoute(engine)
	route.RegisterRoutes(engine, trpt, jwtMaker)
	return engine
}

func startHttp(lc fx.Lifecycle, cnf *config.Config, engine *gin.Engine, telegramClient *telebot.TelegramClient) {
	server := http.Server{
		Addr:    cnf.AppInfo.ApiPort,
		Handler: engine,
	}

	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			if telegramClient != nil {
				go telegramClient.Handle(ctx)
			}

			go func() {
				log.Info().Ctx(ctx).Msg(fmt.Sprintf("Running API on port %s...", cnf.AppInfo.ApiPort))
				tracing.InitLogger("api-service")

				err := server.ListenAndServe()
				if err != nil && !errors.Is(err, http.ErrServerClosed) {
					log.Error().Ctx(ctx).Err(err).Msg("Run app error")
				}
			}()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Info().Msg("Shutting down server...")
			timeoutCtx, cancel := context.WithTimeout(ctx, 10*time.Second)
			defer cancel()

			if err := server.Shutdown(ctx); err != nil {
				log.Error().Ctx(timeoutCtx).Err(err).Msg("Error shutting down server")
			} else {
				log.Info().Ctx(timeoutCtx).Msg("Server shutdown complete.")
			}
			return nil
		},
	})
}

func main() {
	appFx := fx.New(
		fx.Provide(
			func() *config.Config {
				cnf := config.Init()
				return &cnf
			},
		),
		ConnectionModule,
		fx.Provide(token.NewJWTMaker),
		fx.Provide(blockchain.NewEthClient),
		RepositoryModule,
		BusinessModule,
		fx.Provide(transport.NewTransport),
		fx.Provide(NewGinEngine),
		fx.Invoke(startHttp),
	)
	appFx.Run()
}

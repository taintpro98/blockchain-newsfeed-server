package business

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/module/core/repository"
	"context"

	"github.com/rs/zerolog/log"
)

type IMovieBiz interface {
	ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error)
}

type movieBiz struct {
	movieRepo repository.IMovieRepository
}

func NewMovieBiz(
	movieRepo repository.IMovieRepository,
) IMovieBiz {
	return &movieBiz{
		movieRepo: movieRepo,
	}
}

func (b *movieBiz) ListMovies(ctx context.Context, data dto.ListMoviesRequest) (dto.ListMoviesResponse, *int64, error) {
	log.Info().Ctx(ctx).Interface("data", data).Msg("movieBiz ListMovies")
	moviesDB, err := b.movieRepo.List(ctx, dto.FilterMovie{
		CommonFilter: dto.CommonFilter{
			Select: []string{"id", "title", "content"},
		},
	})
	if err != nil {
		return dto.ListMoviesResponse{}, nil, err
	}
	response := dto.ListMoviesResponse{
		Movies: moviesDB,
	}
	count, err := b.movieRepo.Count(ctx, dto.FilterMovie{})
	if err != nil {
		return response, nil, err
	}
	return response, count, nil
}

package transport

import (
	"blockchain-newsfeed-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz     business.IAuthenticateBiz
	movieBiz    business.IMovieBiz
	customerBiz business.ICustomerBiz
	postBiz     business.IPostBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
	movieBiz business.IMovieBiz,
	customerBiz business.ICustomerBiz,
	postBiz business.IPostBiz,
) *Transport {
	return &Transport{
		authBiz:     authBiz,
		movieBiz:    movieBiz,
		customerBiz: customerBiz,
		postBiz:     postBiz,
	}
}

func HandleHealthCheck(ctx *gin.Context) {
	ctx.JSON(
		http.StatusOK,
		nil,
	)
}

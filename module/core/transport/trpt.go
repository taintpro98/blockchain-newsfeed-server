package transport

import (
	"blockchain-newsfeed-server/module/core/business"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Transport struct {
	authBiz     business.IAuthenticateBiz
	customerBiz business.ICustomerBiz
	postBiz     business.IPostBiz
}

func NewTransport(
	authBiz business.IAuthenticateBiz,
	customerBiz business.ICustomerBiz,
	postBiz business.IPostBiz,
) *Transport {
	return &Transport{
		authBiz:     authBiz,
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

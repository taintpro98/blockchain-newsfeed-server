package transport

import (
	"blockchain-newsfeed-server/module/core/dto"
	"blockchain-newsfeed-server/pkg/constants"

	"github.com/gin-gonic/gin"
)

func (t *Transport) GetCustomerProfile(ctx *gin.Context) {
	userID := ctx.MustGet(constants.XUserID).(string)
	profile, err := t.customerBiz.GetCustomerProfile(ctx, userID)
	dto.HandleResponse(ctx, profile, err)
}

package handler

import (
	"arvanwallet/services"
	"github.com/labstack/echo/v4"
)

type BaseHandler struct {
	Credit Credit
}

type Credit interface {
	HandleIncreaseRequestFromChannel(requestChannel chan string)
	HandleGetBalanceRequest() func(c echo.Context) error
}

func NewBaseHandler(services *services.Services) *BaseHandler {
	return &BaseHandler{
		Credit: newCreditHandler(services),
	}
}

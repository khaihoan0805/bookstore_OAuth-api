package http

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khaihoan0805/bookstore_oauth-api/src/domain/access_token"
)

type AccessTokenHandler interface {
	GetById(*gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{
		service: service,
	}
}

func (handler *accessTokenHandler) GetById(c *gin.Context) {
	// accessTokenId := strings.TrimSpace(c.Param("access_token_id"))

	accessToken, err := handler.service.GetById(c.Param("access_token_id"))
	if err != nil {
		c.JSON(err.Status, err)
	}
	c.JSON(http.StatusOK, accessToken)
}

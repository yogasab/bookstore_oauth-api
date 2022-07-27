package http

import (
	"net/http"

	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(ctx *gin.Context)
}

type accessTokenHandler struct {
	service access_token.Service
}

func NewAccesTokenHandler(service access_token.Service) AccessTokenHandler {
	return &accessTokenHandler{service: service}
}

func (h *accessTokenHandler) GetByID(ctx *gin.Context) {
	ctx.JSON(
		http.StatusNotImplemented,
		gin.H{"code": http.StatusNotImplemented, "status": "not implemented"},
	)
}

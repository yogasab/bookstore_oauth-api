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

	at := ctx.Param("access_token")
	access_token, err := h.service.GetAccessTokenByID(at)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError, "status": "internal server error", "message": err.Error(), "error": err},
		)
		return
	}

	ctx.JSON(
		http.StatusOK,
		gin.H{"code": http.StatusOK, "status": "success", "message": "access token fetched successfully", "data": access_token},
	)
}

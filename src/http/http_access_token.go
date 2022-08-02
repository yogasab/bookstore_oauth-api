package http

import (
	"net/http"

	at "github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

type AccessTokenHandler interface {
	GetByID(ctx *gin.Context)
	CreateAccessToken(ctx *gin.Context)
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

func (h *accessTokenHandler) CreateAccessToken(ctx *gin.Context) {
	var input at.AccessTokenInput

	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			gin.H{"code": http.StatusUnprocessableEntity,
				"status":  "failed",
				"message": "failed to process request",
				"error":   err.Error(),
			},
		)
		return
	}

	result, err := h.service.CreateAccessToken(input)
	if err != nil {
		ctx.JSON(
			http.StatusInternalServerError,
			gin.H{"code": http.StatusInternalServerError,
				"status":  "error",
				"message": "internal status error",
				"error":   err.Error(),
			},
		)
		return
	}

	ctx.JSON(
		http.StatusCreated,
		gin.H{"code": http.StatusCreated,
			"status":  "success",
			"message": "access token created successfully",
			"data":    result,
		},
	)
}

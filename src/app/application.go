package app

import (
	"github.com/bookstore_oauth-api/src/http"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/bookstore_oauth-api/src/repository/rest"
	"github.com/bookstore_oauth-api/src/services/access_token"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	applicationHandler := http.NewAccesTokenHandler(
		access_token.NewService(db.NewDBRepository(), rest.NewRestUserRepository()))

	apiRouter := router.Group("/api/v1/")
	apiRouter.POST("oauth/access-token/", applicationHandler.CreateAccessToken)
	apiRouter.GET("oauth/access-token/:access_token", applicationHandler.GetByID)

	router.Run(":5000")
}

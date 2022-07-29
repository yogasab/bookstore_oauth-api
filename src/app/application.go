package app

import (
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/http"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	applicationHandler := http.NewAccesTokenHandler(access_token.NewService(db.NewDBRepository()))

	apiRouter := router.Group("/api/v1/")
	apiRouter.GET("oauth/access_token/:access_token", applicationHandler.GetByID)

	router.Run(":8000")
}

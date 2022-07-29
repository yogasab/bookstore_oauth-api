package app

import (
	"log"

	"github.com/bookstore_oauth-api/src/clients/cassandra"
	"github.com/bookstore_oauth-api/src/domain/access_token"
	"github.com/bookstore_oauth-api/src/http"
	"github.com/bookstore_oauth-api/src/repository/db"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApplication() {
	session, err := cassandra.GetSession()
	if err != nil {
		panic(err)
	}
	defer session.Close()

	log.Println(session)

	applicationHandler := http.NewAccesTokenHandler(access_token.NewService(db.NewDBRepository()))

	router.GET("/oauth/access_token/:access_token", applicationHandler.GetByID)

	router.Run(":8000")
}

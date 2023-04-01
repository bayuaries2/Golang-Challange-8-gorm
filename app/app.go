package app

import (
	"Challange-7/config"
	"Challange-7/repository"
	"Challange-7/routes"
	"Challange-7/service"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
)

var router = gin.New()

func StartApplication() {
	repo := repository.NewRepo(config.GORM.DB)
	app := service.NewService(repo)
	routes.RegisterApi(router, app)

	port := os.Getenv("APP_PORT")

	router.Run(fmt.Sprintf(":%s", port))
}

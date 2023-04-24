package main

import (
	"log"
	"os"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/pklimuk-eng-thesis/air-conditioning/pkg/domain"
	acHttp "github.com/pklimuk-eng-thesis/air-conditioning/pkg/http"
	acService "github.com/pklimuk-eng-thesis/air-conditioning/pkg/service"
)

func main() {
	ac := domain.AC{Enabled: true, Temperature: 20.0, Humidity: 50.0}

	acService := acService.NewACService(&ac)
	acHandler := acHttp.NewACHandler(acService)

	r := gin.Default()
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "PATCH"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	acHttp.SetupRouter(r, acHandler)

	serviceAddress := os.Getenv("ADDRESS")
	if serviceAddress == "" {
		serviceAddress = ":8086"
	}
	log.Fatal(r.Run(serviceAddress))
}

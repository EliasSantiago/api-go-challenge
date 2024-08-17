package main

import (
	"github.com/EliasSantiago/api-go-challenge/controller"
	"github.com/EliasSantiago/api-go-challenge/db"
	"github.com/EliasSantiago/api-go-challenge/repository"
	"github.com/EliasSantiago/api-go-challenge/usecase"
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	dbConnection, err := db.ConnectDB()
	if err != nil {
		panic(err)
	}

	DriverRepository := repository.NewDriverRepository(dbConnection)
	DriverUseCase := usecase.NewDriverUseCase(DriverRepository)
	DriverController := controller.NewDriverController(DriverUseCase)

	server.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/drivers", DriverController.GetDrivers)
	server.POST("/drivers", DriverController.CreateDriver)
	server.GET("/drivers/:driverId", DriverController.GetDriverByID)

	server.Run(":8081")
}

package main

import (
	"github.com/EliasSantiago/api-go-challenge/controller"
	"github.com/EliasSantiago/api-go-challenge/db"
	"github.com/EliasSantiago/api-go-challenge/repository"
	"github.com/EliasSantiago/api-go-challenge/routes"
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

	VehicleRepository := repository.NewVehicleRepository(dbConnection)
	VehicleUseCase := usecase.NewVehicleUseCase(VehicleRepository, DriverRepository)
	VehicleController := controller.NewVehicleController(VehicleUseCase)

	routes.SetupRouter(server, &DriverController, &VehicleController)

	server.Run(":8081")
}

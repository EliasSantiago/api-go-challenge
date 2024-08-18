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

	VehicleRepository := repository.NewVehicleRepository(dbConnection)
	VehicleUseCase := usecase.NewVehicleUseCase(VehicleRepository, DriverRepository)
	VehicleController := controller.NewVehicleController(VehicleUseCase)

	server.GET("/health", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "ok",
		})
	})

	server.GET("/drivers", DriverController.GetDrivers)
	server.POST("/drivers", DriverController.CreateDriver)
	server.GET("/drivers/:driverId", DriverController.GetDriverByID)
	server.PUT("/drivers", DriverController.UpdateDriver)
	server.DELETE("/drivers/:driverId", DriverController.DeleteDriver)

	server.GET("/vehicles", VehicleController.GetVehicles)
	server.POST("/vehicles", VehicleController.CreateVehicle)
	server.GET("/vehicles/:vehicleId", VehicleController.GetVehicleByID)
	server.PUT("/vehicles", VehicleController.UpdateVehicle)
	server.DELETE("/vehicles/:vehicleId", VehicleController.DeleteVehicle)

	server.POST("/vehicles/:vehicleId/assign-driver/:driverId", VehicleController.AssignDriver)

	server.Run(":8081")
}

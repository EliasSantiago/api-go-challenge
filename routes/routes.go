package routes

import (
	"github.com/EliasSantiago/api-go-challenge/controller"
	"github.com/gin-gonic/gin"
)

func SetupRouter(server *gin.Engine, DriverController *controller.DriverController, VehicleController *controller.VehicleController) {
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
}

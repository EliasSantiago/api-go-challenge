package controller

import (
	"net/http"
	"strconv"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/usecase"
	"github.com/gin-gonic/gin"
)

type VehicleController struct {
	vehicleUsecase usecase.VehicleUsecase
}

func NewVehicleController(vu usecase.VehicleUsecase) VehicleController {
	return VehicleController{
		vehicleUsecase: vu,
	}
}

func (v VehicleController) GetVehicles(ctx *gin.Context) {
	vehicles, err := v.vehicleUsecase.GetVehicles()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, vehicles)
}

func (v VehicleController) CreateVehicle(ctx *gin.Context) {
	var vehicle model.Vehicle
	err := ctx.BindJSON(&vehicle)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newDriver, err := v.vehicleUsecase.CreateVehicle(vehicle)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newDriver)
}

func (v VehicleController) GetVehicleByID(ctx *gin.Context) {
	id := ctx.Param("vehicleId")
	if id == "" {
		response := model.Response{
			Message: "O ID do veículo é obrigatório",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	vehicleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "O ID do veículo deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	vehicle, err := v.vehicleUsecase.GetVehicleByID(vehicleID)
	if err != nil {
		if err.Error() == "Veículo não encontrado" {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusNotFound, response)
		} else {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, response)
		}
		return
	}

	ctx.JSON(http.StatusOK, vehicle)
}

func (v VehicleController) UpdateVehicle(ctx *gin.Context) {
	var vehicleUpdateRequest model.VehicleUpdateRequest
	if err := ctx.ShouldBindJSON(&vehicleUpdateRequest); err != nil {
		response := model.Response{
			Message: "Dados inválidos",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	vehicle, err := v.vehicleUsecase.GetVehicleByID(vehicleUpdateRequest.ID)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao buscar veículo"})
		return
	}

	if vehicle == nil {
		response := model.Response{
			Message: "Veículo não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	vehicleUpdated, err := v.vehicleUsecase.UpdateVehicle(vehicleUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, vehicleUpdated)
}

func (v VehicleController) DeleteVehicle(ctx *gin.Context) {
	id := ctx.Param("vehicleId")
	if id == "" {
		response := model.Response{
			Message: "O ID do veículo é obrigatório",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	vehicleID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "O ID do veículo deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	err = v.vehicleUsecase.DeleteVehicle(vehicleID)
	if err != nil {
		if err.Error() == "Veículo não encontrado" {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusNotFound, response)
		} else {
			response := model.Response{
				Message: err.Error(),
			}
			ctx.JSON(http.StatusBadRequest, response)
		}
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Veículo deletado com sucesso"})
}

func (v VehicleController) AssignDriver(ctx *gin.Context) {
	vId := ctx.Param("vehicleId")
	dId := ctx.Param("driverId")

	vehicleID, err := strconv.ParseInt(vId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do veículo inválido"})
		return
	}

	driverID, err := strconv.ParseInt(dId, 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, model.Response{Message: "ID do motorista inválido"})
		return
	}

	err = v.vehicleUsecase.AssignDriver(vehicleID, driverID)
	if err != nil {
		if err.Error() == "veículo não encontrado" {
			ctx.JSON(http.StatusNotFound, model.Response{Message: "Veículo não encontrado"})
		} else if err.Error() == "motorista não encontrado" {
			ctx.JSON(http.StatusNotFound, model.Response{Message: "Motorista não encontrado"})
		} else {
			ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao vincular motorista ao veículo"})
		}
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Motorista atribuído ao veículo com sucesso"})
}

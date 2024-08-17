package controller

import (
	"net/http"
	"strconv"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/usecase"
	"github.com/gin-gonic/gin"
)

type DriverController struct {
	driverUsecase usecase.DriverUsecase
}

func NewDriverController(usecase usecase.DriverUsecase) DriverController {
	return DriverController{
		driverUsecase: usecase,
	}
}

func (d DriverController) GetDrivers(ctx *gin.Context) {
	drivers, err := d.driverUsecase.GetDrivers()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(http.StatusOK, drivers)
}

func (d DriverController) CreateDriver(ctx *gin.Context) {
	var driver model.Driver
	err := ctx.BindJSON(&driver)

	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	newDriver, err := d.driverUsecase.CreateDriver(driver)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusCreated, newDriver)
}

func (d DriverController) GetDriverByID(ctx *gin.Context) {
	id := ctx.Param("driverId")
	if id == "" {
		response := model.Response{
			Message: "O ID do motorista é obrigatório",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	driverID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "O ID do motorista deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	driver, err := d.driverUsecase.GetDriverByID(driverID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if driver == nil {
		response := model.Response{
			Message: "Motorista não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, driver)
}

func (d DriverController) UpdateDriver(ctx *gin.Context) {
	var driverUpdateRequest model.DriverUpdateRequest
	if err := ctx.ShouldBindJSON(&driverUpdateRequest); err != nil {
		response := model.Response{
			Message: "Dados inválidos",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	driver, err := d.driverUsecase.UpdateDriver(driverUpdateRequest)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	if driver == nil {
		response := model.Response{
			Message: "Motorista não encontrado",
		}
		ctx.JSON(http.StatusNotFound, response)
		return
	}

	ctx.JSON(http.StatusOK, driver)
}

func (d DriverController) DeleteDriver(ctx *gin.Context) {
	id := ctx.Param("driverId")
	if id == "" {
		response := model.Response{
			Message: "O ID do motorista é obrigatório",
		}

		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	driverID, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		response := model.Response{
			Message: "O ID do motorista deve ser um número",
		}
		ctx.JSON(http.StatusBadRequest, response)
		return
	}

	driver, err := d.driverUsecase.GetDriverByID(driverID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, model.Response{Message: "Erro ao buscar motorista"})
		return
	}

	if driver == nil {
		ctx.JSON(http.StatusNotFound, model.Response{Message: "Motorista não encontrado"})
		return
	}

	err = d.driverUsecase.DeleteDriver(driverID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, model.Response{Message: "Motorista deletado com sucesso"})
}

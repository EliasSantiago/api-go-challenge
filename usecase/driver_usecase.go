package usecase

import (
	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository"
)

type DriverUsecase struct {
	repository repository.DriverRepository
}

func NewDriverUseCase(repo repository.DriverRepository) DriverUsecase {
	return DriverUsecase{repository: repo}
}

func (du *DriverUsecase) GetDrivers() ([]model.Driver, error) {
	return du.repository.GetDrivers()
}

func (du *DriverUsecase) CreateDriver(driver model.Driver) (model.Driver, error) {
	driverID, err := du.repository.CreateDriver(driver)
	if err != nil {
		return model.Driver{}, err
	}

	driver.ID = int64(driverID)

	return driver, nil
}

func (du *DriverUsecase) GetDriverByID(id int64) (*model.Driver, error) {
	driver, err := du.repository.GetDriverByID(id)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (du *DriverUsecase) UpdateDriver(request model.DriverUpdateRequest) (*model.Driver, error) {
	driver := model.Driver{
		ID:   int64(request.ID),
		CPF:  request.CPF,
		Name: request.Name,
	}

	err := du.repository.UpdateDriver(driver)
	if err != nil {
		return nil, err
	}

	return &driver, nil
}

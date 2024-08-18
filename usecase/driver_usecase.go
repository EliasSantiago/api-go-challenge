package usecase

import (
	"errors"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository"
)

type DriverUsecase interface {
	GetDrivers() ([]model.Driver, error)
	CreateDriver(driver model.Driver) (model.Driver, error)
	GetDriverByID(id int64) (*model.Driver, error)
	UpdateDriver(request model.DriverUpdateRequest) (*model.Driver, error)
	DeleteDriver(id int64) error
}

type driverUsecase struct {
	driverRepo repository.DriverRepository
}

func NewDriverUseCase(repo repository.DriverRepository) DriverUsecase {
	return &driverUsecase{
		driverRepo: repo,
	}
}

func (du *driverUsecase) GetDrivers() ([]model.Driver, error) {
	return du.driverRepo.GetDrivers()
}

func (du *driverUsecase) CreateDriver(driver model.Driver) (model.Driver, error) {
	driverID, err := du.driverRepo.CreateDriver(driver)
	if err != nil {
		return model.Driver{}, err
	}

	driver.ID = int64(driverID)

	return driver, nil
}

func (du *driverUsecase) GetDriverByID(id int64) (*model.Driver, error) {
	if id <= 0 {
		return nil, errors.New("O ID do motorista deve ser um número positivo")
	}

	driver, err := du.driverRepo.GetDriverByID(id)
	if err != nil {
		return nil, err
	}

	if driver == nil {
		return nil, errors.New("Motorista não encontrado")
	}

	return driver, nil
}

func (du *driverUsecase) UpdateDriver(request model.DriverUpdateRequest) (*model.Driver, error) {
	if request.ID <= 0 {
		return nil, errors.New("O ID do motorista deve ser um número positivo")
	}

	driver, err := du.driverRepo.GetDriverByID(request.ID)
	if err != nil {
		return nil, err
	}

	if driver == nil {
		return nil, errors.New("Motorista não encontrado")
	}

	driver.Name = request.Name
	driver.CPF = request.CPF

	err = du.driverRepo.UpdateDriver(*driver)
	if err != nil {
		return nil, err
	}

	return driver, nil
}

func (du *driverUsecase) DeleteDriver(id int64) error {
	if id <= 0 {
		return errors.New("O ID do motorista deve ser um número positivo")
	}

	driver, err := du.driverRepo.GetDriverByID(id)
	if err != nil {
		return err
	}

	if driver == nil {
		return errors.New("Motorista não encontrado")
	}

	err = du.driverRepo.DeleteDriver(id)
	if err != nil {
		return err
	}

	return nil
}

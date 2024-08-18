package usecase

import (
	"errors"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository"
)

type VehicleUsecase interface {
	GetVehicles() ([]model.Vehicle, error)
	CreateVehicle(vehicle model.Vehicle) (model.Vehicle, error)
	AssignDriver(vehicleID, driverID int64) error
	GetVehicleByID(id int64) (*model.Vehicle, error)
	UpdateVehicle(request model.VehicleUpdateRequest) (*model.Vehicle, error)
	DeleteVehicle(id int64) error
}

type vehicleUsecase struct {
	vehicleRepo repository.VehicleRepository
	driverRepo  repository.DriverRepository
}

func NewVehicleUseCase(vr repository.VehicleRepository, dr repository.DriverRepository) VehicleUsecase {
	return &vehicleUsecase{
		vehicleRepo: vr,
		driverRepo:  dr,
	}
}

func (vu *vehicleUsecase) GetVehicles() ([]model.Vehicle, error) {
	return vu.vehicleRepo.GetVehicles()
}

func (vu *vehicleUsecase) CreateVehicle(vehicle model.Vehicle) (model.Vehicle, error) {
	vehicleID, err := vu.vehicleRepo.CreateVehicle(vehicle)
	if err != nil {
		return model.Vehicle{}, err
	}

	vehicle.ID = int64(vehicleID)

	return vehicle, nil
}

func (vu *vehicleUsecase) GetVehicleByID(id int64) (*model.Vehicle, error) {
	if id <= 0 {
		return nil, errors.New("o ID do veículo deve ser um número positivo")
	}

	vehicle, err := vu.vehicleRepo.GetVehicleByID(id)
	if err != nil {
		return nil, err
	}

	if vehicle == nil {
		return nil, errors.New("veículo não encontrado")
	}

	return vehicle, nil
}

func (vu *vehicleUsecase) UpdateVehicle(request model.VehicleUpdateRequest) (*model.Vehicle, error) {
	if request.ID <= 0 {
		return nil, errors.New("o ID do veículo deve ser um número positivo")
	}

	vehicle, err := vu.vehicleRepo.GetVehicleByID(request.ID)
	if err != nil {
		return nil, err
	}

	if vehicle == nil {
		return nil, errors.New("veículo não encontrado")
	}

	data := model.Vehicle(request)

	err = vu.vehicleRepo.UpdateVehicle(data)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (vu *vehicleUsecase) DeleteVehicle(id int64) error {
	if id <= 0 {
		return errors.New("o ID do veículo deve ser um número positivo")
	}

	vehicle, err := vu.vehicleRepo.GetVehicleByID(id)
	if err != nil {
		return err
	}

	if vehicle == nil {
		return errors.New("veículo não encontrado")
	}

	err = vu.vehicleRepo.DeleteVehicle(id)
	if err != nil {
		return err
	}

	return nil
}

func (vu *vehicleUsecase) AssignDriver(vehicleID, driverID int64) error {
	vehicle, err := vu.vehicleRepo.GetVehicleByID(vehicleID)
	if err != nil {
		return err
	}
	if vehicle == nil {
		return errors.New("veículo não encontrado")
	}

	driver, err := vu.driverRepo.GetDriverByID(driverID)
	if err != nil {
		return err
	}
	if driver == nil {
		return errors.New("motorista não encontrado")
	}

	data := model.VehicleAssignDriverRequest{
		VehicleID: vehicleID,
		DriverID:  driverID,
	}

	return vu.vehicleRepo.AssignDriver(data)
}

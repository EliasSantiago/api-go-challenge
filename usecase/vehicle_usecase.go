package usecase

import (
	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository"
)

type VehicleUsecase struct {
	repository repository.VehicleRepository
}

func NewVehicleUseCase(repo repository.VehicleRepository) VehicleUsecase {
	return VehicleUsecase{repository: repo}
}

func (vu *VehicleUsecase) GetVehicles() ([]model.Vehicle, error) {
	return vu.repository.GetVehicles()
}

func (vu *VehicleUsecase) CreateVehicle(vehicle model.Vehicle) (model.Vehicle, error) {
	vehicleID, err := vu.repository.CreateVehicle(vehicle)
	if err != nil {
		return model.Vehicle{}, err
	}

	vehicle.ID = int64(vehicleID)

	return vehicle, nil
}

func (du *VehicleUsecase) GetVehicleByID(id int64) (*model.Vehicle, error) {
	vehicle, err := du.repository.GetVehicleByID(id)
	if err != nil {
		return nil, err
	}

	return vehicle, nil
}

func (du *VehicleUsecase) UpdateVehicle(request model.VehicleUpdateRequest) (*model.Vehicle, error) {
	vehicle := model.Vehicle{
		ID:             int64(request.ID),
		LicenseVehicle: request.LicenseVehicle,
		Model:          request.Model,
	}

	err := du.repository.UpdateVehicle(vehicle)
	if err != nil {
		return nil, err
	}

	return &vehicle, nil
}

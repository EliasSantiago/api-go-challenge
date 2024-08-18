package usecase

import (
	"testing"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetVehicles(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	mockVehicles := []model.Vehicle{
		{ID: 1, Model: "Veículo 1", LicensePlate: "ABC-1234"},
		{ID: 2, Model: "Veículo 2", LicensePlate: "DEF-5678"},
	}

	mockVehicleRepo.EXPECT().GetVehicles().Return(mockVehicles, nil)

	vehicles, err := mockUsecase.GetVehicles()
	assert.NoError(t, err)
	assert.Equal(t, mockVehicles, vehicles)
}

func TestCreateVehicle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	mockVehicle := model.Vehicle{Model: "Novo veículo"}
	expectedVehicle := model.Vehicle{ID: 1, Model: "Novo veículo"}
	mockVehicleRepo.EXPECT().CreateVehicle(mockVehicle).Return(int64(1), nil)

	createdVehicle, err := mockUsecase.CreateVehicle(mockVehicle)
	assert.NoError(t, err)
	assert.Equal(t, expectedVehicle, createdVehicle)
}

func TestAssignDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	vehicleID := int64(1)
	driverID := int64(1)

	mockVehicle := &model.Vehicle{ID: vehicleID, Model: "Veículo 1"}
	mockDriver := &model.Driver{ID: driverID, Name: "Motorista 1"}

	mockVehicleRepo.EXPECT().GetVehicleByID(vehicleID).Return(mockVehicle, nil)
	mockDriverRepo.EXPECT().GetDriverByID(driverID).Return(mockDriver, nil)
	mockVehicleRepo.EXPECT().AssignDriver(model.VehicleAssignDriverRequest{
		VehicleID: vehicleID,
		DriverID:  driverID,
	}).Return(nil)

	err := mockUsecase.AssignDriver(vehicleID, driverID)
	assert.NoError(t, err)
}

func TestGetVehicleByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	mockVehicle := model.Vehicle{ID: 1, Model: "Veículo 1"}
	mockVehicleRepo.EXPECT().GetVehicleByID(int64(1)).Return(&mockVehicle, nil)

	vehicle, err := mockUsecase.GetVehicleByID(1)
	assert.NoError(t, err)
	assert.Equal(t, &mockVehicle, vehicle)
}

func TestUpdateVehicle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	updateRequest := model.VehicleUpdateRequest{ID: 1, Model: "Veículo atualizado"}
	existingVehicle := model.Vehicle{ID: 1, Model: "Veículo antigo"}
	updatedVehicle := model.Vehicle{ID: 1, Model: "Veículo atualizado"}

	mockVehicleRepo.EXPECT().GetVehicleByID(updateRequest.ID).Return(&existingVehicle, nil)
	mockVehicleRepo.EXPECT().UpdateVehicle(updatedVehicle).Return(nil)

	result, err := mockUsecase.UpdateVehicle(updateRequest)
	assert.NoError(t, err)
	assert.Equal(t, &updatedVehicle, result)
}

func TestDeleteVehicle(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockVehicleRepo := mocks.NewMockVehicleRepository(ctrl)
	mockDriverRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewVehicleUseCase(mockVehicleRepo, mockDriverRepo)

	err := mockUsecase.DeleteVehicle(0)
	assert.Error(t, err)
	assert.Equal(t, "o ID do veículo deve ser um número positivo", err.Error())

	mockVehicleRepo.EXPECT().GetVehicleByID(int64(1)).Return(nil, nil)
	err = mockUsecase.DeleteVehicle(1)
	assert.Error(t, err)
	assert.Equal(t, "veículo não encontrado", err.Error())

	mockVehicle := model.Vehicle{ID: 1, Model: "Veículo 1"}
	mockVehicleRepo.EXPECT().GetVehicleByID(int64(1)).Return(&mockVehicle, nil)
	mockVehicleRepo.EXPECT().DeleteVehicle(int64(1)).Return(nil)

	err = mockUsecase.DeleteVehicle(1)
	assert.NoError(t, err)
}

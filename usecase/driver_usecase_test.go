package usecase

import (
	"testing"

	"github.com/EliasSantiago/api-go-challenge/model"
	"github.com/EliasSantiago/api-go-challenge/repository/mocks"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestGetDrivers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewDriverUseCase(mockRepo)

	mockDrivers := []model.Driver{
		{ID: 1, Name: "Motorista 1"},
		{ID: 2, Name: "Motorista 2"},
	}

	mockRepo.EXPECT().GetDrivers().Return(mockDrivers, nil)

	drivers, err := mockUsecase.GetDrivers()
	assert.NoError(t, err)
	assert.Equal(t, mockDrivers, drivers)
}

func TestCreateDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewDriverUseCase(mockRepo)

	mockDriver := model.Driver{CPF: "12345678901", Name: "Novo Motorista"}
	expectedDriver := model.Driver{ID: 1, CPF: "12345678901", Name: "Novo Motorista"}
	mockRepo.EXPECT().CreateDriver(mockDriver).Return(int64(1), nil)

	createdDriver, err := mockUsecase.CreateDriver(mockDriver)
	assert.NoError(t, err)
	assert.Equal(t, expectedDriver, createdDriver)
}

func TestGetDriverByID(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewDriverUseCase(mockRepo)

	mockDriver := model.Driver{ID: 1, Name: "Motorista 1"}
	mockRepo.EXPECT().GetDriverByID(int64(1)).Return(&mockDriver, nil)

	driver, err := mockUsecase.GetDriverByID(1)
	assert.NoError(t, err)
	assert.Equal(t, &mockDriver, driver)
}

func TestUpdateDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewDriverUseCase(mockRepo)

	updateRequest := model.DriverUpdateRequest{ID: 1, Name: "Motorista atualizado", CPF: "98765432100"}
	existingDriver := model.Driver{ID: 1, Name: "Motorista antigo", CPF: "12345678901"}
	updatedDriver := model.Driver{ID: 1, Name: "Motorista atualizado", CPF: "98765432100"}

	mockRepo.EXPECT().GetDriverByID(updateRequest.ID).Return(&existingDriver, nil)
	mockRepo.EXPECT().UpdateDriver(updatedDriver).Return(nil)

	result, err := mockUsecase.UpdateDriver(updateRequest)
	assert.NoError(t, err)
	assert.Equal(t, &updatedDriver, result)
}

func TestDeleteDriver(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockRepo := mocks.NewMockDriverRepository(ctrl)
	mockUsecase := NewDriverUseCase(mockRepo)

	err := mockUsecase.DeleteDriver(0)
	assert.Error(t, err)
	assert.Equal(t, "o ID do motorista deve ser um número positivo", err.Error())

	mockRepo.EXPECT().GetDriverByID(int64(1)).Return(nil, nil)
	err = mockUsecase.DeleteDriver(1)
	assert.Error(t, err)
	assert.Equal(t, "motorista não encontrado", err.Error())

	mockDriver := model.Driver{ID: 1, Name: "Driver 1"}
	mockRepo.EXPECT().GetDriverByID(int64(1)).Return(&mockDriver, nil)
	mockRepo.EXPECT().DeleteDriver(int64(1)).Return(nil)

	err = mockUsecase.DeleteDriver(1)
	assert.NoError(t, err)
}

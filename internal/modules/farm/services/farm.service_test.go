package service_test

import (
	"errors"
	"testing"

	"go-farms/internal/entity"
	service "go-farms/internal/modules/farm/services"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFarmRepository struct {
	mock.Mock
}

func (m *MockFarmRepository) Create(farm *entity.Farm) (*entity.Farm, error) {
	args := m.Called(farm)
	if args.Get(0) != nil {
		return args.Get(0).(*entity.Farm), args.Error(1)
	}
	return nil, args.Error(1)
}

func (m *MockFarmRepository) List() []entity.Farm {
	args := m.Called()
	if args.Get(0) != nil {
		return args.Get(0).([]entity.Farm)
	}
	return []entity.Farm{}
}

func (m *MockFarmRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestFarmService_Create(t *testing.T) {
	mockRepo := new(MockFarmRepository)
	farmService := service.GetFarmService(mockRepo)

	t.Run("Create farm successfully", func(t *testing.T) {
		newFarm := &entity.Farm{ID: 1, FarmName: "New Farm", LandArea: 100, UnitOfMeasure: "acre", Address: "123 Main St"}
		mockRepo.On("Create", newFarm).Return(newFarm, nil)

		createdFarm, err := farmService.Create(newFarm)

		assert.NoError(t, err)
		assert.Equal(t, newFarm, createdFarm)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Create farm with error", func(t *testing.T) {
		newFarm := &entity.Farm{ID: 2, FarmName: "New Farm", LandArea: 100, UnitOfMeasure: "acre", Address: "123 Main St"}
		expectedError := errors.New("failed to create farm")
		mockRepo.On("Create", newFarm).Return((*entity.Farm)(nil), expectedError)

		createdFarm, err := farmService.Create(newFarm)

		assert.Nil(t, createdFarm)
		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
}

func TestFarmService_List(t *testing.T) {
	mockRepo := new(MockFarmRepository)
	farmService := service.GetFarmService(mockRepo)

	t.Run("List farms successfully", func(t *testing.T) {
		expectedFarms := []entity.Farm{
			{ID: 1, FarmName: "Farm 1"},
			{ID: 2, FarmName: "Farm 2"},
		}
		mockRepo.On("List").Return(expectedFarms)

		farms := farmService.List()

		assert.Equal(t, expectedFarms, farms)
		mockRepo.AssertExpectations(t)
	})
}

func TestFarmService_Delete(t *testing.T) {
	mockRepo := new(MockFarmRepository)
	farmService := service.GetFarmService(mockRepo)

	t.Run("Delete farm successfully", func(t *testing.T) {
		farmID := 1
		mockRepo.On("Delete", farmID).Return(nil)

		err := farmService.Delete(farmID)

		assert.NoError(t, err)
		mockRepo.AssertExpectations(t)
	})

	t.Run("Delete farm with error", func(t *testing.T) {
		farmID := 2
		expectedError := errors.New("failed to delete farm")
		mockRepo.On("Delete", farmID).Return(expectedError)

		err := farmService.Delete(farmID)

		assert.EqualError(t, err, expectedError.Error())
		mockRepo.AssertExpectations(t)
	})
}

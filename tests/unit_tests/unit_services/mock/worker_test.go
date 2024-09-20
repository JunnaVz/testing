package mock

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock service
type MockWorkerService struct {
	mock.Mock
}

func (m *MockWorkerService) Create(worker *models.Worker) (*models.Worker, error) {
	args := m.Called(worker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerService) Update(worker *models.Worker) (*models.Worker, error) {
	args := m.Called(worker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerService) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockWorkerService) GetWorkerByID(id uuid.UUID) (*models.Worker, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerService) GetAllWorkers() ([]models.Worker, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Worker), args.Error(1)
}

func (m *MockWorkerService) GetWorkerByEmail(email string) (*models.Worker, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func TestCreateWorker_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	mockService.On("Create", worker).Return(worker, nil)

	createdWorker, err := mockService.Create(worker)

	assert.NoError(t, err)
	assert.Equal(t, worker, createdWorker)
	mockService.AssertExpectations(t)
}

func TestCreateWorker_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	mockService.On("Create", worker).Return((*models.Worker)(nil), errors.New("creation failed"))

	createdWorker, err := mockService.Create(worker)

	assert.Error(t, err)
	assert.Nil(t, createdWorker)
	mockService.AssertExpectations(t)
}

func TestUpdateWorker_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	mockService.On("Update", worker).Return(worker, nil)

	updatedWorker, err := mockService.Update(worker)

	assert.NoError(t, err)
	assert.Equal(t, worker, updatedWorker)
	mockService.AssertExpectations(t)
}

func TestUpdateWorker_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	mockService.On("Update", worker).Return((*models.Worker)(nil), errors.New("update failed"))

	updatedWorker, err := mockService.Update(worker)

	assert.Error(t, err)
	assert.Nil(t, updatedWorker)
	mockService.AssertExpectations(t)
}

func TestDeleteWorker_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	workerID := uuid.New()
	mockService.On("Delete", workerID).Return(nil)

	err := mockService.Delete(workerID)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestDeleteWorker_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	workerID := uuid.New()
	mockService.On("Delete", workerID).Return(errors.New("deletion failed"))

	err := mockService.Delete(workerID)

	assert.Error(t, err)
	mockService.AssertExpectations(t)
}

func TestGetWorkerByID_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	workerID := uuid.New()
	mockService.On("GetWorkerByID", workerID).Return(worker, nil)

	receivedWorker, err := mockService.GetWorkerByID(workerID)

	assert.NoError(t, err)
	assert.Equal(t, worker, receivedWorker)
	mockService.AssertExpectations(t)
}

func TestGetWorkerByID_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	workerID := uuid.New()
	mockService.On("GetWorkerByID", workerID).Return((*models.Worker)(nil), errors.New("worker not found"))

	receivedWorker, err := mockService.GetWorkerByID(workerID)

	assert.Error(t, err)
	assert.Nil(t, receivedWorker)
	mockService.AssertExpectations(t)
}

func TestGetAllWorkers_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	workers := []models.Worker{
		{Name: "WorkerName1", Email: "worker1@gmail.com"},
		{Name: "WorkerName2", Email: "worker2@gmail.com"},
	}
	mockService.On("GetAllWorkers").Return(workers, nil)

	receivedWorkers, err := mockService.GetAllWorkers()

	assert.NoError(t, err)
	assert.Equal(t, workers, receivedWorkers)
	mockService.AssertExpectations(t)
}

func TestGetAllWorkers_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	mockService.On("GetAllWorkers").Return(([]models.Worker)(nil), errors.New("workers not found"))

	receivedWorkers, err := mockService.GetAllWorkers()

	assert.Error(t, err)
	assert.Nil(t, receivedWorkers)
	mockService.AssertExpectations(t)
}

func TestGetWorkerByEmail_Success(t *testing.T) {
	mockService := new(MockWorkerService)
	worker := &models.Worker{Name: "WorkerName", Email: "worker@gmail.com"}
	email := worker.Email
	mockService.On("GetWorkerByEmail", email).Return(worker, nil)

	receivedWorker, err := mockService.GetWorkerByEmail(email)

	assert.NoError(t, err)
	assert.Equal(t, worker, receivedWorker)
	mockService.AssertExpectations(t)
}

func TestGetWorkerByEmail_Failure(t *testing.T) {
	mockService := new(MockWorkerService)
	email := "worker@gmail.com"
	mockService.On("GetWorkerByEmail", email).Return((*models.Worker)(nil), errors.New("worker not found"))

	receivedWorker, err := mockService.GetWorkerByEmail(email)

	assert.Error(t, err)
	assert.Nil(t, receivedWorker)
	mockService.AssertExpectations(t)
}

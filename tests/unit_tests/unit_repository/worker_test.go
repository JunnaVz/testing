package unit_repository

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock repository
type MockWorkerRepository struct {
	mock.Mock
}

func (m *MockWorkerRepository) Create(worker *models.Worker) (*models.Worker, error) {
	args := m.Called(worker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockWorkerRepository) Update(worker *models.Worker) (*models.Worker, error) {
	args := m.Called(worker)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerRepository) GetWorkerByID(id uuid.UUID) (*models.Worker, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func (m *MockWorkerRepository) GetAllWorkers() ([]models.Worker, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Worker), args.Error(1)
}

func (m *MockWorkerRepository) GetWorkerByEmail(email string) (*models.Worker, error) {
	args := m.Called(email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Worker), args.Error(1)
}

func TestCreateWorker_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	mockRepo.On("Create", worker).Return(worker, nil)

	createdWorker, err := mockRepo.Create(worker)

	assert.NoError(t, err)
	assert.Equal(t, worker, createdWorker)
	mockRepo.AssertExpectations(t)
}

func TestCreateWorker_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	mockRepo.On("Create", worker).Return((*models.Worker)(nil), errors.New("creation failed"))

	createdWorker, err := mockRepo.Create(worker)

	assert.Error(t, err)
	assert.Nil(t, createdWorker)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWorker_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	workerID := uuid.New()
	mockRepo.On("Delete", workerID).Return(nil)

	err := mockRepo.Delete(workerID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteWorker_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	workerID := uuid.New()
	mockRepo.On("Delete", workerID).Return(errors.New("deletion failed"))

	err := mockRepo.Delete(workerID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWorker_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	mockRepo.On("Update", worker).Return(worker, nil)

	updatedWorker, err := mockRepo.Update(worker)

	assert.NoError(t, err)
	assert.Equal(t, worker, updatedWorker)
	mockRepo.AssertExpectations(t)
}

func TestUpdateWorker_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	mockRepo.On("Update", worker).Return((*models.Worker)(nil), errors.New("update failed"))

	updatedWorker, err := mockRepo.Update(worker)

	assert.Error(t, err)
	assert.Nil(t, updatedWorker)
	mockRepo.AssertExpectations(t)
}

func TestGetWorkerByID_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	workerID := uuid.New()
	mockRepo.On("GetWorkerByID", workerID).Return(worker, nil)

	receivedWorker, err := mockRepo.GetWorkerByID(workerID)

	assert.NoError(t, err)
	assert.Equal(t, worker, receivedWorker)
	mockRepo.AssertExpectations(t)
}

func TestGetWorkerByID_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	workerID := uuid.New()
	mockRepo.On("GetWorkerByID", workerID).Return((*models.Worker)(nil), errors.New("worker not found"))

	receivedWorker, err := mockRepo.GetWorkerByID(workerID)

	assert.Error(t, err)
	assert.Nil(t, receivedWorker)
	mockRepo.AssertExpectations(t)
}

func TestGetAllWorkers_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	workers := []models.Worker{
		{Name: "First Name 1", Surname: "Last Name 1", Email: "email1@example.com"},
		{Name: "First Name 2", Surname: "Last Name 2", Email: "email2@example.com"},
	}
	mockRepo.On("GetAllWorkers").Return(workers, nil)

	receivedWorkers, err := mockRepo.GetAllWorkers()

	assert.NoError(t, err)
	assert.Equal(t, workers, receivedWorkers)
	mockRepo.AssertExpectations(t)
}

func TestGetAllWorkers_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	mockRepo.On("GetAllWorkers").Return(([]models.Worker)(nil), errors.New("workers not found"))

	receivedWorkers, err := mockRepo.GetAllWorkers()

	assert.Error(t, err)
	assert.Nil(t, receivedWorkers)
	mockRepo.AssertExpectations(t)
}

func TestGetWorkerByEmail_Success(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	worker := &models.Worker{Name: "First Name", Surname: "Last Name", Email: "email@example.com"}
	mockRepo.On("GetWorkerByEmail", "email@example.com").Return(worker, nil)

	receivedWorker, err := mockRepo.GetWorkerByEmail("email@example.com")

	assert.NoError(t, err)
	assert.Equal(t, worker, receivedWorker)
	mockRepo.AssertExpectations(t)
}

func TestGetWorkerByEmail_Failure(t *testing.T) {
	mockRepo := new(MockWorkerRepository)
	mockRepo.On("GetWorkerByEmail", "email@example.com").Return((*models.Worker)(nil), errors.New("worker not found"))

	receivedWorker, err := mockRepo.GetWorkerByEmail("email@example.com")

	assert.Error(t, err)
	assert.Nil(t, receivedWorker)
	mockRepo.AssertExpectations(t)
}

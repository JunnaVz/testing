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
type MockTaskRepository struct {
	mock.Mock
}

func (m *MockTaskRepository) Create(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskRepository) Update(task *models.Task) (*models.Task, error) {
	args := m.Called(task)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByID(id uuid.UUID) (*models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTasksInCategory(category int) ([]models.Task, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskRepository) GetTaskByName(name string) (*models.Task, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func TestCreateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockRepo.On("Create", task).Return(task, nil)

	createdTask, err := mockRepo.Create(task)

	assert.NoError(t, err)
	assert.Equal(t, task, createdTask)
	mockRepo.AssertExpectations(t)
}

func TestCreateTask_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockRepo.On("Create", task).Return((*models.Task)(nil), errors.New("creation failed"))

	createdTask, err := mockRepo.Create(task)

	assert.Error(t, err)
	assert.Nil(t, createdTask)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskID := uuid.New()
	mockRepo.On("Delete", taskID).Return(nil)

	err := mockRepo.Delete(taskID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteTask_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskID := uuid.New()
	mockRepo.On("Delete", taskID).Return(errors.New("deletion failed"))

	err := mockRepo.Delete(taskID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockRepo.On("Update", task).Return(task, nil)

	updatedTask, err := mockRepo.Update(task)

	assert.NoError(t, err)
	assert.Equal(t, task, updatedTask)
	mockRepo.AssertExpectations(t)
}

func TestUpdateTask_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockRepo.On("Update", task).Return((*models.Task)(nil), errors.New("update failed"))

	updatedTask, err := mockRepo.Update(task)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	taskID := uuid.New()
	mockRepo.On("GetTaskByID", taskID).Return(task, nil)

	receivedTask, err := mockRepo.GetTaskByID(taskID)

	assert.NoError(t, err)
	assert.Equal(t, task, receivedTask)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByID_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	taskID := uuid.New()
	mockRepo.On("GetTaskByID", taskID).Return((*models.Task)(nil), errors.New("task not found"))

	receivedTask, err := mockRepo.GetTaskByID(taskID)

	assert.Error(t, err)
	assert.Nil(t, receivedTask)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	tasks := []models.Task{
		{Name: "TaskName1", PricePerSingle: 100.0, Category: 1},
		{Name: "TaskName2", PricePerSingle: 200.0, Category: 2},
	}
	mockRepo.On("GetAllTasks").Return(tasks, nil)

	receivedTasks, err := mockRepo.GetAllTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, receivedTasks)
	mockRepo.AssertExpectations(t)
}

func TestGetAllTasks_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockRepo.On("GetAllTasks").Return(([]models.Task)(nil), errors.New("tasks not found"))

	receivedTasks, err := mockRepo.GetAllTasks()

	assert.Error(t, err)
	assert.Nil(t, receivedTasks)
	mockRepo.AssertExpectations(t)
}

func TestGetTasksInCategory_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	tasks := []models.Task{
		{Name: "TaskName1", PricePerSingle: 100.0, Category: 1},
		{Name: "TaskName2", PricePerSingle: 200.0, Category: 1},
	}
	mockRepo.On("GetTasksInCategory", 1).Return(tasks, nil)

	receivedTasks, err := mockRepo.GetTasksInCategory(1)

	assert.NoError(t, err)
	assert.Equal(t, tasks, receivedTasks)
	mockRepo.AssertExpectations(t)
}

func TestGetTasksInCategory_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockRepo.On("GetTasksInCategory", 1).Return(([]models.Task)(nil), errors.New("tasks not found"))

	receivedTasks, err := mockRepo.GetTasksInCategory(1)

	assert.Error(t, err)
	assert.Nil(t, receivedTasks)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByName_Success(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockRepo.On("GetTaskByName", "TaskName").Return(task, nil)

	receivedTask, err := mockRepo.GetTaskByName("TaskName")

	assert.NoError(t, err)
	assert.Equal(t, task, receivedTask)
	mockRepo.AssertExpectations(t)
}

func TestGetTaskByName_Failure(t *testing.T) {
	mockRepo := new(MockTaskRepository)
	mockRepo.On("GetTaskByName", "TaskName").Return((*models.Task)(nil), errors.New("task not found"))

	receivedTask, err := mockRepo.GetTaskByName("TaskName")

	assert.Error(t, err)
	assert.Nil(t, receivedTask)
	mockRepo.AssertExpectations(t)
}

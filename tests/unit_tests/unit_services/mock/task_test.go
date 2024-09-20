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
type MockTaskService struct {
	mock.Mock
}

func (m *MockTaskService) Create(name string, price float64, category int) (*models.Task, error) {
	args := m.Called(name, price, category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockTaskService) Update(id uuid.UUID, category int, name string, price float64) (*models.Task, error) {
	args := m.Called(id, category, name, price)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskByID(id uuid.UUID) (*models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func (m *MockTaskService) GetAllTasks() ([]models.Task, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskService) GetTasksInCategory(category int) ([]models.Task, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockTaskService) GetTaskByName(name string) (*models.Task, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Task), args.Error(1)
}

func TestCreateTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockService.On("Create", "TaskName", 100.0, 1).Return(task, nil)

	createdTask, err := mockService.Create("TaskName", 100.0, 1)

	assert.NoError(t, err)
	assert.Equal(t, task, createdTask)
	mockService.AssertExpectations(t)
}

func TestCreateTask_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("Create", "TaskName", 100.0, 1).Return((*models.Task)(nil), errors.New("creation failed"))

	createdTask, err := mockService.Create("TaskName", 100.0, 1)

	assert.Error(t, err)
	assert.Nil(t, createdTask)
	mockService.AssertExpectations(t)
}

func TestDeleteTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	taskID := uuid.New()
	mockService.On("Delete", taskID).Return(nil)

	err := mockService.Delete(taskID)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestDeleteTask_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	taskID := uuid.New()
	mockService.On("Delete", taskID).Return(errors.New("deletion failed"))

	err := mockService.Delete(taskID)

	assert.Error(t, err)
	mockService.AssertExpectations(t)
}

func TestUpdateTask_Success(t *testing.T) {
	mockService := new(MockTaskService)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	taskID := uuid.New()
	mockService.On("Update", taskID, 1, "TaskName", 100.0).Return(task, nil)

	updatedTask, err := mockService.Update(taskID, 1, "TaskName", 100.0)

	assert.NoError(t, err)
	assert.Equal(t, task, updatedTask)
	mockService.AssertExpectations(t)
}

func TestUpdateTask_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	taskID := uuid.New()
	mockService.On("Update", taskID, 1, "TaskName", 100.0).Return((*models.Task)(nil), errors.New("update failed"))

	updatedTask, err := mockService.Update(taskID, 1, "TaskName", 100.0)

	assert.Error(t, err)
	assert.Nil(t, updatedTask)
	mockService.AssertExpectations(t)
}

func TestGetTaskByID_Success(t *testing.T) {
	mockService := new(MockTaskService)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	taskID := uuid.New()
	mockService.On("GetTaskByID", taskID).Return(task, nil)

	receivedTask, err := mockService.GetTaskByID(taskID)

	assert.NoError(t, err)
	assert.Equal(t, task, receivedTask)
	mockService.AssertExpectations(t)
}

func TestGetTaskByID_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	taskID := uuid.New()
	mockService.On("GetTaskByID", taskID).Return((*models.Task)(nil), errors.New("task not found"))

	receivedTask, err := mockService.GetTaskByID(taskID)

	assert.Error(t, err)
	assert.Nil(t, receivedTask)
	mockService.AssertExpectations(t)
}

func TestGetAllTasks_Success(t *testing.T) {
	mockService := new(MockTaskService)
	tasks := []models.Task{
		{Name: "TaskName1", PricePerSingle: 100.0, Category: 1},
		{Name: "TaskName2", PricePerSingle: 200.0, Category: 2},
	}
	mockService.On("GetAllTasks").Return(tasks, nil)

	receivedTasks, err := mockService.GetAllTasks()

	assert.NoError(t, err)
	assert.Equal(t, tasks, receivedTasks)
	mockService.AssertExpectations(t)
}

func TestGetAllTasks_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("GetAllTasks").Return(([]models.Task)(nil), errors.New("tasks not found"))

	receivedTasks, err := mockService.GetAllTasks()

	assert.Error(t, err)
	assert.Nil(t, receivedTasks)
	mockService.AssertExpectations(t)
}

func TestGetTasksInCategory_Success(t *testing.T) {
	mockService := new(MockTaskService)
	tasks := []models.Task{
		{Name: "TaskName1", PricePerSingle: 100.0, Category: 1},
		{Name: "TaskName2", PricePerSingle: 200.0, Category: 1},
	}
	mockService.On("GetTasksInCategory", 1).Return(tasks, nil)

	receivedTasks, err := mockService.GetTasksInCategory(1)

	assert.NoError(t, err)
	assert.Equal(t, tasks, receivedTasks)
	mockService.AssertExpectations(t)
}

func TestGetTasksInCategory_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("GetTasksInCategory", 1).Return(([]models.Task)(nil), errors.New("tasks not found"))

	receivedTasks, err := mockService.GetTasksInCategory(1)

	assert.Error(t, err)
	assert.Nil(t, receivedTasks)
	mockService.AssertExpectations(t)
}

func TestGetTaskByName_Success(t *testing.T) {
	mockService := new(MockTaskService)
	task := &models.Task{Name: "TaskName", PricePerSingle: 100.0, Category: 1}
	mockService.On("GetTaskByName", "TaskName").Return(task, nil)

	receivedTask, err := mockService.GetTaskByName("TaskName")

	assert.NoError(t, err)
	assert.Equal(t, task, receivedTask)
	mockService.AssertExpectations(t)
}

func TestGetTaskByName_Failure(t *testing.T) {
	mockService := new(MockTaskService)
	mockService.On("GetTaskByName", "TaskName").Return((*models.Task)(nil), errors.New("task not found"))

	receivedTask, err := mockService.GetTaskByName("TaskName")

	assert.Error(t, err)
	assert.Nil(t, receivedTask)
	mockService.AssertExpectations(t)
}

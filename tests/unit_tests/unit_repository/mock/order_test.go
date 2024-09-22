//package mock
//
//import (
//	"errors"
//	"github.com/google/uuid"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"lab3/internal/models"
//	"testing"
//)
//
//// Mock repository
//type MockOrderRepository struct {
//	mock.Mock
//}
//
//func (m *MockOrderRepository) Create(order *models.Order, orderedTasks []models.OrderedTask) (*models.Order, error) {
//	args := m.Called(order, orderedTasks)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Order), args.Error(1)
//}
//
//func (m *MockOrderRepository) Delete(id uuid.UUID) error {
//	args := m.Called(id)
//	return args.Error(0)
//}
//
//func (m *MockOrderRepository) Update(order *models.Order) (*models.Order, error) {
//	args := m.Called(order)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Order), args.Error(1)
//}
//
//func (m *MockOrderRepository) GetOrderByID(id uuid.UUID) (*models.Order, error) {
//	args := m.Called(id)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Order), args.Error(1)
//}
//
//func (m *MockOrderRepository) GetTasksInOrder(id uuid.UUID) ([]models.Task, error) {
//	args := m.Called(id)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).([]models.Task), args.Error(1)
//}
//
//func (m *MockOrderRepository) GetCurrentOrderByUserID(id uuid.UUID) (*models.Order, error) {
//	args := m.Called(id)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Order), args.Error(1)
//}
//
//func (m *MockOrderRepository) GetAllOrdersByUserID(id uuid.UUID) ([]models.Order, error) {
//	args := m.Called(id)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).([]models.Order), args.Error(1)
//}
//
//func (m *MockOrderRepository) AddTaskToOrder(orderID uuid.UUID, taskID uuid.UUID) error {
//	args := m.Called(orderID, taskID)
//	return args.Error(0)
//}
//
//func (m *MockOrderRepository) RemoveTaskFromOrder(orderID uuid.UUID, taskID uuid.UUID) error {
//	args := m.Called(orderID, taskID)
//	return args.Error(0)
//}
//
//func (m *MockOrderRepository) UpdateTaskQuantity(orderID uuid.UUID, taskID uuid.UUID, quantity int) error {
//	args := m.Called(orderID, taskID, quantity)
//	return args.Error(0)
//}
//
//func (m *MockOrderRepository) GetTaskQuantity(orderID uuid.UUID, taskID uuid.UUID) (int, error) {
//	args := m.Called(orderID, taskID)
//	return args.Int(0), args.Error(1)
//}
//
//func TestCreateOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	orderedTasks := []models.OrderedTask{
//		{Task: &models.Task{ID: uuid.New(), Name: "Task1", PricePerSingle: 100, Category: 1}, Quantity: 2},
//	}
//	mockRepo.On("Create", order, orderedTasks).Return(order, nil)
//
//	createdOrder, err := mockRepo.Create(order, orderedTasks)
//
//	assert.NoError(t, err)
//	assert.Equal(t, order, createdOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestCreateOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	orderedTasks := []models.OrderedTask{
//		{Task: &models.Task{ID: uuid.New(), Name: "Task1", PricePerSingle: 100, Category: 1}, Quantity: 2},
//	}
//	mockRepo.On("Create", order, orderedTasks).Return((*models.Order)(nil), errors.New("creation failed"))
//
//	createdOrder, err := mockRepo.Create(order, orderedTasks)
//
//	assert.Error(t, err)
//	assert.Nil(t, createdOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestDeleteOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	mockRepo.On("Delete", orderID).Return(nil)
//
//	err := mockRepo.Delete(orderID)
//
//	assert.NoError(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestDeleteOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	mockRepo.On("Delete", orderID).Return(errors.New("deletion failed"))
//
//	err := mockRepo.Delete(orderID)
//
//	assert.Error(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	mockRepo.On("Update", order).Return(order, nil)
//
//	updatedOrder, err := mockRepo.Update(order)
//
//	assert.NoError(t, err)
//	assert.Equal(t, order, updatedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	mockRepo.On("Update", order).Return((*models.Order)(nil), errors.New("update failed"))
//
//	updatedOrder, err := mockRepo.Update(order)
//
//	assert.Error(t, err)
//	assert.Nil(t, updatedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetOrderByID_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	orderID := order.ID
//	mockRepo.On("GetOrderByID", orderID).Return(order, nil)
//
//	receivedOrder, err := mockRepo.GetOrderByID(orderID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, order, receivedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetOrderByID_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	mockRepo.On("GetOrderByID", orderID).Return((*models.Order)(nil), errors.New("order not found"))
//
//	receivedOrder, err := mockRepo.GetOrderByID(orderID)
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetTasksInOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	tasks := []models.Task{
//		{ID: uuid.New(), Name: "Task1", PricePerSingle: 100, Category: 1},
//		{ID: uuid.New(), Name: "Task2", PricePerSingle: 200, Category: 2},
//	}
//	orderID := uuid.New()
//	mockRepo.On("GetTasksInOrder", orderID).Return(tasks, nil)
//
//	receivedTasks, err := mockRepo.GetTasksInOrder(orderID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, tasks, receivedTasks)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetTasksInOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	mockRepo.On("GetTasksInOrder", orderID).Return(([]models.Task)(nil), errors.New("tasks not found"))
//
//	receivedTasks, err := mockRepo.GetTasksInOrder(orderID)
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedTasks)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetCurrentOrderByUserID_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
//	userID := order.UserID
//	mockRepo.On("GetCurrentOrderByUserID", userID).Return(order, nil)
//
//	receivedOrder, err := mockRepo.GetCurrentOrderByUserID(userID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, order, receivedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetCurrentOrderByUserID_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	userID := uuid.New()
//	mockRepo.On("GetCurrentOrderByUserID", userID).Return((*models.Order)(nil), errors.New("order not found"))
//
//	receivedOrder, err := mockRepo.GetCurrentOrderByUserID(userID)
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedOrder)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetAllOrdersByUserID_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orders := []models.Order{
//		{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address1"},
//		{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 2, Address: "Address2"},
//	}
//	userID := orders[0].UserID
//	mockRepo.On("GetAllOrdersByUserID", userID).Return(orders, nil)
//
//	receivedOrders, err := mockRepo.GetAllOrdersByUserID(userID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, orders, receivedOrders)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetAllOrdersByUserID_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	userID := uuid.New()
//	mockRepo.On("GetAllOrdersByUserID", userID).Return(([]models.Order)(nil), errors.New("orders not found"))
//
//	receivedOrders, err := mockRepo.GetAllOrdersByUserID(userID)
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedOrders)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestAddTaskToOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	mockRepo.On("AddTaskToOrder", orderID, taskID).Return(nil)
//
//	err := mockRepo.AddTaskToOrder(orderID, taskID)
//
//	assert.NoError(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestAddTaskToOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	mockRepo.On("AddTaskToOrder", orderID, taskID).Return(errors.New("addition failed"))
//
//	err := mockRepo.AddTaskToOrder(orderID, taskID)
//
//	assert.Error(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestRemoveTaskFromOrder_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	mockRepo.On("RemoveTaskFromOrder", orderID, taskID).Return(nil)
//
//	err := mockRepo.RemoveTaskFromOrder(orderID, taskID)
//
//	assert.NoError(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestRemoveTaskFromOrder_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	mockRepo.On("RemoveTaskFromOrder", orderID, taskID).Return(errors.New("removal failed"))
//
//	err := mockRepo.RemoveTaskFromOrder(orderID, taskID)
//
//	assert.Error(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateTaskQuantity_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	quantity := 5
//	mockRepo.On("UpdateTaskQuantity", orderID, taskID, quantity).Return(nil)
//
//	err := mockRepo.UpdateTaskQuantity(orderID, taskID, quantity)
//
//	assert.NoError(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateTaskQuantity_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	quantity := 5
//	mockRepo.On("UpdateTaskQuantity", orderID, taskID, quantity).Return(errors.New("update failed"))
//
//	err := mockRepo.UpdateTaskQuantity(orderID, taskID, quantity)
//
//	assert.Error(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetTaskQuantity_Success(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	quantity := 5
//	mockRepo.On("GetTaskQuantity", orderID, taskID).Return(quantity, nil)
//
//	receivedQuantity, err := mockRepo.GetTaskQuantity(orderID, taskID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, quantity, receivedQuantity)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetTaskQuantity_Failure(t *testing.T) {
//	mockRepo := new(MockOrderRepository)
//	orderID := uuid.New()
//	taskID := uuid.New()
//	mockRepo.On("GetTaskQuantity", orderID, taskID).Return(0, errors.New("quantity not found"))
//
//	receivedQuantity, err := mockRepo.GetTaskQuantity(orderID, taskID)
//
//	assert.Error(t, err)
//	assert.Equal(t, 0, receivedQuantity)
//	mockRepo.AssertExpectations(t)
//}

package mock

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock repository
type MockOrderRepository struct {
	mock.Mock
}

func (m *MockOrderRepository) Create(order *models.Order, orderedTasks []models.OrderedTask) (*models.Order, error) {
	args := m.Called(order, orderedTasks)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockOrderRepository) Update(order *models.Order) (*models.Order, error) {
	args := m.Called(order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) GetOrderByID(id uuid.UUID) (*models.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) GetTasksInOrder(id uuid.UUID) ([]models.Task, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Task), args.Error(1)
}

func (m *MockOrderRepository) GetCurrentOrderByUserID(id uuid.UUID) (*models.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderRepository) GetAllOrdersByUserID(id uuid.UUID) ([]models.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Order), args.Error(1)
}

func (m *MockOrderRepository) AddTaskToOrder(orderID uuid.UUID, taskID uuid.UUID) error {
	args := m.Called(orderID, taskID)
	return args.Error(0)
}

func (m *MockOrderRepository) RemoveTaskFromOrder(orderID uuid.UUID, taskID uuid.UUID) error {
	args := m.Called(orderID, taskID)
	return args.Error(0)
}

func (m *MockOrderRepository) UpdateTaskQuantity(orderID uuid.UUID, taskID uuid.UUID, quantity int) error {
	args := m.Called(orderID, taskID, quantity)
	return args.Error(0)
}

func (m *MockOrderRepository) GetTaskQuantity(orderID uuid.UUID, taskID uuid.UUID) (int, error) {
	args := m.Called(orderID, taskID)
	return args.Int(0), args.Error(1)
}

// fixture
func setupMockОrderRepo() *MockOrderRepository {
	return new(MockOrderRepository)
}

func TestCreateOrder_Success(t *testing.T) {
	mockRepo := setupMockОrderRepo()
	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
	orderedTasks := []models.OrderedTask{
		{Task: &models.Task{ID: uuid.New(), Name: "Task1", PricePerSingle: 100, Category: 1}, Quantity: 2},
	}
	mockRepo.On("Create", order, orderedTasks).Return(order, nil)

	createdOrder, err := mockRepo.Create(order, orderedTasks)

	assert.NoError(t, err)
	assert.Equal(t, order, createdOrder)
	mockRepo.AssertExpectations(t)
}

func TestCreateOrder_Failure(t *testing.T) {
	mockRepo := setupMockОrderRepo()
	order := &models.Order{ID: uuid.New(), WorkerID: uuid.New(), UserID: uuid.New(), Status: 1, Address: "Address"}
	orderedTasks := []models.OrderedTask{
		{Task: &models.Task{ID: uuid.New(), Name: "Task1", PricePerSingle: 100, Category: 1}, Quantity: 2},
	}
	mockRepo.On("Create", order, orderedTasks).Return((*models.Order)(nil), errors.New("creation failed"))

	createdOrder, err := mockRepo.Create(order, orderedTasks)

	assert.Error(t, err)
	assert.Nil(t, createdOrder)
	mockRepo.AssertExpectations(t)
}

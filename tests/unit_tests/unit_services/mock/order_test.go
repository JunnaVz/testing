package mock

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

type MockOrderService struct {
	mock.Mock
}

func (m *MockOrderService) Create(order *models.Order) (*models.Order, error) {
	args := m.Called(order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) Update(order *models.Order) (*models.Order, error) {
	args := m.Called(order)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockOrderService) GetOrderByID(id uuid.UUID) (*models.Order, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Order), args.Error(1)
}

func (m *MockOrderService) GetAllOrders() ([]models.Order, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Order), args.Error(1)
}

func TestCreateOrder_Success(t *testing.T) {
	mockService := new(MockOrderService)
	order := &models.Order{ID: uuid.New(), Address: "Test Order"}
	mockService.On("Create", order).Return(order, nil)

	createdOrder, err := mockService.Create(order)

	assert.NoError(t, err)
	assert.Equal(t, order, createdOrder)
	mockService.AssertExpectations(t)
}

func TestCreateOrder_Failure(t *testing.T) {
	mockService := new(MockOrderService)
	order := &models.Order{ID: uuid.New(), Address: "Test Order"}
	mockService.On("Create", order).Return((*models.Order)(nil), errors.New("creation failed"))

	createdOrder, err := mockService.Create(order)

	assert.Error(t, err)
	assert.Nil(t, createdOrder)
	mockService.AssertExpectations(t)
}

func TestUpdateOrder_Success(t *testing.T) {
	mockService := new(MockOrderService)
	order := &models.Order{ID: uuid.New(), Address: "Updated Order"}
	mockService.On("Update", order).Return(order, nil)

	updatedOrder, err := mockService.Update(order)

	assert.NoError(t, err)
	assert.Equal(t, order, updatedOrder)
	mockService.AssertExpectations(t)
}

func TestUpdateOrder_Failure(t *testing.T) {
	mockService := new(MockOrderService)
	order := &models.Order{ID: uuid.New(), Address: "Updated Order"}
	mockService.On("Update", order).Return((*models.Order)(nil), errors.New("update failed"))

	updatedOrder, err := mockService.Update(order)

	assert.Error(t, err)
	assert.Nil(t, updatedOrder)
	mockService.AssertExpectations(t)
}

func TestDeleteOrder_Success(t *testing.T) {
	mockService := new(MockOrderService)
	orderID := uuid.New()
	mockService.On("Delete", orderID).Return(nil)

	err := mockService.Delete(orderID)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestDeleteOrder_Failure(t *testing.T) {
	mockService := new(MockOrderService)
	orderID := uuid.New()
	mockService.On("Delete", orderID).Return(errors.New("deletion failed"))

	err := mockService.Delete(orderID)

	assert.Error(t, err)
	mockService.AssertExpectations(t)
}

func TestGetOrderByID_Success(t *testing.T) {
	mockService := new(MockOrderService)
	order := &models.Order{ID: uuid.New(), Address: "Test Order"}
	orderID := order.ID
	mockService.On("GetOrderByID", orderID).Return(order, nil)

	receivedOrder, err := mockService.GetOrderByID(orderID)

	assert.NoError(t, err)
	assert.Equal(t, order, receivedOrder)
	mockService.AssertExpectations(t)
}

func TestGetOrderByID_Failure(t *testing.T) {
	mockService := new(MockOrderService)
	orderID := uuid.New()
	mockService.On("GetOrderByID", orderID).Return((*models.Order)(nil), errors.New("order not found"))

	receivedOrder, err := mockService.GetOrderByID(orderID)

	assert.Error(t, err)
	assert.Nil(t, receivedOrder)
	mockService.AssertExpectations(t)
}

func TestGetAllOrders_Success(t *testing.T) {
	mockService := new(MockOrderService)
	orders := []models.Order{
		{ID: uuid.New(), Address: "Order 1"},
		{ID: uuid.New(), Address: "Order 2"},
	}
	mockService.On("GetAllOrders").Return(orders, nil)

	receivedOrders, err := mockService.GetAllOrders()

	assert.NoError(t, err)
	assert.Equal(t, orders, receivedOrders)
	mockService.AssertExpectations(t)
}

func TestGetAllOrders_Failure(t *testing.T) {
	mockService := new(MockOrderService)
	mockService.On("GetAllOrders").Return(([]models.Order)(nil), errors.New("orders not found"))

	receivedOrders, err := mockService.GetAllOrders()

	assert.Error(t, err)
	assert.Nil(t, receivedOrders)
	mockService.AssertExpectations(t)
}

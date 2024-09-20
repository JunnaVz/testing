package no_mock

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"lab3/internal/models"
	"lab3/internal/repository/postgres"
	services "lab3/internal/services"
	"os"
	"testing"
	"time"
)

func TestOrderServiceCreateOrder_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	task := models.Task{
		ID:             uuid.New(),
		Name:           "Test Task",
		PricePerSingle: 10,
		Category:       1,
	}
	userID := uuid.New()
	tasks := []models.OrderedTask{
		{Task: &task, Quantity: 2},
	}

	// Act
	order, err := orderService.CreateOrder(userID, "Test Address", time.Now().Add(24*time.Hour), tasks)

	// Assert
	require.Error(t, err)
	require.Nil(t, order)
}

func TestOrderServiceCreateOrder_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	task := models.Task{
		ID:             uuid.New(),
		Name:           "Test Task",
		PricePerSingle: 10,
		Category:       1,
	}

	invalidTasks := []models.OrderedTask{
		{Task: &task, Quantity: -1},
	}

	// Act
	order, err := orderService.CreateOrder(uuid.New(), "", time.Now().Add(-24*time.Hour), invalidTasks)

	// Assert
	require.Error(t, err)
	require.Nil(t, order)
}

func TestOrderServiceDeleteOrder_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	task := models.Task{
		ID:             uuid.New(),
		Name:           "Test Task",
		PricePerSingle: 10,
		Category:       1,
	}

	tasks := []models.OrderedTask{
		{Task: &task, Quantity: 1},
	}

	_, err = orderService.CreateOrder(uuid.New(), "Test Address", time.Now().Add(24*time.Hour), tasks)

	require.Error(t, err)
}

func TestOrderServiceDeleteOrder_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	err = orderService.DeleteOrder(uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetOrderByID_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	task := models.Task{
		ID:             uuid.New(),
		Name:           "Test Task",
		PricePerSingle: 10,
		Category:       1,
	}

	tasks := []models.OrderedTask{
		{Task: &task, Quantity: 1},
	}

	_, err = orderService.CreateOrder(uuid.New(), "Test Address", time.Now().Add(24*time.Hour), tasks)

	require.Error(t, err)
}

func TestOrderServiceGetOrderByID_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	invalidOrderID := uuid.New()

	// Act
	fetchedOrder, err := orderService.GetOrderByID(invalidOrderID)

	// Assert
	require.Error(t, err)
	require.Nil(t, fetchedOrder)
}

func TestOrderServiceGetCurrentOrderByUserID_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	userID := uuid.New()

	// Act
	_, err = orderService.GetCurrentOrderByUserID(userID)

	// Assert
	require.Error(t, err)
	//require.NotNil(t, order)
}

func TestOrderServiceGetCurrentOrderByUserID_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.GetCurrentOrderByUserID(uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetAllOrdersByUserID_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	userID := uuid.New()

	// Act
	_, err = orderService.GetAllOrdersByUserID(userID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetAllOrdersByUserID_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.GetAllOrdersByUserID(uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceUpdate_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	workerID := uuid.New()

	// Act
	_, err = orderService.Update(orderID, 1, 5, workerID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceUpdate_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.Update(uuid.New(), 1, 5, uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceAddTask_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	err = orderService.AddTask(orderID, taskID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceAddTask_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	err = orderService.AddTask(uuid.New(), uuid.New())

	// Assert
	require.Error(t, err)
}
func TestOrderServiceRemoveTask_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	err = orderService.RemoveTask(orderID, taskID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceRemoveTask_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	err = orderService.RemoveTask(uuid.New(), uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceIncrementTaskQuantity_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	_, err = orderService.IncrementTaskQuantity(orderID, taskID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceIncrementTaskQuantity_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.IncrementTaskQuantity(uuid.New(), uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceDecrementTaskQuantity_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	_, err = orderService.DecrementTaskQuantity(orderID, taskID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceDecrementTaskQuantity_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.DecrementTaskQuantity(uuid.New(), uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceSetTaskQuantity_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	err = orderService.SetTaskQuantity(orderID, taskID, 5)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceSetTaskQuantity_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	err = orderService.SetTaskQuantity(uuid.New(), uuid.New(), 5)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetTaskQuantity_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()
	taskID := uuid.New()

	// Act
	_, err = orderService.GetTaskQuantity(orderID, taskID)

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetTaskQuantity_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.GetTaskQuantity(uuid.New(), uuid.New())

	// Assert
	require.Error(t, err)
}

func TestOrderServiceFilter_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	params := map[string]string{"status": "1"}

	// Act
	_, err = orderService.Filter(params)

	// Assert
	require.Nil(t, err)
}

func TestOrderServiceFilter_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.Filter(map[string]string{"status": "invalid"})

	// Assert
	require.Error(t, err)
}

func TestOrderServiceGetTotalPrice_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	orderID := uuid.New()

	// Act
	_, err = orderService.GetTotalPrice(orderID)

	// Assert
	require.Nil(t, err)
}

func TestOrderServiceGetTotalPrice_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	orderRepository := postgres.NewOrderRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	userRepository := postgres.NewUserRepository(db)
	workerRepository := postgres.NewWorkerRepository(db)
	logger := log.New(f)
	orderService := services.NewOrderService(orderRepository, workerRepository, taskRepository, userRepository, logger)

	// Act
	_, err = orderService.GetTotalPrice(uuid.New())

	// Assert
	require.Nil(t, err)
}

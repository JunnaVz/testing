package itc_repository

import (
	"context"
	"github.com/google/uuid"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"lab3/internal/models"
	"lab3/internal/repository/postgres"
	"log"
	"testing"
	"time"
)

func TestOrderRepositoryCreate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:   uuid.New(),
		Status:   1,
		Address:  "Test Address",
		Deadline: time.Now().Add(24 * time.Hour),
	}

	createdOrder := order
	require.NotNil(t, createdOrder)
	require.Equal(t, order.UserID, createdOrder.UserID)
	require.Equal(t, order.Status, createdOrder.Status)
	require.Equal(t, order.Address, createdOrder.Address)
	require.WithinDuration(t, order.Deadline, createdOrder.Deadline, time.Second)
}

func TestOrderRepositoryCreate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder, err := orderRepository.Create(order, nil)

	require.Error(t, err)
	require.Nil(t, createdOrder)
}

func TestOrderRepositoryDelete_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}

	createdOrder := order
	require.NotNil(t, createdOrder)

	err := orderRepository.Delete(createdOrder.ID)
	require.Error(t, err)

	receivedOrder := (*models.Order)(nil)
	require.Nil(t, receivedOrder)
}

func TestOrderRepositoryDelete_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	err := orderRepository.Delete(uuid.New())
	require.Error(t, err)
}

func TestOrderRepositoryUpdate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}

	createdOrder := order
	require.NotNil(t, createdOrder)

	createdOrder.Status = 2
	updatedOrder := createdOrder
	require.NotNil(t, updatedOrder)
	require.Equal(t, 2, updatedOrder.Status)
}

func TestOrderRepositoryUpdate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}

	createdOrder := order
	require.NotNil(t, createdOrder)

	createdOrder.UserID = uuid.Nil
	updatedOrder := (*models.Order)(nil)
	require.Nil(t, updatedOrder)
}

func TestOrderRepositoryGetOrderByID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	receivedOrder := createdOrder
	require.NotNil(t, receivedOrder)
	require.Equal(t, createdOrder.ID, receivedOrder.ID)
}

func TestOrderRepositoryGetOrderByID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	receivedOrder, err := orderRepository.GetOrderByID(uuid.New())
	require.Error(t, err)
	require.Nil(t, receivedOrder)
}

func TestOrderRepositoryGetTasksInOrder_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)
	_ = postgres.NewTaskRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask := task
	require.NotNil(t, createdTask)

	tasks := []*models.Task{createdTask}
	require.NotNil(t, tasks)
	require.Len(t, tasks, 1)
	require.Equal(t, createdTask.ID, tasks[0].ID)
}

func TestOrderRepositoryGetTasksInOrder_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	_, err := orderRepository.GetTasksInOrder(uuid.New())
	require.Nil(t, err)
}

func TestOrderRepositoryGetCurrentOrderByUserID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	receivedOrder := createdOrder
	require.NotNil(t, receivedOrder)
	require.Equal(t, createdOrder.ID, receivedOrder.ID)
}

func TestOrderRepositoryGetCurrentOrderByUserID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	receivedOrder, err := orderRepository.GetCurrentOrderByUserID(uuid.New())
	require.Error(t, err)
	require.Nil(t, receivedOrder)
}

func TestOrderRepositoryGetAllOrdersByUserID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	orders, err := orderRepository.GetAllOrdersByUserID(order.UserID)
	require.Nil(t, err)
	require.Len(t, orders, 0)
}

func TestOrderRepositoryGetAllOrdersByUserID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	orders, err := orderRepository.GetAllOrdersByUserID(uuid.New())
	require.NoError(t, err)
	require.Len(t, orders, 0)
}

func TestOrderRepositoryAddTaskToOrder_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)
	_ = postgres.NewTaskRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask := task
	require.NotNil(t, createdTask)
}

func TestOrderRepositoryAddTaskToOrder_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	err := orderRepository.AddTaskToOrder(uuid.New(), uuid.New())
	require.Error(t, err)
}

func TestOrderRepositoryRemoveTaskFromOrder_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)
	_ = postgres.NewTaskRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask := task
	require.NotNil(t, createdTask)
}

func TestOrderRepositoryRemoveTaskFromOrder_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	err := orderRepository.RemoveTaskFromOrder(uuid.New(), uuid.New())
	require.Nil(t, err)
}

func TestOrderRepositoryUpdateTaskQuantity_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)
	_ = postgres.NewTaskRepository(db)

	order := &models.Order{
		UserID:       uuid.New(),
		Status:       1,
		Address:      "Test Address",
		CreationDate: time.Now(),
		Deadline:     time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask := task
	require.NotNil(t, createdTask)
}

func TestOrderRepositoryUpdateTaskQuantity_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	err := orderRepository.UpdateTaskQuantity(uuid.New(), uuid.New(), 5)
	require.Nil(t, err)
}

func TestOrderRepositoryGetTaskQuantity_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	_ = postgres.NewOrderRepository(db)
	_ = postgres.NewTaskRepository(db)

	order := &models.Order{
		ID:       uuid.New(),
		UserID:   uuid.New(),
		Status:   1,
		Address:  "Test Address",
		Deadline: time.Now().Add(24 * time.Hour),
	}
	createdOrder := order
	require.NotNil(t, createdOrder)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask := task
	require.NotNil(t, createdTask)

	quantity := 5
	require.Equal(t, 5, quantity)
}

func TestOrderRepositoryGetTaskQuantity_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	orderRepository := postgres.NewOrderRepository(db)

	quantity, err := orderRepository.GetTaskQuantity(uuid.New(), uuid.New())
	require.Error(t, err)
	require.Equal(t, 0, quantity)
}

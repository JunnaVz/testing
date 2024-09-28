package itc_services

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
)

func TestTaskServiceCreate_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	task, err := taskService.Create("Test Task", 100.0, 1)

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Test Task", task.Name)
	require.Equal(t, 100.0, task.PricePerSingle)
	require.Equal(t, 1, task.Category)
}

func TestTaskServiceCreate_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	task, err := taskService.Create("", -10.0, 99)

	// Assert
	require.Error(t, err)
	require.Nil(t, task)
}

func TestTaskServiceUpdate_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	task, err := taskService.Create("Test Task", 100.0, 1)
	require.NoError(t, err)

	// Act
	updatedTask, err := taskService.Update(task.ID, 2, "Updated Task", 200.0)

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Updated Task", updatedTask.Name)
	require.Equal(t, 200.0, updatedTask.PricePerSingle)
	require.Equal(t, 2, updatedTask.Category)
}

func TestTaskServiceUpdate_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	task, err := taskService.Create("Test Task", 100.0, 1)
	require.NoError(t, err)

	// Act
	updatedTask, err := taskService.Update(task.ID, 2, "", -10.0)

	// Assert
	require.Error(t, err)
	require.Nil(t, updatedTask)
}

func TestTaskServiceDelete_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	task, err := taskService.Create("Task to Delete", 50.0, 1)
	require.NoError(t, err)

	// Act
	err = taskService.Delete(task.ID)

	// Assert
	require.NoError(t, err)
}

func TestTaskServiceDelete_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	err = taskService.Delete(uuid.New())

	// Assert
	require.Error(t, err)
}

func TestTaskServiceGetTaskByID_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	task, err := taskService.Create("Task to Retrieve", 150.0, 1)
	require.NoError(t, err)

	// Act
	receivedTask, err := taskService.GetTaskByID(task.ID)

	// Assert
	require.NoError(t, err)
	require.Equal(t, task.Name, receivedTask.Name)
}

func TestTaskServiceGetTaskByID_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	foundTask, err := taskService.GetTaskByID(uuid.New())

	// Assert
	require.Error(t, err)
	require.Nil(t, foundTask)
}

func TestTaskServiceGetTasksInCategory_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	_, err = taskService.Create("Task 1", 100.0, 1)
	require.NoError(t, err)
	_, err = taskService.Create("Task 2", 200.0, 1)
	require.NoError(t, err)

	// Act
	tasks, err := taskService.GetTasksInCategory(1)

	// Assert
	require.NoError(t, err)
	require.Len(t, tasks, 2)
	require.Equal(t, 1, tasks[0].Category)
	require.Equal(t, 1, tasks[1].Category)
}

func TestTaskServiceGetTasksInCategory_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	invalidCategory := -1

	// Act
	tasks, err := taskService.GetTasksInCategory(invalidCategory)

	// Assert
	require.Error(t, err)
	require.Nil(t, tasks)
}

func TestTaskServiceGetAllTasks_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Create tasks
	task1 := &models.Task{Name: "Task1", PricePerSingle: 100.0, Category: 1}
	task2 := &models.Task{Name: "Task2", PricePerSingle: 150.0, Category: 2}
	_, _ = taskService.Create(task1.Name, task1.PricePerSingle, task1.Category)
	_, _ = taskService.Create(task2.Name, task2.PricePerSingle, task2.Category)

	// Act
	tasks, err := taskService.GetAllTasks()

	// Assert
	require.NoError(t, err)
	require.NotEmpty(t, tasks)
	require.Equal(t, 2, len(tasks))
}

func TestTaskServiceGetAllTasks_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	tasks, err := taskService.GetAllTasks()

	// Assert
	require.NoError(t, err)
	require.Empty(t, tasks) // If no tasks are present
}

func TestTaskServiceGetTaskByName_Success(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	task := &models.Task{Name: "Unique Task", PricePerSingle: 200.0, Category: 1}
	_, _ = taskService.Create(task.Name, task.PricePerSingle, task.Category)

	// Act
	foundTask, err := taskService.GetTaskByName("Unique Task")

	// Assert
	require.NoError(t, err)
	require.NotNil(t, foundTask)
	require.Equal(t, task.Name, foundTask.Name)
}

func TestTaskServiceGetTaskByName_Failure(t *testing.T) {
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

	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	taskService := services.NewTaskService(taskRepository, logger)

	// Act
	foundTask, err := taskService.GetTaskByName("Nonexistent Task")

	// Assert
	require.Error(t, err)
	require.Nil(t, foundTask)
}

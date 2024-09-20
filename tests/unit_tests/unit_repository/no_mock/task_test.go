package no_mock

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
)

func TestTaskRepositoryCreate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask, err := taskRepository.Create(task)

	require.NoError(t, err)
	require.Equal(t, task.Name, createdTask.Name)
	require.Equal(t, task.PricePerSingle, createdTask.PricePerSingle)
	require.Equal(t, task.Category, createdTask.Category)
}

func TestTaskRepositoryCreate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask, err := taskRepository.Create(task)

	require.Error(t, err)
	require.Nil(t, createdTask)
}

func TestTaskRepositoryDelete_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}

	createdTask, err := taskRepository.Create(task)
	require.NoError(t, err)

	err = taskRepository.Delete(createdTask.ID)
	require.NoError(t, err)

	receivedTask, err := taskRepository.GetTaskByID(createdTask.ID)
	require.Error(t, err)
	require.Nil(t, receivedTask)
}

func TestTaskRepositoryDelete_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	_ = taskRepository.Delete(uuid.New())
	require.Nil(t, nil)
}

func TestTaskRepositoryUpdate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}

	createdTask, err := taskRepository.Create(task)
	require.NoError(t, err)

	createdTask.Name = "Updated Name"
	updatedTask, err := taskRepository.Update(createdTask)
	require.NoError(t, err)
	require.Equal(t, "Updated Name", updatedTask.Name)
}

func TestTaskRepositoryUpdate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}

	createdTask, err := taskRepository.Create(task)
	require.NoError(t, err)

	createdTask.Name = ""
	updatedTask, err := taskRepository.Update(createdTask)
	require.Error(t, err)
	require.Nil(t, updatedTask)
}

func TestTaskRepositoryGetTaskByID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask, err := taskRepository.Create(task)
	require.NoError(t, err)

	receivedTask, err := taskRepository.GetTaskByID(createdTask.ID)
	require.NoError(t, err)
	require.Equal(t, createdTask.Name, receivedTask.Name)
}

func TestTaskRepositoryGetTaskByID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	receivedTask, err := taskRepository.GetTaskByID(uuid.New())
	require.Error(t, err)
	require.Nil(t, receivedTask)
}

func TestTaskRepositoryGetTaskByName_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task := &models.Task{
		Name:           "TaskName",
		PricePerSingle: 100.0,
		Category:       1,
	}
	createdTask, err := taskRepository.Create(task)
	require.NoError(t, err)

	receivedTask, err := taskRepository.GetTaskByName(createdTask.Name)
	require.NoError(t, err)
	require.Equal(t, createdTask.Name, receivedTask.Name)
}

func TestTaskRepositoryGetTaskByName_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	receivedTask, err := taskRepository.GetTaskByName("Unknown Name")
	require.Error(t, err)
	require.Nil(t, receivedTask)
}

func TestTaskRepositoryGetAllTasks_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task1 := &models.Task{
		Name:           "TaskName1",
		PricePerSingle: 100.0,
		Category:       1,
	}
	task2 := &models.Task{
		Name:           "TaskName2",
		PricePerSingle: 100.0,
		Category:       1,
	}
	_, err := taskRepository.Create(task1)
	require.NoError(t, err)
	_, err = taskRepository.Create(task2)
	require.NoError(t, err)

	receivedCategories, err := taskRepository.GetAllTasks()
	require.NoError(t, err)
	require.Len(t, receivedCategories, 2)
}

func TestTaskRepositoryGetAllTasks_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	receivedCategories, err := taskRepository.GetAllTasks()
	require.NoError(t, err)
	require.Len(t, receivedCategories, 0)
}

func TestTaskRepositoryGetTasksInCategory_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	task1 := &models.Task{
		Name:           "TaskName1",
		PricePerSingle: 100.0,
		Category:       1,
	}
	task2 := &models.Task{
		Name:           "TaskName2",
		PricePerSingle: 100.0,
		Category:       1,
	}
	_, err := taskRepository.Create(task1)
	require.NoError(t, err)
	_, err = taskRepository.Create(task2)
	require.NoError(t, err)

	receivedCategories, err := taskRepository.GetTasksInCategory(1)
	require.NoError(t, err)
	require.Len(t, receivedCategories, 2)
}

func TestTaskRepositoryGetTasksInCategory_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	taskRepository := postgres.NewTaskRepository(db)

	receivedCategories, err := taskRepository.GetTasksInCategory(0)
	require.NoError(t, err)
	require.Len(t, receivedCategories, 0)
}

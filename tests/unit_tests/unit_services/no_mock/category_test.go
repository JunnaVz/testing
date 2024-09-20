package no_mock

import (
	"context"
	"github.com/charmbracelet/log"
	_ "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"lab3/internal/models"
	_ "lab3/internal/models"
	"lab3/internal/repository/postgres"
	services "lab3/internal/services"
	"os"
	"testing"

	_ "lab3/internal/repository/repository_errors"
	_ "lab3/internal/services/service_errors"
)

func TestCategoryServiceCreate_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	category, err := categoryService.Create("New Category")

	// Assert
	require.NoError(t, err)
	require.Equal(t, "New Category", category.Name)
}

func TestCategoryServiceCreate_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	category, err := categoryService.Create("")

	// Assert
	require.Error(t, err)
	require.Nil(t, category)
}

func TestCategoryServiceUpdate_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	category, err := categoryService.Create("Category")
	require.NoError(t, err)

	// Act
	updatedCategory, err := categoryService.Update(category)

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Category", updatedCategory.Name)
}

func TestCategoryServiceUpdate_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	category, err := categoryService.Create("Category")
	require.NoError(t, err)

	// Act
	updatedCategory, err := categoryService.Update(category)

	// Assert
	require.Nil(t, err)
	require.NotNil(t, updatedCategory)
}

func TestCategoryServiceDelete_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	category, err := categoryService.Create("Category to Delete")
	require.NoError(t, err)

	// Act
	err = categoryService.Delete(category.ID)

	// Assert
	require.NoError(t, err)
}

func TestCategoryServiceDelete_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	err = categoryService.Delete(0)

	// Assert
	require.Nil(t, err)
}

func TestCategoryServiceGetByID_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	category, err := categoryService.Create("Category to Retrieve")
	require.NoError(t, err)

	// Act
	receivedCategory, err := categoryService.GetByID(category.ID)

	// Assert
	require.NoError(t, err)
	require.Equal(t, category.Name, receivedCategory.Name)
}

func TestCategoryServiceGetByID_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	receivedCategory, err := categoryService.GetByID(0)

	// Assert
	require.Error(t, err)
	require.Nil(t, receivedCategory)
}

func TestCategoryServiceGetAll_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	_, _ = categoryService.Create("Category 1")
	_, _ = categoryService.Create("Category 2")

	categories, err := categoryService.GetAll()

	// Assert
	require.NoError(t, err)
	require.Len(t, categories, 2)
}

func TestCategoryServiceGetAll_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	categories, err := categoryService.GetAll()

	// Assert
	require.NoError(t, err)
	require.Len(t, categories, 0)
}

func TestCategoryServiceGetTasksInCategory_Success(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	category, err := categoryService.Create("Category with Tasks")
	require.NoError(t, err)

	task, err := taskRepository.Create(&models.Task{
		Name:           "Task",
		Category:       category.ID,
		PricePerSingle: 10,
	})
	require.NoError(t, err)

	// Act
	tasks, err := categoryService.GetTasksInCategory(category.ID)

	// Assert
	require.NoError(t, err)
	require.Len(t, tasks, 1)
	require.Equal(t, task.Name, tasks[0].Name)
}

func TestCategoryServiceGetTasksInCategory_Failure(t *testing.T) {
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

	categoryRepository := postgres.NewCategoryRepository(db)
	taskRepository := postgres.NewTaskRepository(db)
	logger := log.New(f)
	categoryService := services.NewCategoryService(categoryRepository, taskRepository, logger)

	// Act
	tasks, err := categoryService.GetTasksInCategory(0)

	// Assert
	require.NoError(t, err)
	require.Len(t, tasks, 0)
}

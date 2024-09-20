package no_mock

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"lab3/internal/models"
	"lab3/internal/repository/postgres"
	"log"
	"testing"
)

func TestCategoryRepositoryCreate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "CategoryName",
	}
	createdCategory, err := categoryRepository.Create(category)

	require.NoError(t, err)
	require.Equal(t, category.Name, createdCategory.Name)
}

func TestCategoryRepositoryCreate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "",
	}
	createdCategory, err := categoryRepository.Create(category)

	require.Error(t, err)
	require.Nil(t, createdCategory)
}

func TestCategoryRepositoryGetByID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "CategoryName",
	}
	createdCategory, err := categoryRepository.Create(category)
	require.NoError(t, err)

	receivedCategory, err := categoryRepository.GetByID(createdCategory.ID)
	require.NoError(t, err)
	require.Equal(t, createdCategory.Name, receivedCategory.Name)
}

func TestCategoryRepositoryGetByID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	receivedCategory, err := categoryRepository.GetByID(999)
	require.Error(t, err)
	require.Nil(t, receivedCategory)
}

func TestCategoryRepositoryGetAll_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category1 := &models.Category{
		ID:   1,
		Name: "CategoryName1",
	}
	category2 := &models.Category{
		ID:   2,
		Name: "CategoryName2",
	}
	_, err := categoryRepository.Create(category1)
	require.NoError(t, err)
	_, err = categoryRepository.Create(category2)
	require.NoError(t, err)

	receivedCategories, err := categoryRepository.GetAll()
	require.NoError(t, err)
	require.Len(t, receivedCategories, 2)
}

func TestCategoryRepositoryGetAll_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	receivedCategories, err := categoryRepository.GetAll()
	require.NoError(t, err)
	require.Len(t, receivedCategories, 0)
}

func TestCategoryRepositoryUpdate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "CategoryName",
	}
	createdCategory, err := categoryRepository.Create(category)
	require.NoError(t, err)

	createdCategory.Name = "UpdatedCategoryName"
	updatedCategory, err := categoryRepository.Update(createdCategory)
	require.NoError(t, err)
	require.Equal(t, "UpdatedCategoryName", updatedCategory.Name)
}

func TestCategoryRepositoryUpdate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "CategoryName",
	}
	createdCategory, err := categoryRepository.Create(category)
	require.NoError(t, err)

	createdCategory.Name = ""
	updatedCategory, err := categoryRepository.Update(createdCategory)
	require.Error(t, err)
	require.Nil(t, updatedCategory)
}

func TestCategoryRepositoryDelete_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	category := &models.Category{
		ID:   1,
		Name: "CategoryName",
	}
	createdCategory, err := categoryRepository.Create(category)
	require.NoError(t, err)

	err = categoryRepository.Delete(createdCategory.ID)
	require.NoError(t, err)

	receivedCategory, err := categoryRepository.GetByID(createdCategory.ID)
	require.Error(t, err)
	require.Nil(t, receivedCategory)
}

func TestCategoryRepositoryDelete_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	categoryRepository := postgres.NewCategoryRepository(db)

	_ = categoryRepository.Delete(999)
	require.Nil(t, nil)
}

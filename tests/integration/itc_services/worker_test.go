// tests/unit_services/no_mock/worker_test.go
package itc_services

import (
	"context"
	"github.com/charmbracelet/log"
	"github.com/google/uuid"
	_ "github.com/google/uuid"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	"lab3/internal/models"
	"lab3/internal/repository/postgres"
	_ "lab3/internal/services"
	services "lab3/internal/services"
	"lab3/password_hash"
	"os"
	"testing"
)

func TestWorkerServiceCreate_Success(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}

	// Act
	createdWorker, err := workerService.Create(worker, "password123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, worker.Name, createdWorker.Name)
	require.Equal(t, worker.Surname, createdWorker.Surname)
	require.Equal(t, worker.Address, createdWorker.Address)
	require.Equal(t, worker.PhoneNumber, createdWorker.PhoneNumber)
	require.Equal(t, worker.Email, createdWorker.Email)
	require.NotEqual(t, "password123", createdWorker.Password) // Password should be hashed
}

func TestWorkerServiceCreate_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatal(err)
	}

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		Name:        "",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "password123",
	}

	// Act
	createdWorker, err := workerService.Create(worker, "password123")

	// Assert
	require.Error(t, err)
	require.Nil(t, createdWorker)
}

func TestWorkerServiceLogin_Success(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	_, err = workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	loggedInWorker, err := workerService.Login(worker.Email, "password123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, worker.Email, loggedInWorker.Email)
}

func TestWorkerServiceLogin_Failure(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	// Act
	loggedInWorker, err := workerService.Login("nonexistent@email.com", "password123")

	// Assert
	require.Error(t, err)
	require.Nil(t, loggedInWorker)
}

func TestWorkerServiceGetWorkerByID_Success(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	receivedWorker, err := workerService.GetWorkerByID(createdWorker.ID)

	// Assert
	require.NoError(t, err)
	require.Equal(t, createdWorker.Name, receivedWorker.Name)
	require.Equal(t, createdWorker.Surname, receivedWorker.Surname)
	require.Equal(t, createdWorker.Address, receivedWorker.Address)
	require.Equal(t, createdWorker.PhoneNumber, receivedWorker.PhoneNumber)
	require.Equal(t, createdWorker.Email, receivedWorker.Email)
	require.Equal(t, createdWorker.Password, receivedWorker.Password)
}

func TestWorkerServiceGetWorkerByID_Failure(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	// Act
	receivedWorker, err := workerService.GetWorkerByID(uuid.New())

	// Assert
	require.Error(t, err)
	require.Nil(t, receivedWorker)
}

func TestWorkerServiceUpdate_Success(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	updatedWorker, err := workerService.Update(createdWorker.ID, "Updated Name", createdWorker.Surname, createdWorker.Email, createdWorker.Address, createdWorker.PhoneNumber, createdWorker.Role, "newpassword123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Updated Name", updatedWorker.Name)
}

func TestWorkerServiceUpdate_Failure(t *testing.T) {
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

	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	updatedWorker, err := workerService.Update(createdWorker.ID, "", createdWorker.Surname, createdWorker.Email, createdWorker.Address, createdWorker.PhoneNumber, createdWorker.Role, "newpassword123")

	// Assert
	require.Error(t, err)
	require.Nil(t, updatedWorker)
}

func TestWorkerService_Delete_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	err = workerService.Delete(createdWorker.ID)

	// Assert
	require.NoError(t, err)

	// Try to get deleted worker
	deletedWorker, err := workerService.GetWorkerByID(createdWorker.ID)
	require.Error(t, err)
	require.Nil(t, deletedWorker)
}

func TestWorkerService_Delete_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	// Act
	err = workerService.Delete(uuid.New()) // Attempt to delete non-existent worker

	// Assert
	require.Error(t, err)
}

func TestWorkerService_GetWorkerByID_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	foundWorker, err := workerService.GetWorkerByID(createdWorker.ID)

	// Assert
	require.NoError(t, err)
	require.Equal(t, createdWorker.Email, foundWorker.Email)
}

func TestWorkerService_GetWorkerByID_Failure(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	// Act
	worker, err := workerService.GetWorkerByID(uuid.New()) // Try to get non-existent worker

	// Assert
	require.Error(t, err)
	require.Nil(t, worker)
}

func TestWorkerService_GetWorkersByRole_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker1 := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	worker2 := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test1",
		Surname:     "Test1",
		Email:       "test1@email.com",
		Address:     "Test1",
		PhoneNumber: "+79999999998",
		Role:        1,
		Password:    "password123",
	}
	_, err = workerService.Create(worker1, "password123")
	require.NoError(t, err)
	_, err = workerService.Create(worker2, "password123")
	require.NoError(t, err)

	// Act
	workers, err := workerService.GetWorkersByRole(1)

	// Assert
	require.NoError(t, err)
	require.Len(t, workers, 2)
	require.Equal(t, "test@email.com", workers[0].Email)
	require.Equal(t, "test1@email.com", workers[1].Email)
}

func TestWorkerService_GetAverageOrderRate_Success(t *testing.T) {
	// Arrange
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Fatal(err)
		}
	}(dbContainer, context.Background())

	f, err := os.OpenFile("tests.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	require.NoError(t, err)

	logger := log.New(f)
	workerRepository := postgres.NewWorkerRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	workerService := services.NewWorkerService(workerRepository, passwordHash, logger)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "Test",
		Surname:     "Test",
		Email:       "test@email.com",
		Address:     "Test",
		PhoneNumber: "+79999999999",
		Role:        1,
		Password:    "password123",
	}
	createdWorker, err := workerService.Create(worker, "password123")
	require.NoError(t, err)

	// Act
	averageRate, err := workerService.GetAverageOrderRate(createdWorker)

	// Assert
	require.Error(t, err)
	require.GreaterOrEqual(t, averageRate, 0.0)
}

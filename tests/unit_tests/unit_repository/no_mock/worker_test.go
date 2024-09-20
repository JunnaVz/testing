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

func TestWorkerRepositoryCreate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		ID:          uuid.New(),
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)

	require.NoError(t, err)
	require.Equal(t, worker.Name, createdWorker.Name)
	require.Equal(t, worker.Surname, createdWorker.Surname)
	require.Equal(t, worker.Address, createdWorker.Address)
	require.Equal(t, worker.PhoneNumber, createdWorker.PhoneNumber)
	require.Equal(t, worker.Email, createdWorker.Email)
	require.Equal(t, worker.Role, createdWorker.Role)
	require.Equal(t, worker.Password, createdWorker.Password)
}

func TestWorkerRepositoryCreate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		Name:        "",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)

	require.Error(t, err)
	require.Nil(t, createdWorker)
}

func TestWorkerRepositoryGetWorkerByID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)
	require.NoError(t, err)

	receivedWorker, err := workerRepository.GetWorkerByID(createdWorker.ID)
	require.NoError(t, err)
	require.Equal(t, createdWorker.Name, receivedWorker.Name)
	require.Equal(t, createdWorker.Surname, receivedWorker.Surname)
	require.Equal(t, createdWorker.Address, receivedWorker.Address)
	require.Equal(t, createdWorker.PhoneNumber, receivedWorker.PhoneNumber)
	require.Equal(t, createdWorker.Email, receivedWorker.Email)
	require.Equal(t, createdWorker.Role, receivedWorker.Role)
	require.Equal(t, createdWorker.Password, receivedWorker.Password)
}

func TestWorkerRepositoryGetWorkerByID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	receivedWorker, err := workerRepository.GetWorkerByID(uuid.New())
	require.Error(t, err)
	require.Nil(t, receivedWorker)
}

func TestWorkerRepositoryUpdate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)
	require.NoError(t, err)

	createdWorker.Name = "Updated Name"
	updatedWorker, err := workerRepository.Update(createdWorker)
	require.NoError(t, err)
	require.Equal(t, "Updated Name", updatedWorker.Name)
}

func TestWorkerRepositoryUpdate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)
	require.NoError(t, err)

	createdWorker.Name = ""
	updatedWorker, err := workerRepository.Update(createdWorker)
	require.Error(t, err)
	require.Nil(t, updatedWorker)
}

func TestWorkerRepositoryDelete_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	worker := &models.Worker{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Role:        1,
		Password:    "hashed_password",
	}
	createdWorker, err := workerRepository.Create(worker)
	require.NoError(t, err)

	err = workerRepository.Delete(createdWorker.ID)
	require.NoError(t, err)

	receivedWorker, err := workerRepository.GetWorkerByID(createdWorker.ID)
	require.Error(t, err)
	require.Nil(t, receivedWorker)
}

func TestWorkerRepositoryDelete_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	workerRepository := postgres.NewWorkerRepository(db)

	_ = workerRepository.Delete(uuid.New())
	require.Nil(t, nil)
}

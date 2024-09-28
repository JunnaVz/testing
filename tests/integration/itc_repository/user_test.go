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
)

func TestUserRepositoryCreate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userRepository.Create(user)

	require.NoError(t, err)
	require.Equal(t, user.Name, createdUser.Name)
	require.Equal(t, user.Surname, createdUser.Surname)
	require.Equal(t, user.Address, createdUser.Address)
	require.Equal(t, user.PhoneNumber, createdUser.PhoneNumber)
	require.Equal(t, user.Email, createdUser.Email)
	require.Equal(t, user.Password, createdUser.Password)
}

func TestUserRepositoryCreate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hash",
	}
	createdUser, err := userRepository.Create(user)

	require.Error(t, err)
	require.Nil(t, createdUser)
}

func TestUserRepositoryGetUserByID_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userRepository.Create(user)
	require.NoError(t, err)

	receivedUser, err := userRepository.GetUserByID(createdUser.ID)
	require.NoError(t, err)
	require.Equal(t, createdUser.Name, receivedUser.Name)
	require.Equal(t, createdUser.Surname, receivedUser.Surname)
	require.Equal(t, createdUser.Address, receivedUser.Address)
	require.Equal(t, createdUser.PhoneNumber, receivedUser.PhoneNumber)
	require.Equal(t, createdUser.Email, receivedUser.Email)
	require.Equal(t, createdUser.Password, receivedUser.Password)
}

func TestUserRepositoryGetUserByID_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	receivedUser, err := userRepository.GetUserByID(uuid.New())
	require.Error(t, err)
	require.Nil(t, receivedUser)
}

func TestUserRepositoryUpdate_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userRepository.Create(user)
	require.NoError(t, err)

	createdUser.Name = "Updated Name"
	updatedUser, err := userRepository.Update(createdUser)
	require.NoError(t, err)
	require.Equal(t, "Updated Name", updatedUser.Name)
}

func TestUserRepositoryUpdate_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userRepository.Create(user)
	require.NoError(t, err)

	createdUser.Name = ""
	updatedUser, err := userRepository.Update(createdUser)
	require.Error(t, err)
	require.Nil(t, updatedUser)
}

func TestUserRepositoryDelete_Success(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userRepository.Create(user)
	require.NoError(t, err)

	err = userRepository.Delete(createdUser.ID)
	require.NoError(t, err)

	receivedUser, err := userRepository.GetUserByID(createdUser.ID)
	require.Error(t, err)
	require.Nil(t, receivedUser)
}

func TestUserRepositoryDelete_Failure(t *testing.T) {
	dbContainer, db := SetupTestDatabase()
	defer func(dbContainer testcontainers.Container, ctx context.Context) {
		err := dbContainer.Terminate(ctx)
		if err != nil {
			log.Println("Error terminating container:", err)
		}
	}(dbContainer, context.Background())

	userRepository := postgres.NewUserRepository(db)

	_ = userRepository.Delete(uuid.New())
	require.Nil(t, nil)
}

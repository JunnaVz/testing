// tests/unit_services/no_mock/user_test.go
package no_mock

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

func TestUserServiceRegister_Success(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}

	// Act
	createdUser, err := userService.Register(user, "password123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, user.Name, createdUser.Name)
	require.Equal(t, user.Surname, createdUser.Surname)
	require.Equal(t, user.Address, createdUser.Address)
	require.Equal(t, user.PhoneNumber, createdUser.PhoneNumber)
	require.Equal(t, user.Email, createdUser.Email)
	require.NotEqual(t, "password123", createdUser.Password) // Password should be hashed
}

func TestUserServiceRegister_Failure(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hash",
	}

	// Act
	createdUser, err := userService.Register(user, "password123")

	// Assert
	require.Error(t, err)
	require.Nil(t, createdUser)
}

func TestUserServiceLogin_Success(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	_, err = userService.Register(user, "password123")
	require.NoError(t, err)

	// Act
	loggedInUser, err := userService.Login(user.Email, "password123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, user.Email, loggedInUser.Email)
}

func TestUserServiceLogin_Failure(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	// Act
	loggedInUser, err := userService.Login("nonexistent@email.com", "password123")

	// Assert
	require.Error(t, err)
	require.Nil(t, loggedInUser)
}

func TestUserServiceGetUserByID_Success(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userService.Register(user, "password123")
	require.NoError(t, err)

	// Act
	receivedUser, err := userService.GetUserByID(createdUser.ID)

	// Assert
	require.NoError(t, err)
	require.Equal(t, createdUser.Name, receivedUser.Name)
	require.Equal(t, createdUser.Surname, receivedUser.Surname)
	require.Equal(t, createdUser.Address, receivedUser.Address)
	require.Equal(t, createdUser.PhoneNumber, receivedUser.PhoneNumber)
	require.Equal(t, createdUser.Email, receivedUser.Email)
	require.Equal(t, createdUser.Password, receivedUser.Password)
}

func TestUserServiceGetUserByID_Failure(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	// Act
	receivedUser, err := userService.GetUserByID(uuid.New())

	// Assert
	require.Error(t, err)
	require.Nil(t, receivedUser)
}

func TestUserServiceGetUserByEmail_Success(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	_, err = userService.Register(user, "password123")
	require.NoError(t, err)

	// Act
	receivedUser, err := userService.GetUserByEmail(user.Email)

	// Assert
	require.NoError(t, err)
	require.Equal(t, user.Email, receivedUser.Email)
}

func TestUserServiceGetUserByEmail_Failure(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	// Act
	receivedUser, err := userService.GetUserByEmail("nonexistent@email.com")

	// Assert
	require.Error(t, err)
	require.Nil(t, receivedUser)
}

func TestUserServiceUpdate_Success(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userService.Register(user, "password123")
	require.NoError(t, err)

	// Act
	updatedUser, err := userService.Update(createdUser.ID, "Updated Name", createdUser.Surname, createdUser.Email, createdUser.Address, createdUser.PhoneNumber, "newpassword123")

	// Assert
	require.NoError(t, err)
	require.Equal(t, "Updated Name", updatedUser.Name)
}

func TestUserServiceUpdate_Failure(t *testing.T) {
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

	userRepository := postgres.NewUserRepository(db)
	passwordHash := password_hash.NewPasswordHash()
	logger := log.New(f)

	userService := services.NewUserService(userRepository, passwordHash, logger)

	user := &models.User{
		Name:        "First Name",
		Surname:     "Last Name",
		Address:     "Address",
		PhoneNumber: "+79999999999",
		Email:       "test@email.com",
		Password:    "hashed_password",
	}
	createdUser, err := userService.Register(user, "password123")
	require.NoError(t, err)

	// Act
	updatedUser, err := userService.Update(createdUser.ID, "", createdUser.Surname, createdUser.Email, createdUser.Address, createdUser.PhoneNumber, "newpassword123")

	// Assert
	require.Error(t, err)
	require.Nil(t, updatedUser)
}

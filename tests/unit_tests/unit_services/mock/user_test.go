package mock

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock service
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(name string, email string) (*models.User, error) {
	args := m.Called(name, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockUserService) Update(id uuid.UUID, name string, email string) (*models.User, error) {
	args := m.Called(id, name, email)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetUserByID(id uuid.UUID) (*models.User, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func (m *MockUserService) GetAllUsers() ([]models.User, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.User), args.Error(1)
}

func (m *MockUserService) GetUserByName(name string) (*models.User, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.User), args.Error(1)
}

func TestCreateUser_Success(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "UserName", Email: "test@gmail.com"}
	mockService.On("Create", "UserName", "test@gmail.com").Return(user, nil)

	createdUser, err := mockService.Create("UserName", "test@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, user, createdUser)
	mockService.AssertExpectations(t)
}

func TestCreateUser_Failure(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("Create", "UserName", "test@gmail.com").Return((*models.User)(nil), errors.New("creation failed"))

	createdUser, err := mockService.Create("UserName", "test@gmail.com")

	assert.Error(t, err)
	assert.Nil(t, createdUser)
	mockService.AssertExpectations(t)
}

func TestDeleteUser_Success(t *testing.T) {
	mockService := new(MockUserService)
	userID := uuid.New()
	mockService.On("Delete", userID).Return(nil)

	err := mockService.Delete(userID)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestDeleteUser_Failure(t *testing.T) {
	mockService := new(MockUserService)
	userID := uuid.New()
	mockService.On("Delete", userID).Return(errors.New("deletion failed"))

	err := mockService.Delete(userID)

	assert.Error(t, err)
	mockService.AssertExpectations(t)
}

func TestUpdateUser_Success(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "UserName", Email: "test@gmail.com"}
	userID := uuid.New()
	mockService.On("Update", userID, "UserName", "test@gmail.com").Return(user, nil)

	updatedUser, err := mockService.Update(userID, "UserName", "test@gmail.com")

	assert.NoError(t, err)
	assert.Equal(t, user, updatedUser)
	mockService.AssertExpectations(t)
}

func TestUpdateUser_Failure(t *testing.T) {
	mockService := new(MockUserService)
	userID := uuid.New()
	mockService.On("Update", userID, "UserName", "test@gmail.com").Return((*models.User)(nil), errors.New("update failed"))

	updatedUser, err := mockService.Update(userID, "UserName", "test@gmail.com")

	assert.Error(t, err)
	assert.Nil(t, updatedUser)
	mockService.AssertExpectations(t)
}

func TestGetUserByID_Success(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "UserName", Email: "test@gmail.com"}
	userID := uuid.New()
	mockService.On("GetUserByID", userID).Return(user, nil)

	receivedUser, err := mockService.GetUserByID(userID)

	assert.NoError(t, err)
	assert.Equal(t, user, receivedUser)
	mockService.AssertExpectations(t)
}

func TestGetUserByID_Failure(t *testing.T) {
	mockService := new(MockUserService)
	userID := uuid.New()
	mockService.On("GetUserByID", userID).Return((*models.User)(nil), errors.New("user not found"))

	receivedUser, err := mockService.GetUserByID(userID)

	assert.Error(t, err)
	assert.Nil(t, receivedUser)
	mockService.AssertExpectations(t)
}

func TestGetAllUsers_Success(t *testing.T) {
	mockService := new(MockUserService)
	users := []models.User{
		{Name: "UserName1", Email: "test1@gmail.com"},
		{Name: "UserName2", Email: "test2@gmail.com"},
	}
	mockService.On("GetAllUsers").Return(users, nil)

	receivedUsers, err := mockService.GetAllUsers()

	assert.NoError(t, err)
	assert.Equal(t, users, receivedUsers)
	mockService.AssertExpectations(t)
}

func TestGetAllUsers_Failure(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetAllUsers").Return(([]models.User)(nil), errors.New("users not found"))

	receivedUsers, err := mockService.GetAllUsers()

	assert.Error(t, err)
	assert.Nil(t, receivedUsers)
	mockService.AssertExpectations(t)
}

func TestGetUserByName_Success(t *testing.T) {
	mockService := new(MockUserService)
	user := &models.User{Name: "UserName", Email: "test@gmail.com"}
	mockService.On("GetUserByName", "UserName").Return(user, nil)

	receivedUser, err := mockService.GetUserByName("UserName")

	assert.NoError(t, err)
	assert.Equal(t, user, receivedUser)
	mockService.AssertExpectations(t)
}

func TestGetUserByName_Failure(t *testing.T) {
	mockService := new(MockUserService)
	mockService.On("GetUserByName", "UserName").Return((*models.User)(nil), errors.New("user not found"))

	receivedUser, err := mockService.GetUserByName("UserName")

	assert.Error(t, err)
	assert.Nil(t, receivedUser)
	mockService.AssertExpectations(t)
}

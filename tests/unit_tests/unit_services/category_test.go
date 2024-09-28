package unit_services

import (
	"errors"
	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock service
type MockCategoryService struct {
	mock.Mock
}

func (m *MockCategoryService) Create(name string) (*models.Category, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) Update(id uuid.UUID, name string) (*models.Category, error) {
	args := m.Called(id, name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) Delete(id uuid.UUID) error {
	args := m.Called(id)
	return args.Error(0)
}

func (m *MockCategoryService) GetCategoryByID(id uuid.UUID) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryService) GetAllCategories() ([]models.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryService) GetCategoryByName(name string) (*models.Category, error) {
	args := m.Called(name)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func TestCreateCategory_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	category := &models.Category{Name: "CategoryName"}
	mockService.On("Create", "CategoryName").Return(category, nil)

	createdCategory, err := mockService.Create("CategoryName")

	assert.NoError(t, err)
	assert.Equal(t, category, createdCategory)
	mockService.AssertExpectations(t)
}

func TestCreateCategory_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	mockService.On("Create", "CategoryName").Return((*models.Category)(nil), errors.New("creation failed"))

	createdCategory, err := mockService.Create("CategoryName")

	assert.Error(t, err)
	assert.Nil(t, createdCategory)
	mockService.AssertExpectations(t)
}

func TestUpdateCategory_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	category := &models.Category{Name: "CategoryName"}
	categoryID := uuid.New()
	mockService.On("Update", categoryID, "CategoryName").Return(category, nil)

	updatedCategory, err := mockService.Update(categoryID, "CategoryName")

	assert.NoError(t, err)
	assert.Equal(t, category, updatedCategory)
	mockService.AssertExpectations(t)
}

func TestUpdateCategory_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	categoryID := uuid.New()
	mockService.On("Update", categoryID, "CategoryName").Return((*models.Category)(nil), errors.New("update failed"))

	updatedCategory, err := mockService.Update(categoryID, "CategoryName")

	assert.Error(t, err)
	assert.Nil(t, updatedCategory)
	mockService.AssertExpectations(t)
}

func TestDeleteCategory_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	categoryID := uuid.New()
	mockService.On("Delete", categoryID).Return(nil)

	err := mockService.Delete(categoryID)

	assert.NoError(t, err)
	mockService.AssertExpectations(t)
}

func TestDeleteCategory_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	categoryID := uuid.New()
	mockService.On("Delete", categoryID).Return(errors.New("deletion failed"))

	err := mockService.Delete(categoryID)

	assert.Error(t, err)
	mockService.AssertExpectations(t)
}

func TestGetCategoryByID_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	category := &models.Category{Name: "CategoryName"}
	categoryID := uuid.New()
	mockService.On("GetCategoryByID", categoryID).Return(category, nil)

	receivedCategory, err := mockService.GetCategoryByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, category, receivedCategory)
	mockService.AssertExpectations(t)
}

func TestGetCategoryByID_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	categoryID := uuid.New()
	mockService.On("GetCategoryByID", categoryID).Return((*models.Category)(nil), errors.New("category not found"))

	receivedCategory, err := mockService.GetCategoryByID(categoryID)

	assert.Error(t, err)
	assert.Nil(t, receivedCategory)
	mockService.AssertExpectations(t)
}

func TestGetAllCategories_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	categories := []models.Category{
		{Name: "CategoryName1"},
		{Name: "CategoryName2"},
	}
	mockService.On("GetAllCategories").Return(categories, nil)

	receivedCategories, err := mockService.GetAllCategories()

	assert.NoError(t, err)
	assert.Equal(t, categories, receivedCategories)
	mockService.AssertExpectations(t)
}

func TestGetAllCategories_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	mockService.On("GetAllCategories").Return(([]models.Category)(nil), errors.New("categories not found"))

	receivedCategories, err := mockService.GetAllCategories()

	assert.Error(t, err)
	assert.Nil(t, receivedCategories)
	mockService.AssertExpectations(t)
}

func TestGetCategoryByName_Success(t *testing.T) {
	mockService := new(MockCategoryService)
	category := &models.Category{Name: "CategoryName"}
	mockService.On("GetCategoryByName", "CategoryName").Return(category, nil)

	receivedCategory, err := mockService.GetCategoryByName("CategoryName")

	assert.NoError(t, err)
	assert.Equal(t, category, receivedCategory)
	mockService.AssertExpectations(t)
}

func TestGetCategoryByName_Failure(t *testing.T) {
	mockService := new(MockCategoryService)
	mockService.On("GetCategoryByName", "CategoryName").Return((*models.Category)(nil), errors.New("category not found"))

	receivedCategory, err := mockService.GetCategoryByName("CategoryName")

	assert.Error(t, err)
	assert.Nil(t, receivedCategory)
	mockService.AssertExpectations(t)
}

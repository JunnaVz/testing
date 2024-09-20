package mock

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"lab3/internal/models"
	"testing"
)

// Mock repository
type MockCategoryRepository struct {
	mock.Mock
}

func (m *MockCategoryRepository) GetAll() ([]models.Category, error) {
	args := m.Called()
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).([]models.Category), args.Error(1)
}

func (m *MockCategoryRepository) GetByID(id int) (*models.Category, error) {
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Create(category *models.Category) (*models.Category, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Update(category *models.Category) (*models.Category, error) {
	args := m.Called(category)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*models.Category), args.Error(1)
}

func (m *MockCategoryRepository) Delete(id int) error {
	args := m.Called(id)
	return args.Error(0)
}

func TestCreateCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &models.Category{Name: "CategoryName"}
	mockRepo.On("Create", category).Return(category, nil)

	createdCategory, err := mockRepo.Create(category)

	assert.NoError(t, err)
	assert.Equal(t, category, createdCategory)
	mockRepo.AssertExpectations(t)
}

func TestCreateCategory_Failure(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &models.Category{Name: "CategoryName"}
	mockRepo.On("Create", category).Return((*models.Category)(nil), errors.New("creation failed"))

	createdCategory, err := mockRepo.Create(category)

	assert.Error(t, err)
	assert.Nil(t, createdCategory)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &models.Category{Name: "CategoryName"}
	mockRepo.On("Update", category).Return(category, nil)

	updatedCategory, err := mockRepo.Update(category)

	assert.NoError(t, err)
	assert.Equal(t, category, updatedCategory)
	mockRepo.AssertExpectations(t)
}

func TestUpdateCategory_Failure(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &models.Category{Name: "CategoryName"}
	mockRepo.On("Update", category).Return((*models.Category)(nil), errors.New("update failed"))

	updatedCategory, err := mockRepo.Update(category)

	assert.Error(t, err)
	assert.Nil(t, updatedCategory)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCategory_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryID := 1
	mockRepo.On("Delete", categoryID).Return(nil)

	err := mockRepo.Delete(categoryID)

	assert.NoError(t, err)
	mockRepo.AssertExpectations(t)
}

func TestDeleteCategory_Failure(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryID := 1
	mockRepo.On("Delete", categoryID).Return(errors.New("deletion failed"))

	err := mockRepo.Delete(categoryID)

	assert.Error(t, err)
	mockRepo.AssertExpectations(t)
}

func TestGetCategoryByID_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	category := &models.Category{Name: "CategoryName"}
	categoryID := 1
	mockRepo.On("GetByID", categoryID).Return(category, nil)

	receivedCategory, err := mockRepo.GetByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, category, receivedCategory)
	mockRepo.AssertExpectations(t)
}

func TestGetCategoryByID_Failure(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categoryID := 1
	mockRepo.On("GetByID", categoryID).Return((*models.Category)(nil), errors.New("category not found"))

	receivedCategory, err := mockRepo.GetByID(categoryID)

	assert.Error(t, err)
	assert.Nil(t, receivedCategory)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCategories_Success(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	categories := []models.Category{
		{Name: "CategoryName1"},
		{Name: "CategoryName2"},
	}
	mockRepo.On("GetAll").Return(categories, nil)

	receivedCategories, err := mockRepo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, categories, receivedCategories)
	mockRepo.AssertExpectations(t)
}

func TestGetAllCategories_Failure(t *testing.T) {
	mockRepo := new(MockCategoryRepository)
	mockRepo.On("GetAll").Return(([]models.Category)(nil), errors.New("categories not found"))

	receivedCategories, err := mockRepo.GetAll()

	assert.Error(t, err)
	assert.Nil(t, receivedCategories)
	mockRepo.AssertExpectations(t)
}

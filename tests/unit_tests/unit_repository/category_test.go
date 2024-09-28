//package mock
//
//import (
//	"errors"
//	"github.com/stretchr/testify/assert"
//	"github.com/stretchr/testify/mock"
//	"lab3/internal/models"
//	"testing"
//)
//
//// Mock repository
//type MockCategoryRepository struct {
//	mock.Mock
//}
//
//func (m *MockCategoryRepository) GetAll() ([]models.Category, error) {
//	args := m.Called()
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).([]models.Category), args.Error(1)
//}
//
//func (m *MockCategoryRepository) GetByID(id int) (*models.Category, error) {
//	args := m.Called(id)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Category), args.Error(1)
//}
//
//func (m *MockCategoryRepository) Create(category *models.Category) (*models.Category, error) {
//	args := m.Called(category)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Category), args.Error(1)
//}
//
//func (m *MockCategoryRepository) Update(category *models.Category) (*models.Category, error) {
//	args := m.Called(category)
//	if args.Get(0) == nil {
//		return nil, args.Error(1)
//	}
//	return args.Get(0).(*models.Category), args.Error(1)
//}
//
//func (m *MockCategoryRepository) Delete(id int) error {
//	args := m.Called(id)
//	return args.Error(0)
//}
//
//var mockRepo *MockCategoryRepository
//
//func TestMain(m *testing.M) {
//	mockRepo = new(MockCategoryRepository)
//	m.Run()
//}
//
//func TestCreateCategory_Success(t *testing.T) {
//	category := &models.Category{Name: "CategoryName"}
//	mockRepo.On("Create", category).Return(category, nil)
//
//	createdCategory, err := mockRepo.Create(category)
//
//	assert.Nil(t, err)
//	assert.Equal(t, category, createdCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestCreateCategory_Failure(t *testing.T) {
//	category := &models.Category{Name: "CategoryName"}
//	mockRepo.On("Create", category).Return((*models.Category)(nil), errors.New("creation failed"))
//
//	createdCategory, err := mockRepo.Create(category)
//
//	assert.Error(t, err)
//	assert.Nil(t, createdCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateCategory_Success(t *testing.T) {
//	category := &models.Category{Name: "CategoryName"}
//	mockRepo.On("Update", category).Return(category, nil)
//
//	updatedCategory, err := mockRepo.Update(category)
//
//	assert.NoError(t, err)
//	assert.Equal(t, category, updatedCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestUpdateCategory_Failure(t *testing.T) {
//	category := &models.Category{Name: "CategoryName"}
//	mockRepo.On("Update", category).Return((*models.Category)(nil), errors.New("update failed"))
//
//	updatedCategory, err := mockRepo.Update(category)
//
//	assert.Error(t, err)
//	assert.Nil(t, updatedCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestDeleteCategory_Success(t *testing.T) {
//	categoryID := 1
//	mockRepo.On("Delete", categoryID).Return(nil)
//
//	err := mockRepo.Delete(categoryID)
//
//	assert.NoError(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestDeleteCategory_Failure(t *testing.T) {
//	categoryID := 1
//	mockRepo.On("Delete", categoryID).Return(errors.New("deletion failed"))
//
//	err := mockRepo.Delete(categoryID)
//
//	assert.Error(t, err)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetCategoryByID_Success(t *testing.T) {
//	category := &models.Category{Name: "CategoryName"}
//	categoryID := 1
//	mockRepo.On("GetByID", categoryID).Return(category, nil)
//
//	receivedCategory, err := mockRepo.GetByID(categoryID)
//
//	assert.NoError(t, err)
//	assert.Equal(t, category, receivedCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetCategoryByID_Failure(t *testing.T) {
//	categoryID := 1
//	mockRepo.On("GetByID", categoryID).Return((*models.Category)(nil), errors.New("category not found"))
//
//	receivedCategory, err := mockRepo.GetByID(categoryID)
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedCategory)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetAllCategories_Success(t *testing.T) {
//	categories := []models.Category{
//		{Name: "CategoryName1"},
//		{Name: "CategoryName2"},
//	}
//	mockRepo.On("GetAll").Return(categories, nil)
//
//	receivedCategories, err := mockRepo.GetAll()
//
//	assert.NoError(t, err)
//	assert.Equal(t, categories, receivedCategories)
//	mockRepo.AssertExpectations(t)
//}
//
//func TestGetAllCategories_Failure(t *testing.T) {
//	mockRepo.On("GetAll").Return(([]models.Category)(nil), errors.New("categories not found"))
//
//	receivedCategories, err := mockRepo.GetAll()
//
//	assert.Error(t, err)
//	assert.Nil(t, receivedCategories)
//	mockRepo.AssertExpectations(t)
//}

package unit_repository

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"lab3/internal/models"
	"lab3/tests/unit_tests/unit_repository/testdata"
	"os"
	"testing"
)

// Mock repository
type MockCategoryRepository struct {
	GetAllFunc  func() ([]models.Category, error)
	GetByIDFunc func(int) (*models.Category, error)
	CreateFunc  func(*models.Category) (*models.Category, error)
	UpdateFunc  func(*models.Category) (*models.Category, error)
	DeleteFunc  func(int) error
}

func (m *MockCategoryRepository) GetAll() ([]models.Category, error) {
	return m.GetAllFunc()
}

func (m *MockCategoryRepository) GetByID(id int) (*models.Category, error) {
	return m.GetByIDFunc(id)
}

func (m *MockCategoryRepository) Create(category *models.Category) (*models.Category, error) {
	return m.CreateFunc(category)
}

func (m *MockCategoryRepository) Update(category *models.Category) (*models.Category, error) {
	return m.UpdateFunc(category)
}

func (m *MockCategoryRepository) Delete(id int) error {
	return m.DeleteFunc(id)
}

// Глобальные переменные для использования в тестах
var mockRepo *MockCategoryRepository

// TestMain используется для глобальной фикстуры, которая запускается один раз для всех тестов
func TestMain(m *testing.M) {
	setupSuite()    // Глобальная фикстура: настройка перед всеми тестами
	code := m.Run() // Запуск всех тестов
	teardownSuite() // Очистка после выполнения всех тестов
	os.Exit(code)   // Завершение программы
}

// setupSuite для настройки перед в��еми тестами
func setupSuite() {
	mockRepo = &MockCategoryRepository{} // Инициализация общего мока
}

// teardownSuite для очистки после всех тестов
func teardownSuite() {
	mockRepo = nil // Очистка
}

// setup для каждого теста (инициализация)
func setupTest(t *testing.T) *MockCategoryRepository {
	repo := &MockCategoryRepository{} // Каждый тест получает свою копию мок-репозитория
	t.Cleanup(func() {
		teardownTest() // Удаление после теста
	})
	return repo
}

// Teardown для каждого теста (очистка)
func teardownTest() {
	// Можно добавлять дополнительные действия по очистке
}

func TestCreateCategory_Success(t *testing.T) {
	mockRepo := setupTest(t)
	category := testdata.NewCategoryBuilder().WithName("CategoryName").Build()

	mockRepo.CreateFunc = func(c *models.Category) (*models.Category, error) {
		return c, nil
	}

	createdCategory, err := mockRepo.Create(category)

	assert.Nil(t, err)
	assert.Equal(t, category, createdCategory)
}

func TestCreateCategory_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	category := testdata.CategoryMother{}.CustomCategory("CategoryName")

	mockRepo.CreateFunc = func(c *models.Category) (*models.Category, error) {
		return nil, errors.New("creation failed")
	}

	createdCategory, err := mockRepo.Create(category)

	assert.Error(t, err)
	assert.Nil(t, createdCategory)
}

func TestUpdateCategory_Success(t *testing.T) {
	mockRepo := setupTest(t)
	category := testdata.NewCategoryBuilder().WithName("CategoryName").Build()

	mockRepo.UpdateFunc = func(c *models.Category) (*models.Category, error) {
		return c, nil
	}

	updatedCategory, err := mockRepo.Update(category)

	assert.NoError(t, err)
	assert.Equal(t, category, updatedCategory)
}

func TestUpdateCategory_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	category := testdata.CategoryMother{}.CustomCategory("CategoryName")

	mockRepo.UpdateFunc = func(c *models.Category) (*models.Category, error) {
		return nil, errors.New("update failed")
	}

	updatedCategory, err := mockRepo.Update(category)

	assert.Error(t, err)
	assert.Nil(t, updatedCategory)
}

func TestDeleteCategory_Success(t *testing.T) {
	mockRepo := setupTest(t)
	categoryID := 1

	mockRepo.DeleteFunc = func(id int) error {
		return nil
	}

	err := mockRepo.Delete(categoryID)

	assert.NoError(t, err)
}

func TestDeleteCategory_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	categoryID := 1

	mockRepo.DeleteFunc = func(id int) error {
		return errors.New("deletion failed")
	}

	err := mockRepo.Delete(categoryID)

	assert.Error(t, err)
}

func TestGetCategoryByID_Success(t *testing.T) {
	mockRepo := setupTest(t)
	category := testdata.NewCategoryBuilder().WithName("CategoryName").Build()
	categoryID := 1

	mockRepo.GetByIDFunc = func(id int) (*models.Category, error) {
		return category, nil
	}

	receivedCategory, err := mockRepo.GetByID(categoryID)

	assert.NoError(t, err)
	assert.Equal(t, category, receivedCategory)
}

func TestGetCategoryByID_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	categoryID := 1

	mockRepo.GetByIDFunc = func(id int) (*models.Category, error) {
		return nil, errors.New("category not found")
	}

	receivedCategory, err := mockRepo.GetByID(categoryID)

	assert.Error(t, err)
	assert.Nil(t, receivedCategory)
}

func TestGetAllCategories_Success(t *testing.T) {
	mockRepo := setupTest(t)
	categories := []models.Category{
		*testdata.NewCategoryBuilder().WithName("CategoryName1").Build(),
		*testdata.NewCategoryBuilder().WithName("CategoryName2").Build(),
	}

	mockRepo.GetAllFunc = func() ([]models.Category, error) {
		return categories, nil
	}

	receivedCategories, err := mockRepo.GetAll()

	assert.NoError(t, err)
	assert.Equal(t, categories, receivedCategories)
}

func TestGetAllCategories_Failure(t *testing.T) {
	mockRepo := setupTest(t)

	mockRepo.GetAllFunc = func() ([]models.Category, error) {
		return nil, errors.New("categories not found")
	}

	receivedCategories, err := mockRepo.GetAll()

	assert.Error(t, err)
	assert.Nil(t, receivedCategories)
}

package unit_repository

import (
	"errors"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
	"lab3/internal/models"
)

// CategoryBuilder реализует паттерн Data Builder для Category
type CategoryBuilder struct {
	category models.Category
}

// NewCategoryBuilder создает новый экземпляр CategoryBuilder с настройками по умолчанию
func NewCategoryBuilder() *CategoryBuilder {
	return &CategoryBuilder{
		category: models.Category{
			ID:   1,
			Name: "DefaultCategory",
		},
	}
}

// WithID устанавливает ID категории
func (b *CategoryBuilder) WithID(id int) *CategoryBuilder {
	b.category.ID = id
	return b
}

// WithName устанавливает имя категории
func (b *CategoryBuilder) WithName(name string) *CategoryBuilder {
	b.category.Name = name
	return b
}

// Build возвращает готовый объект Category
func (b *CategoryBuilder) Build() *models.Category {
	return &b.category
}

// CategoryMother реализует паттерн Object Mother для Category
var CategoryMother = struct {
	Default        func() *models.Category
	WithID         func(id int) *models.Category
	WithName       func(name string) *models.Category
	CustomCategory func(id int, name string) *models.Category
}{
	Default: func() *models.Category {
		return &models.Category{
			ID:   1,
			Name: "DefaultCategory",
		}
	},
	WithID: func(id int) *models.Category {
		return &models.Category{
			ID:   id,
			Name: "CategoryWithSpecificID",
		}
	},
	WithName: func(name string) *models.Category {
		return &models.Category{
			ID:   2,
			Name: name,
		}
	},
	CustomCategory: func(id int, name string) *models.Category {
		return &models.Category{
			ID:   id,
			Name: name,
		}
	},
}

// MockCategoryRepository представляет мок-репозиторий для Category
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

// Глобальная переменная для мок-репозитория (может быть полезна для глобальных настроек)
var mockRepo *MockCategoryRepository

// TestMain используется для глобальной фикстуры, которая запускается один раз для всех тестов
func TestMain(m *testing.M) {
	setupSuite()    // Глобальная фикстура: настройка перед всеми тестами
	code := m.Run() // Запуск всех тестов
	teardownSuite() // Очистка после выполнения всех тестов
	os.Exit(code)   // Завершение программы
}

// setupSuite для настройки перед всеми тестами
func setupSuite() {
	mockRepo = &MockCategoryRepository{} // Инициализация общего мока, если необходимо
}

// teardownSuite для очистки после всех тестов
func teardownSuite() {
	mockRepo = nil // Очистка
}

// setupTest и teardownTest для каждого теста
func setupTest(t *testing.T) *MockCategoryRepository {
	repo := &MockCategoryRepository{}
	t.Cleanup(func() {
		teardownTest()
	})
	return repo
}

func teardownTest() {
	// Дополнительная очистка, если требуется
}

// Пример теста с использованием Data Builder
func TestCreateCategory_Success(t *testing.T) {
	mockRepo := setupTest(t)

	// Используем Data Builder для создания категории
	category := NewCategoryBuilder().
		WithName("CategoryName").
		Build()

	mockRepo.CreateFunc = func(c *models.Category) (*models.Category, error) {
		return c, nil
	}

	createdCategory, err := mockRepo.Create(category)

	assert.Nil(t, err)
	assert.Equal(t, category, createdCategory)
}

// Пример теста с использованием Object Mother
func TestCreateCategory_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	// Исправляем вызов CustomCategory, передавая оба параметра: id и name
	category := CategoryMother.CustomCategory(0, "CategoryName")

	mockRepo.CreateFunc = func(c *models.Category) (*models.Category, error) {
		return nil, errors.New("creation failed")
	}

	createdCategory, err := mockRepo.Create(category)

	assert.Error(t, err)
	assert.Nil(t, createdCategory)
}

func TestUpdateCategory_Success(t *testing.T) {
	mockRepo := setupTest(t)
	category := NewCategoryBuilder().WithName("CategoryName").Build()

	mockRepo.UpdateFunc = func(c *models.Category) (*models.Category, error) {
		return c, nil
	}

	updatedCategory, err := mockRepo.Update(category)

	assert.NoError(t, err)
	assert.Equal(t, category, updatedCategory)
}

func TestUpdateCategory_Failure(t *testing.T) {
	mockRepo := setupTest(t)
	// Исправляем вызов CustomCategory, передавая оба параметра: id и name
	category := CategoryMother.CustomCategory(0, "CategoryName")

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
	category := NewCategoryBuilder().WithName("CategoryName").Build()
	categoryID := category.ID

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
		*NewCategoryBuilder().WithName("CategoryName1").Build(),
		*NewCategoryBuilder().WithName("CategoryName2").Build(),
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

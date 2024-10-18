package utils

//// CategoryBuilder реализует паттерн Data Builder для Category
//type CategoryBuilder struct {
//	category models.Category
//}
//
//// NewCategoryBuilder создает новый экземпляр CategoryBuilder с настройками по умолчанию
//func NewCategoryBuilder() *CategoryBuilder {
//	return &CategoryBuilder{
//		category: models.Category{
//			ID:   1,
//			Name: "DefaultCategory",
//		},
//	}
//}
//
//// WithID устанавливает ID категории
//func (b *CategoryBuilder) WithID(id int) *CategoryBuilder {
//	b.category.ID = id
//	return b
//}
//
//// WithName устанавливает имя категории
//func (b *CategoryBuilder) WithName(name string) *CategoryBuilder {
//	b.category.Name = name
//	return b
//}
//
//// Build возвращает готовый объект Category
//func (b *CategoryBuilder) Build() *models.Category {
//	return &b.category
//}
//
//// CategoryMother реализует паттерн Object Mother для Category
//var CategoryMother = struct {
//	Default        func() *models.Category
//	WithID         func(id int) *models.Category
//	WithName       func(name string) *models.Category
//	CustomCategory func(id int, name string) *models.Category
//}{
//	Default: func() *models.Category {
//		return &models.Category{
//			ID:   1,
//			Name: "DefaultCategory",
//		}
//	},
//	WithID: func(id int) *models.Category {
//		return &models.Category{
//			ID:   id,
//			Name: "CategoryWithSpecificID",
//		}
//	},
//	WithName: func(name string) *models.Category {
//		return &models.Category{
//			ID:   2,
//			Name: name,
//		}
//	},
//	CustomCategory: func(id int, name string) *models.Category {
//		return &models.Category{
//			ID:   id,
//			Name: name,
//		}
//	},
//}

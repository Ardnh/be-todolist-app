package category

import "gorm.io/gorm"

type CategoryRepository interface {
}

type CategoryRepositoryImpl struct {
	Db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &CategoryRepositoryImpl{
		Db: db,
	}
}

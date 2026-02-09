package repository

import (
	"levelearn-backend/internal/entity"

	"gorm.io/gorm"
)

type ChapterRepository interface {
	FindAll() ([]entity.Chapter, error)
	FindById(id int) (entity.Chapter, error)
	Create(chapter *entity.Chapter) error
	Update(id int, data map[string]any) error
	Delete(id int) error
}

type chapterRepository struct {
	db *gorm.DB
}

func NewChapterRepository(db *gorm.DB) ChapterRepository {
	return &chapterRepository{db}
}

func (r *chapterRepository) FindAll() ([]entity.Chapter, error) {
	var chapters []entity.Chapter

	err := r.db.Preload("Material").Preload("Assessment").Preload("Assignment").Find(&chapters).Error
	return chapters, err
}

func (r *chapterRepository) FindById(id int) (entity.Chapter, error) {
	var chapter entity.Chapter

	err := r.db.Preload("Material").Preload("Assessment").Preload("Assignment").Find(&chapter, id).Error
	return chapter, err
}

func (r *chapterRepository) Create(chapter *entity.Chapter) error {
	return r.db.Create(chapter).Error
}

func (r *chapterRepository) Update(id int, data map[string]any) error {
	return r.db.Model(&entity.Chapter{}).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *chapterRepository) Delete(id int) error {
	return r.db.Delete(&entity.Chapter{}, id).Error
}

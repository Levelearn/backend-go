package repository

import (
	"levelearn-backend/internal/entity"

	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll() ([]entity.Course, error)
	FindById(id int) (*entity.Course, error)
	Create(course *entity.Course) error
	Update(id int, data map[string]interface{}) error
	Delete(id int) error
}

type courseRepository struct {
	db *gorm.DB
}

func NewCourseRepository(db *gorm.DB) CourseRepository {
	return &courseRepository{db}
}

func (r *courseRepository) FindAll() ([]entity.Course, error) {
	var courses []entity.Course

	err := r.db.Preload("Chapters").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindById(id int) (*entity.Course, error) {
	var course entity.Course

	err := r.db.Preload("Chapters").
		First(&course, id).Error
	return &course, err
}

func (r *courseRepository) Create(course *entity.Course) error {
	return r.db.Create(course).Error
}

func (r *courseRepository) Update(id int, data map[string]interface{}) error {
	return r.db.Model(&entity.Course{}).
		Where("id = ?", id).
		Updates(data).Error
}

func (r *courseRepository) Delete(id int) error {
	return r.db.Delete(&entity.Course{}, id).Error
}

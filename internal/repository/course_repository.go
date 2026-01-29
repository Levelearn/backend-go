package repository

import (
	"levelearn-backend/internal/entity"
	"gorm.io/gorm"
)

type CourseRepository interface {
	FindAll() ([]entity.Course, error)
	FindByID(id int) (entity.Course, error)
	Create(course entity.Course) (entity.Course, error)
	Update(id int, course entity.Course) (entity.Course, error)
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
	// Preload chapters agar data related ikut terambil (Eager Loading)
	err := r.db.Preload("Chapters").Find(&courses).Error
	return courses, err
}

func (r *courseRepository) FindByID(id int) (entity.Course, error) {
	var course entity.Course
	err := r.db.Preload("Chapters").First(&course, id).Error
	return course, err
}

func (r *courseRepository) Create(course entity.Course) (entity.Course, error) {
	err := r.db.Create(&course).Error
	return course, err
}

func (r *courseRepository) Update(id int, course entity.Course) (entity.Course, error) {
	var existingCourse entity.Course
	if err := r.db.First(&existingCourse, id).Error; err != nil {
		return course, err
	}
	// Update fields
	err := r.db.Model(&existingCourse).Updates(course).Error
	return existingCourse, err
}

func (r *courseRepository) Delete(id int) error {
	return r.db.Delete(&entity.Course{}, id).Error
}
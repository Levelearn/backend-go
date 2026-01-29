package service

import (
	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/repository"
)

type CourseService interface {
	GetAllCourses() ([]entity.Course, error)
	GetCourseByID(id int) (entity.Course, error)
	CreateCourse(input entity.Course) (entity.Course, error)
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo}
}

func (s *courseService) GetAllCourses() ([]entity.Course, error) {
	return s.repo.FindAll()
}

func (s *courseService) GetCourseByID(id int) (entity.Course, error) {
	return s.repo.FindByID(id)
}

func (s *courseService) CreateCourse(input entity.Course) (entity.Course, error) {
	// Contoh Business Logic: Validasi code tidak boleh kosong
	if input.Code == "" {
		// return empty dan error custom
	}
	return s.repo.Create(input)
}
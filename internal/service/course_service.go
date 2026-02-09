package service

import (
	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/repository"
)

type CourseService interface {
	FindAll() ([]entity.Course, error)
	FindById(id int) (*entity.Course, error)
	Create(course *entity.Course) error
	Update(id int, code, name, description, image *string) error
	Delete(id int) error
}

type courseService struct {
	repo repository.CourseRepository
}

func NewCourseService(repo repository.CourseRepository) CourseService {
	return &courseService{repo}
}

func (s *courseService) FindAll() ([]entity.Course, error) {
	return s.repo.FindAll()
}

func (s *courseService) FindById(id int) (*entity.Course, error) {
	return s.repo.FindById(id)
}

func (s *courseService) Create(course *entity.Course) error {
	return s.repo.Create(course)
}

func (s *courseService) Update(id int, code, name, description, image *string) error {
	data := map[string]any{}

	if code != nil {
		data["code"] = *code
	}

	if name != nil {
		data["name"] = *name
	}

	if description != nil {
		data["description"] = *description
	}

	if image != nil {
		data["image"] = *image
	}

	if len(data) == 0 {
		return nil
	}

	return s.repo.Update(id, data)
}

func (s *courseService) Delete(id int) error {
	return s.repo.Delete(id)
}

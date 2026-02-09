package service

import (
	"fmt"
	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/repository"
)

type ChapterService interface {
	FindAll() ([]entity.Chapter, error)
	FindById(id int) (entity.Chapter, error)
	Create(chapter *entity.Chapter, courseId int) error
	Update(id int, name, description *string, level, IsCheckpoint *int) error
	Delete(id int) error
}

type chapterService struct {
	chapterRepo repository.ChapterRepository
	courseRepo  repository.CourseRepository
}

func NewChapterService(chapterRepo repository.ChapterRepository, courseRepo repository.CourseRepository) ChapterService {
	return &chapterService{chapterRepo, courseRepo}
}

func (s *chapterService) FindAll() ([]entity.Chapter, error) {
	return s.chapterRepo.FindAll()
}

func (s *chapterService) FindById(id int) (entity.Chapter, error) {
	return s.chapterRepo.FindById(id)
}

func (s *chapterService) Create(chapter *entity.Chapter, courseId int) error {
	course, err := s.courseRepo.FindById(courseId)
	if err != nil {
		return fmt.Errorf("course not found")
	}

	chapter.CourseID = course.ID

	if err := s.chapterRepo.Create(chapter); err != nil {
		return err
	}

	return nil
}

func (s *chapterService) Update(id int, name, description *string, level, IsCheckpoint *int) error {
	data := map[string]any{}

	if name != nil {
		data["name"] = *name
	}

	if description != nil {
		data["description"] = *description
	}

	if level != nil {
		data["level"] = *level
	}

	if IsCheckpoint != nil {
		data["isCheckpoint"] = *IsCheckpoint
	}

	return s.chapterRepo.Update(id, data)
}

func (s *chapterService) Delete(id int) error {
	return s.chapterRepo.Delete(id)
}

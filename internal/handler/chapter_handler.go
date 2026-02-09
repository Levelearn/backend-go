package handler

import (
	"net/http"
	"strconv"

	"levelearn-backend/internal/entity"
	"levelearn-backend/internal/service"

	"github.com/gin-gonic/gin"
)

type ChapterHandler struct {
	chapterService service.ChapterService
}

func NewChapterHandler(chapterService service.ChapterService) *ChapterHandler {
	return &ChapterHandler{chapterService}
}

func (h *ChapterHandler) FindAll(c *gin.Context) {
	chapters, err := h.chapterService.FindAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chapters)
}

func (h *ChapterHandler) FindById(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	chapter, err := h.chapterService.FindById(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, chapter)
}

type CreateChapterInput struct {
	Name         string `json:"name" binding:"required"`
	Description  string `json:"description"`
	CourseID     int    `json:"courseId" binding:"required"`
	Level        int    `json:"level"`
	IsCheckpoint int    `json:"isCheckpoint"`
}

func (h *ChapterHandler) Create(c *gin.Context) {
	var input CreateChapterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	chapter := &entity.Chapter{
		Name:         input.Name,
		Description:  input.Description,
		Level:        input.Level,
		IsCheckpoint: input.IsCheckpoint,
	}

	if err := h.chapterService.Create(chapter, input.CourseID); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, chapter)
}

type UpdateChapterInput struct {
	Name         *string `json:"name"`
	Description  *string `json:"description"`
	Level        *int    `json:"level"`
	IsCheckpoint *int    `json:"isCheckpoint"`
}

func (h *ChapterHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	var input UpdateChapterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := h.chapterService.Update(id, input.Name, input.Description, input.Level, input.IsCheckpoint); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "chapter updated"})
}

func (h *ChapterHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid id"})
		return
	}

	if err := h.chapterService.Delete(id); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "chapter deleted"})
}

package main

import (
	"levelearn-backend/config"
	"levelearn-backend/internal/handler"
	"levelearn-backend/internal/repository"
	"levelearn-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	config.RunMigration(db)

	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		courses := api.Group("/courses")
		{
			courses.GET("/", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
			courses.POST("/", courseHandler.Create)
		}
	}

	r.Run(":8080")
}

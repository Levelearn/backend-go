package main

import (
	"levelearn-backend/config"
	"levelearn-backend/internal/handler"
	"levelearn-backend/internal/middleware"
	"levelearn-backend/internal/repository"
	"levelearn-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	db := config.ConnectDB()

	config.RunMigration(db)

	// Repo
	courseRepo := repository.NewCourseRepository(db)
	userRepo := repository.NewUserRepository(db)

	// Service
	authService := service.NewAuthService(userRepo)
	courseService := service.NewCourseService(courseRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	courseHandler := handler.NewCourseHandler(courseService)

	r := gin.Default()

	api := r.Group("/api/v1")
	{
		auth := api.Group("/auth")
		{
			auth.POST("/register", authHandler.Register)
			auth.POST("/login", authHandler.Login)
		}

		protected := api.Group("/")
		protected.Use(middleware.JWTAuth())
		{
			courses := protected.Group("/courses")
			{
				courses.GET("/", courseHandler.GetAll)
				courses.GET("/:id", courseHandler.GetByID)
				courses.POST("/", courseHandler.Create)
			}
		}
	}

	r.Run(":8080")
}

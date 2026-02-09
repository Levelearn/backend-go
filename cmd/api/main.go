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
	chapterRepo := repository.NewChapterRepository(db)

	// Service
	authService := service.NewAuthService(userRepo)
	courseService := service.NewCourseService(courseRepo)
	userService := service.NewUserService(userRepo)
	chapterService := service.NewChapterService(chapterRepo, courseRepo)

	// Handler
	authHandler := handler.NewAuthHandler(authService)
	courseHandler := handler.NewCourseHandler(courseService)
	userHandler := handler.NewUserHandler(userService)
	chapterHandler := handler.NewChapterHandler(chapterService)

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
			users := protected.Group("/users")
			{
				users.GET("/me", userHandler.GetMe)
				users.PUT("/me", userHandler.UpdateMe)
				users.GET("/:id", userHandler.GetById)
			}

			courses := protected.Group("/courses")
			{
				courses.GET("", courseHandler.FindAll)
				courses.GET("/:id", courseHandler.FindById)
				courses.POST("", courseHandler.Create)
				courses.PUT("/:id", courseHandler.Update)
				courses.DELETE("/:id", courseHandler.Delete)
			}

			chapters := protected.Group("/chapters")
			{
				chapters.GET("", chapterHandler.FindAll)
				chapters.GET("/:id", chapterHandler.FindById)
				chapters.POST("", chapterHandler.Create)
				chapters.POST("/:id", chapterHandler.Update)
				chapters.POST("/:id", chapterHandler.Delete)
			}
		}
	}

	r.Run(":8080")
}

package main

import (
	"levelearn-backend/config"
	"levelearn-backend/internal/handler"
	"levelearn-backend/internal/repository"
	"levelearn-backend/internal/service"

	"github.com/gin-gonic/gin"
)

func main() {
	// 1. Setup Database
	db := config.ConnectDB()

	// 2. Setup Repository, Service, Handler (Dependency Injection)
	// --- Course Module ---
	courseRepo := repository.NewCourseRepository(db)
	courseService := service.NewCourseService(courseRepo)
	courseHandler := handler.NewCourseHandler(courseService)

	// --- User Module (Anda buat sendiri mengikuti pola di atas) ---
	// userRepo := repository.NewUserRepository(db) ...

	// 3. Setup Router
	r := gin.Default()

	api := r.Group("/api/v1")
	{
		courses := api.Group("/courses")
		{
			courses.GET("/", courseHandler.GetAll)
			courses.GET("/:id", courseHandler.GetByID)
			courses.POST("/", courseHandler.Create)
		}

		// Setup route lain disini:
		// users := api.Group("/users") ...
	}

	r.Run(":8080") // Jalan di port 8080
}
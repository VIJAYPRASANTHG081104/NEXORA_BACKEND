package router

import (
	"database/sql"
	"nexora_backend/api/middleware"
	"nexora_backend/internal/users"
	"nexora_backend/internal/videos"

	"github.com/gin-gonic/gin"
)

func InitializeRouter(db *sql.DB, router *gin.Engine) {
	userStore := users.NewUserStore(db)
	userHandler := users.CreateUserHandler(userStore)
	userHandler.RegisterRoutes(router)

	// video service
	// videoService := &VideoService{}
	api := router.Group("/api")
	videoHandler := videos.CreateVideoServiceHandler(userStore)
	api.Use(middleware.AuthMiddleware()) // Example middleware
	videoHandler.VideoServiceRouter(api)
}

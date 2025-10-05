package router

import (
	"database/sql"
	"nexora_backend/internal/users"

	"github.com/gorilla/mux"
)

func InitializeRouter(db *sql.DB, router *mux.Router) {
	userStore := users.NewUserStore(db)
	userHandler := users.NewHandler(userStore)
	userHandler.RegisterRoutes(router)
}

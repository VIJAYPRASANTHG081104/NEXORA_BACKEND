package server

import (
	"database/sql"
	"log"
	"net/http"
	"nexora_backend/api/router"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

type APIServer struct{
	addr string
	db *sql.DB
}

func NewAPIServer(addr string,db *sql.DB) *APIServer{
	return &APIServer{
		addr: addr,
		db:db,
	}
}

func (s *APIServer) Run() error{
	// Create a router
	initialRouter := gin.Default();
	router.InitializeRouter(s.db,initialRouter)
	// cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string {"http://localhost:3000"},
		AllowedMethods: []string {"GET","POST","DELETE","OPTIONS","PUT","PATCH"},
		AllowedHeaders: []string {"Authorization","Content-type"},
		AllowCredentials: true,
	})

	corsHandle := c.Handler(initialRouter);

	log.Println("Listening on port", s.addr);
	return http.ListenAndServe(s.addr, corsHandle);
}
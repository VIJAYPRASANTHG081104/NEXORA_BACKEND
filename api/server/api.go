package server

import (
	"database/sql"
	"log"
	"net/http"
	"nexora_backend/api/router"

	"github.com/gorilla/mux"
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
	muxRouter := mux.NewRouter()
	
	router.InitializeRouter(s.db,muxRouter)
	// cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string {"http://localhost:3000"},
		AllowedMethods: []string {"GET","POST","DELETE","OPTIONS","PUT","PATCH"},
		AllowedHeaders: []string {"Authorization","Content-type"},
		AllowCredentials: true,
	})

	corsHandle := c.Handler(muxRouter);

	log.Println("Listening on port", s.addr);
	return http.ListenAndServe(s.addr, corsHandle);
}
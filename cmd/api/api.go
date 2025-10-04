package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type APIServer struct{
	addr string
}

func NewAPIServer(addr string) *APIServer{
	return &APIServer{
		addr: addr,
	}
}

func (s *APIServer) Run() error{
	// Create a router
	router := mux.NewRouter()
	
	// cors
	c := cors.New(cors.Options{
		AllowedOrigins: []string {"http://localhost:3000"},
		AllowedMethods: []string {"GET","POST","DELETE","OPTIONS","PUT","PATCH"},
		AllowedHeaders: []string {"Authorization","Content-type"},
		AllowCredentials: true,
	})

	corsHandle := c.Handler(router);

	log.Println("Listening on port", s.addr);
	return http.ListenAndServe(s.addr, corsHandle);
}
package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"

	"github.com/dodo/ecom/service/user"
)

type APIServer struct {
	Addrs string
	DB    *gorm.DB
}

func NewAPIServer(addrs string, db *gorm.DB) *APIServer {
	return &APIServer{
		Addrs: addrs,
		DB:    db,
	}
}

func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	userStore := user.NewStore(s.DB)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Println("Listening on PORT", s.Addrs)

	return http.ListenAndServe(s.Addrs, router)
}

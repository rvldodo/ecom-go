package api

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"

	_ "github.com/dodo/ecom/cmd/docs"
	"github.com/dodo/ecom/config"
	"github.com/dodo/ecom/service/cart"
	"github.com/dodo/ecom/service/order"
	"github.com/dodo/ecom/service/product"
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

// @title Swagger Example API
// @version 1.0
// @description This is a sample server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:1337
// @BasePath /api/v1
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()
	subrouter.Handle("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("http://localhost:1337/swagger/swagger.json"),
	)).Methods("GET")

	userStore := user.NewStore(s.DB)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	productStore := product.NewStore(s.DB)
	productHandler := product.NewHandler(productStore)
	productHandler.RegisterRoutes(subrouter)

	orderStore := order.NewStore(s.DB)

	cartHandler := cart.NewHandler(orderStore, productStore, userStore)
	cartHandler.RegisterRoutes(subrouter)

	log.Println("Listening on PORT", s.Addrs)
	log.Println(config.Envs.SwaggerAddress)

	return http.ListenAndServe(s.Addrs, router)
}

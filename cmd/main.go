package main

import (
	"log"

	"github.com/dodo/ecom/cmd/api"
	"github.com/dodo/ecom/config"
	"github.com/dodo/ecom/db"
)

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
func main() {
	// run server
	server := api.NewAPIServer(config.Envs.Port, db.DB)
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

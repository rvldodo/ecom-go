package main

import (
	"log"

	"github.com/dodo/ecom/cmd/api"
	"github.com/dodo/ecom/config"
	"github.com/dodo/ecom/db"
)

func main() {
	// run server
	server := api.NewAPIServer(config.Envs.Port, db.DB)
	if err := server.Run(); err != nil {
		log.Fatal(err.Error())
	}
}

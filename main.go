package main

import (
	"github.com/darcops/dialgorithm-server/models"
	"github.com/darcops/dialgorithm-server/routes"
)

func main() {
	models.Connect()
	models.Migrate()
	routes.Init()
}

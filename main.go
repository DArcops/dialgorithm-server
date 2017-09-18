package main

import (
	"github.com/darcops/dialgorithm-server/models"
	"github.com/darcops/dialgorithm-server/routes"
)

func main() {
	routes.Init()
	models.Connect()
}

package main

import (
	"fmt"

	"github.com/darcops/dialgorithm-server/models"
	"github.com/darcops/dialgorithm-server/modules/encrypt"
	"github.com/darcops/dialgorithm-server/routes"
)

func main() {
	models.Connect()
	models.Migrate()
	message := "que pedal"
	fmt.Println(message)
	e, _ := encrypt.Encrypt([]byte(message))
	fmt.Println(e)
	de, _ := encrypt.Decrypt(e)
	fmt.Println("ahora...", string(de))
	routes.Init()
}

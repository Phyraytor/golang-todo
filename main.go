package main

import (
	"fmt"
	"./packages/handler"
	"log"
)

func main() {
	handlers := new(handler.Handler)

	runSever := new(server.Server)
	if err := runSever.Run("8001"); err != nil {
		log.Fatalf("Error: %s", err.Error())
	}

}
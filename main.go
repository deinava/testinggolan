package main

import (
	"log"

	"github.com/deinava/testinggolan/bd"
	"github.com/deinava/testinggolan/handlers"
)

func main() {
	if bd.ChequeoConnection() == 0 {
		log.Fatal("Sin conexion a la base de datos")
		return
	}
	handlers.Manejadores()
}

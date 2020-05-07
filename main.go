package main

import (
	"log"

	"github.com/BrangyCastro/twitter-backend/bd"
	"github.com/BrangyCastro/twitter-backend/handlers"
)

func main() {
	if bd.ChequeoConexion() == 0 {
		log.Fatal("Sin conecxion a la BD")
		return
	}
	handlers.Manejadores()
}

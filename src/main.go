package main

import (
	"log"
	"net/http"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/routes"
)

func main() {
	go database.ConectarBanco()
	go routes.SetupAPI()
	log.Println("Servidor rodando na porta 8080...")

	// ESTA LINHA É O QUE SEGURA O CONTAINER VIVO:
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

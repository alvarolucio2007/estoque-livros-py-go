package main

import (
	"log"
	"net/http"
)

func main() {
	go ConectarBanco()
	go setupAPI()
	log.Println("Servidor rodando na porta 8080...")

	// ESTA LINHA É O QUE SEGURA O CONTAINER VIVO:
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

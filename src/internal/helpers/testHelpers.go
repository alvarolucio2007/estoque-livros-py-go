// Package helpers ajuda com funções de auxílio
package helpers

import (
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
)

var Livro = models.Livro{Titulo: "Livro de Teste", Autor: "Autor Teste", Preco: 120, Quantidade: 100, Ano: 1984}

func CriarLivroTesteAux(t *testing.T) models.Livro {
	if err := database.DB.Create(&Livro).Error; err != nil {
		t.Fatalf("Erro ao criar livro teste: %v", err)
	}

	// A partir daqui, qualquer teste que chamar essa função
	// será limpo automaticamente no final.
	t.Cleanup(func() {
		database.DB.Unscoped().Delete(&Livro)
	})

	return Livro
}

package main

import "testing"

func TestVerificadorCompleto(t *testing.T) {
}

var mapLivroTeste = map[string]Livro{
	"Normal":       {Titulo: "Normal", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20},
	"Sem titulo":   {Titulo: "", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20},
	"Sem Autor":    {Titulo: "Sem autor", Autor: "", Preco: 100, Ano: 1984, Quantidade: 20},
	"Preço errado": {Titulo: "Preço errado", Autor: "Preço errado", Preco: -100, Ano: 1984, Quantidade: 20},
	"Ano errado":   {Titulo: "Ano errado", Autor: "Ano errado", Preco: 100, Ano: 9000, Quantidade: 20},
}

func testarServicoAdicionarLivro(t *testing.T) {
	for nomeLivro, livro := range mapLivroTeste {

		err := servicoAdicionarLivro(livro.Titulo, livro.Autor, livro.Preco, livro.Ano, livro.Quantidade)
		if err != nil {
			t.Errorf("Livro %v  não salvo! Erro: %v ", nomeLivro, err)
		}
	}
}

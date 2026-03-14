package main

import "testing"

func TestVerifiers(t *testing.T) {
	t.Run("verifierAdicionar", func(t *testing.T) {
		testarServicoAdicionarLivro(t)
	})
	t.Run("verifierDeletar", func(t *testing.T) {
		servicoAdicionarLivro("AA", "BB", 100, 2007, 20)
		testarServicoDeletarLivro(t)
	})
}

var mapLivroTeste = map[string]struct {
	livro      Livro
	esperaErro bool
}{
	"Normal":       {livro: Livro{Titulo: "Normal", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: false},
	"Sem titulo":   {livro: Livro{Titulo: "", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Sem Autor":    {livro: Livro{Titulo: "Sem autor", Autor: "", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Preço errado": {livro: Livro{Titulo: "Preço errado", Autor: "Preço errado", Preco: -100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Ano errado":   {livro: Livro{Titulo: "Ano errado", Autor: "Ano errado", Preco: 100, Ano: 9000, Quantidade: 20}, esperaErro: true},
}

func testarServicoAdicionarLivro(t *testing.T) {
	for nome, tc := range mapLivroTeste {
		t.Run(nome, func(t *testing.T) {
			err := servicoAdicionarLivro(tc.livro.Titulo, tc.livro.Autor, tc.livro.Preco, tc.livro.Ano, tc.livro.Quantidade)
			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapIDTeste = map[string]struct {
	ID         uint
	esperaErro bool
}{
	"Não retorna erro":    {ID: 1, esperaErro: false},
	"Retorna erro":        {ID: 0, esperaErro: true},
	"Não existe no banco": {ID: 1225, esperaErro: true},
}

func testarServicoDeletarLivro(t *testing.T) {
	livroParaDeletar := Livro{Titulo: "Livro de Teste", Autor: "Autor Teste"}
	DB.Create(&livroParaDeletar)
	mapIDTeste["Não retorna erro"] = struct {
		ID         uint
		esperaErro bool
	}{ID: livroParaDeletar.ID, esperaErro: false}
	for nome, tc := range mapIDTeste {
		t.Run(nome, func(t *testing.T) {
			err := servicoDeletarLivro(tc.ID)
			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapLivroEditarTeste = map[string]struct {
	id         uint
	livro      Livro
	esperaErro bool
}{
	"Normal": {id: 1, livro: Livro{Livro{Titulo: "Normal", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: false}},
	"Sem titulo":   {id: 2, livro: Livro{Titulo: "", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Sem Autor":    {id: 3, livro: Livro{Titulo: "Sem autor", Autor: "", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Preço errado": {id: 4, livro: Livro{Titulo: "Preço errado", Autor: "Preço errado", Preco: -100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Ano errado":   {id: 5, livro: Livro{Titulo: "Ano errado", Autor: "Ano errado", Preco: 100, Ano: 9000, Quantidade: 20}, esperaErro: true},
	},

func testarServicoAtualizarLivro(t *testing.T) {
}

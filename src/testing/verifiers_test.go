package test

import (
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/verifiers"
)

func TestVerifiers(t *testing.T) {
	t.Run("verifierAdicionar", func(t *testing.T) {
		testarServicoAdicionarLivro(t)
	})
	t.Run("verifierDeletar", func(t *testing.T) {
		testarServicoDeletarLivro(t)
	})
	t.Run("verifierEditar", func(t *testing.T) {
		testarServicoAtualizarLivro(t)
	})
	t.Run("verifierBuscarTitulo", func(t *testing.T) {
		testarServicoBuscarLivroTitulo(t)
	})
	t.Run("verifierBuscarAutor", func(t *testing.T) {
		testarServicoBuscarLivroAutor(t)
	})
	t.Run("verifierBuscarID", func(t *testing.T) {
		testarServicoBuscarLivroID(t)
	})
}

func criarLivroTeste(t *testing.T) models.Livro {
	livro := models.Livro{Titulo: "Livro de Teste", Autor: "Autor Teste", Preco: 120, Quantidade: 100, Ano: 1984}
	if err := database.DB.Create(&livro).Error; err != nil {
		t.Fatalf("Erro ao criar livro teste: %v", err)
	}
	return livro
}

var mapLivroTeste = map[string]struct {
	livro      models.Livro
	esperaErro bool
}{
	"Normal":       {livro: models.Livro{Titulo: "Normal", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: false},
	"Sem titulo":   {livro: models.Livro{Titulo: "", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Sem Autor":    {livro: models.Livro{Titulo: "Sem autor", Autor: "", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Preço errado": {livro: models.Livro{Titulo: "Preço errado", Autor: "Preço errado", Preco: -100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Ano errado":   {livro: models.Livro{Titulo: "Ano errado", Autor: "Ano errado", Preco: 100, Ano: 9000, Quantidade: 20}, esperaErro: true},
}

func testarServicoAdicionarLivro(t *testing.T) {
	criarLivroTeste(t)

	for nome, tc := range mapLivroTeste {
		t.Run(nome, func(t *testing.T) {
			err := verifiers.ServicoAdicionarLivro(tc.livro.Titulo, tc.livro.Autor, tc.livro.Preco, tc.livro.Ano, tc.livro.Quantidade)
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
	livroParaDeletar := criarLivroTeste(t)
	mapIDTeste["Não retorna erro"] = struct {
		ID         uint
		esperaErro bool
	}{ID: livroParaDeletar.ID, esperaErro: false}
	for nome, tc := range mapIDTeste {
		t.Run(nome, func(t *testing.T) {
			err := verifiers.ServicoDeletarLivro(tc.ID)
			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapLivroEditarTeste = map[string]struct {
	id         uint
	livro      models.Livro
	esperaErro bool
}{
	"Normal":       {id: 1, livro: models.Livro{Titulo: "Normal", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: false},
	"Sem titulo":   {id: 2, livro: models.Livro{Titulo: "", Autor: "Normal", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Sem Autor":    {id: 3, livro: models.Livro{Titulo: "Sem autor", Autor: "", Preco: 100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Preço errado": {id: 4, livro: models.Livro{Titulo: "Preço errado", Autor: "Preço errado", Preco: -100, Ano: 1984, Quantidade: 20}, esperaErro: true},
	"Ano errado":   {id: 5, livro: models.Livro{Titulo: "Ano errado", Autor: "Ano errado", Preco: 100, Ano: 9000, Quantidade: 20}, esperaErro: true},
}

func testarServicoAtualizarLivro(t *testing.T) {
	livroParaEditar := criarLivroTeste(t)
	for nome, tc := range mapLivroEditarTeste {
		t.Run(nome, func(t *testing.T) {
			err := verifiers.ServicoAtualizarLivro(livroParaEditar.ID, tc.livro)
			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapLivroTituloBuscar = map[string]struct {
	Título     string
	esperaErro bool
}{
	"Normal":        {Título: "Livro de Teste", esperaErro: false},
	"Título vazio":  {Título: "", esperaErro: true},
	"Título errado": {Título: "AAAAAAAAAAA", esperaErro: true},
}

func testarServicoBuscarLivroTitulo(t *testing.T) {
	criarLivroTeste(t)
	for nome, tc := range mapLivroTituloBuscar {
		t.Run(nome, func(t *testing.T) {
			_, err := verifiers.ServicoBuscarLivroTitulo(tc.Título)

			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapLivroAutorBuscar = map[string]struct {
	Autor      string
	esperaErro bool
}{
	"Normal":        {Autor: "Livro de Teste", esperaErro: false},
	"Título vazio":  {Autor: "", esperaErro: true},
	"Título errado": {Autor: "AAAAAAAAAAA", esperaErro: true},
}

func testarServicoBuscarLivroAutor(t *testing.T) {
	criarLivroTeste(t)
	for nome, tc := range mapLivroAutorBuscar {
		t.Run(nome, func(t *testing.T) {
			_, err := verifiers.ServicoBuscarLivroTitulo(tc.Autor)

			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v",
					nome, err, tc.esperaErro)
			}
		})
	}
}

var mapLivroIDBuscar = map[string]struct {
	ID         uint
	esperaErro bool
}{
	"Título vazio":  {ID: 0, esperaErro: true},
	"Título errado": {ID: 12, esperaErro: true},
}

func testarServicoBuscarLivroID(t *testing.T) {
	livro := criarLivroTeste(t)
	mapLivroIDBuscar["Normal"] = struct {
		ID         uint
		esperaErro bool
	}{ID: livro.ID, esperaErro: false}
	for nome, tc := range mapLivroIDBuscar {
		t.Run(nome, func(t *testing.T) {
			_, err := verifiers.ServicoBuscarLivroID(tc.ID)
			if (err != nil) != tc.esperaErro {
				t.Errorf("[%s] Resultado inesperado: erro recebido = %v, esperava erro? %v ",
					nome, err, tc.esperaErro)
			}
		})
	}
}

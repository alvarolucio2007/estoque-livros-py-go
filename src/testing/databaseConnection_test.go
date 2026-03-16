package test

import (
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
)

var livroExemplo = models.Livro{
	Titulo:     "Teste",
	Autor:      "autorTeste",
	Preco:      999,
	Ano:        1500,
	Quantidade: 150,
}

func TestDatabaseConnection(t *testing.T) {
	t.Run("Adicionar", func(t *testing.T) {
		testarAdicionarLivro(t)
	})
	t.Run("Carregar", func(t *testing.T) {
		testarCarregarDados(t)
	})
	t.Run("ChecarID", func(t *testing.T) {
		testarListarID(t)
	})
	t.Run("Editar", func(t *testing.T) {
		testarEditarLivro(t)
	})
	t.Run("Buscar", func(t *testing.T) {
		testarBuscarLivroTitulo(t)
		testarBuscarLivroAutor(t)
	})
	t.Run("Relatório", func(t *testing.T) {
		testarRelatorio(t)
	})
	t.Run("Deletar", func(t *testing.T) {
		testarDeletarLivro(t)
	})
}

func testarAdicionarLivro(t *testing.T) {
	disponivel := true
	livroExemplo.Disponivel = &disponivel
	err := database.AdicionarLivro(livroExemplo)
	if err != nil {
		t.Errorf("Livro não salvo: %v", err)
	}
}

func testarCarregarDados(t *testing.T) {
	listaLivros, err := database.CarregarDados()
	if err != nil {
		t.Fatalf("Erro ao carregar dados: %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("A lista está vazia, esperava 1 livro.")
	}
	recebido := listaLivros[0]
	esperado := livroExemplo
	recebido.ID = 0
	if *recebido.Disponivel != *esperado.Disponivel {
		t.Error("Disponibilidade errada")
	}
	recebido.Disponivel = esperado.Disponivel
	if recebido != esperado {
		t.Errorf("Ainda há diferenças. \nRecebido: %+v\nEsperado: %+v", recebido, esperado)
	}
}

func testarEditarLivro(t *testing.T) {
	disponivel := true
	livroEdicao := models.Livro{Titulo: "Teste2", Autor: "autorTeste2", Preco: 9992, Ano: 1502, Quantidade: 152, Disponivel: &disponivel}
	err := database.AtualizarLivro(1, livroEdicao)
	if err != nil {
		t.Fatalf("log: erro ao atualizar livro, %v", err)
	}
	listaLivros, err := database.CarregarDados()
	if err != nil || len(listaLivros) == 0 {
		t.Fatal("log: erro ao carregar dados após edição")
	}
	recebido := listaLivros[0]
	if livroEdicao.Titulo != recebido.Titulo {
		t.Error("log: titulo nao editado")
	}
	if livroEdicao.Autor != recebido.Autor {
		t.Error("log: autor nao editado")
	}
	if livroEdicao.Ano != recebido.Ano {
		t.Error("log: ano não editado")
	}
	if livroEdicao.Quantidade != recebido.Quantidade {
		t.Error("log: quantidade não editada")
	}
	if *livroEdicao.Disponivel != *recebido.Disponivel {
		t.Error("log: disponivel não editado")
	}
	livroExemplo = livroEdicao
}

func testarListarID(t *testing.T) {
	listaID := database.ListarID()
	if len(listaID) != 1 {
		t.Error("log: quantidade de IDs inesperada, esperado: %i, atual: %i", 1, len(listaID))
	}
}

func testarBuscarLivroTitulo(t *testing.T) {
	listaLivros, err := database.BuscarLivroTitulo(livroExemplo.Titulo)
	if err != nil {
		t.Fatalf("log: erro interno na função buscarLivroTitulo: %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("log: busca não encontrou nenhum livro")
	}
	if listaLivros[0].Titulo != livroExemplo.Titulo {
		t.Error("log: busca errada por título")
	}
}

func testarBuscarLivroAutor(t *testing.T) {
	listaLivros, err := database.BuscarLivroAutor(livroExemplo.Autor)
	if err != nil {
		t.Fatalf("log: erro interno na função buscarLivroAutor %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("log: busca nao encontrou livro algum")
	}
	if listaLivros[0].Autor != livroExemplo.Autor {
		t.Errorf("log: busca errada por autor\nEsperado: %+v\nRecebido: %+v", livroExemplo, listaLivros[0])
	}
}

func testarDeletarLivro(t *testing.T) {
	err := database.DeletarLivro(1)
	if err != nil {
		t.Fatalf("Erro interno: %v", err)
	}
	listaLivros, _ := database.CarregarDados()
	if len(listaLivros) != 0 {
		t.Error("Livro não foi apagado!")
	}
}

func testarRelatorio(t *testing.T) {
	relatorio, err := database.GerarRelatorio()
	if err != nil {
		t.Fatalf("Erro ao gerar relatório: %v", err)
	}
	if relatorio["total_livros"].(int64) != 1 {
		t.Error("log: relatório não retornou 1 total_livros")
	}
	if relatorio["livros_disponiveis"].(int64) != 1 {
		t.Error("log: relatório não retornou 1 livros_disponiveis")
	}
	if relatorio["livros_indisponiveis"].(int64) != 0 {
		t.Error("log: relatório não retornou 0 lívros_indisponiveis")
	}
	if relatorio["valor_total_estoque"].(float64) != 1518784 {
		t.Errorf("Valor total esperado: 1518784, valor dado: %.2f", relatorio["valor_total_estoque"])
	}
}

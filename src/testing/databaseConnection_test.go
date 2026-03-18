package test

import (
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/helpers"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
)

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
	livroExemplo := helpers.Livro
	disponivel := true
	livroExemplo.Disponivel = &disponivel
	err := database.AdicionarLivro(livroExemplo)
	if err != nil {
		t.Errorf("Livro não salvo: %v", err)
	}
	t.Cleanup(func() {
		database.DB.Unscoped().Delete(&livroExemplo)
	})
}

func testarCarregarDados(t *testing.T) {
	esperado := helpers.CriarLivroTesteAux(t)
	listaLivros, err := database.CarregarDados()
	if err != nil {
		t.Fatalf("Erro ao carregar dados: %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("A lista está vazia, esperava 1 livro.")
	}
	recebido := listaLivros[0]
	if recebido.Disponivel == nil || esperado.Disponivel == nil {
		t.Fatal("Erro: campo Disponivel veio nulo (nil)")
	}
	if *recebido.Disponivel != *esperado.Disponivel {
		t.Error("Disponibilidade errada")
	}
	recebido.ID = esperado.ID
	recebido.Disponivel = esperado.Disponivel
	if recebido != esperado {
		t.Errorf("Ainda há diferenças. \nRecebido: %+v\nEsperado: %+v", recebido, esperado)
	}
}

func testarEditarLivro(t *testing.T) {
	livroCriado := helpers.CriarLivroTesteAux(t)
	disponivel := true
	livroEdicao := models.Livro{Titulo: "Teste2", Autor: "autorTeste2", Preco: 9992, Ano: 1502, Quantidade: 152, Disponivel: &disponivel}
	err := database.AtualizarLivro(livroCriado.ID, livroEdicao)
	if err != nil {
		t.Fatalf("log: erro ao atualizar livro, %v", err)
	}
	listaLivros, err := database.CarregarDados()
	if err != nil || len(listaLivros) == 0 {
		t.Fatal("log: erro ao carregar dados após edição")
	}
	recebido, _ := database.BuscarPorID(livroCriado.ID)
	if recebido == nil {
		t.Fatalf("log: livro a ser editado não está na db")
	}
	if livroEdicao.Titulo != recebido.Titulo {
		t.Errorf("log: titulo nao editado, esperado: %v, recebido: %v", livroEdicao.Titulo, recebido.Titulo)
	}
	if livroEdicao.Autor != recebido.Autor {
		t.Errorf("log: autor nao editado,esperado: %v,recebido: %v", livroEdicao.Autor, recebido.Autor)
	}
	if livroEdicao.Ano != recebido.Ano {
		t.Errorf("log: ano não editado,,esperado: %v,recebido: %v", livroEdicao.Ano, recebido.Ano)
	}
	if livroEdicao.Quantidade != recebido.Quantidade {
		t.Errorf("log: quantidade não editada,esperado: %v,recebido: %v", livroEdicao.Quantidade, recebido.Quantidade)
	}
	if *livroEdicao.Disponivel != *recebido.Disponivel {
		t.Errorf("log: disponivel não editado,esperado: %v,recebido: %v", livroEdicao.Disponivel, recebido.Disponivel)
	}
}

func testarListarID(t *testing.T) {
	listaID, err := database.ListarID()
	if len(listaID) != 1 {
		t.Error("log: quantidade de IDs inesperada, esperado: %i, atual: %i", 1, len(listaID))
	}
	if err != nil {
		t.Errorf("log: erro interno: %v", err)
	}
}

func testarBuscarLivroTitulo(t *testing.T) {
	helpers.CriarLivroTesteAux(t)
	listaLivros, err := database.BuscarLivroTitulo("Livro de Teste")
	if err != nil {
		t.Fatalf("log: erro interno na função buscarLivroTitulo: %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("log: busca não encontrou nenhum livro")
	}
	if listaLivros[0].Titulo != helpers.Livro.Titulo {
		t.Error("log: busca errada por título")
	}
}

func testarBuscarLivroAutor(t *testing.T) {
	helpers.CriarLivroTesteAux(t)
	listaLivros, err := database.BuscarLivroAutor("Autor Teste")
	if err != nil {
		t.Fatalf("log: erro interno na função buscarLivroAutor %v", err)
	}
	if len(listaLivros) == 0 {
		t.Fatal("log: busca nao encontrou livro algum")
	}
	if listaLivros[0].Autor != helpers.Livro.Autor {
		t.Errorf("log: busca errada por autor\nEsperado: %+v\nRecebido: %+v", helpers.Livro, listaLivros[0])
	}
}

func testarDeletarLivro(t *testing.T) {
	livroCriado := helpers.CriarLivroTesteAux(t)

	listaLivrosAntesDeletar, _ := database.CarregarDados()
	err := database.DeletarLivro(livroCriado.ID)
	if err != nil {
		t.Fatalf("Erro interno: %v", err)
	}
	listaLivrosAposDeletar, _ := database.CarregarDados()
	if len(listaLivrosAposDeletar) != len(listaLivrosAntesDeletar)-1 {
		t.Error("Livro não foi apagado!")
	}
	if recebido, err := database.BuscarPorID(livroCriado.ID); err == nil && recebido != nil {
		t.Error("log: o livro ainda existe.")
	} else if err != nil && err.Error() != "nenhum livro encontrado" {
		t.Errorf("log: erro na busca por ID para checagem, erro: %v", err)
	}
}

func testarRelatorio(t *testing.T) {
	helpers.CriarLivroTesteAux(t)
	relatorio, err := database.GerarRelatorio()
	if err != nil {
		t.Fatalf("Erro ao gerar relatório: %v", err)
	}
	listaLivros, _ := database.CarregarDados()

	if relatorio["total_livros"].(int64) != int64(len(listaLivros)) {
		t.Error("log: relatório não retornou 1 total_livros")
	}

	if relatorio["livros_disponiveis"].(int64) != int64(len(listaLivros)) {
		t.Error("log: relatório não retornou 1 livros_disponiveis")
	}
	if relatorio["livros_indisponiveis"].(int64) != 0 {
		t.Error("log: relatório não retornou 0 lívros_indisponiveis")
	}
	somaListaLivros := 0.0
	for _, livro := range listaLivros {
		somaListaLivros += float64(livro.Quantidade) * livro.Preco
	}
	if relatorio["valor_total_estoque"].(float64) != somaListaLivros {
		t.Errorf("Valor total esperado: 12000, valor dado: %.2f", relatorio["valor_total_estoque"])
	}
}

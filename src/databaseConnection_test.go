package main

import (
	"context"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestIntegracaoPostgres(t *testing.T) {
	ctx := context.Background()
	dbName := "db_estoque_test"
	dbUser := "user"
	dbPass := "password"
	container, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:18-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPass),
		testcontainers.WithWaitStrategy(wait.ForLog("database system is ready to accept connections").WithOccurrence(2).WithStartupTimeout(5*time.Second)),
	)
	if err != nil {
		t.Fatal("Falha ao subir o container:", err)
	}
	defer func() {
		if err := container.Terminate(ctx); err != nil {
			t.Fatalf("Falha ao parar o container: %s", err)
		}
	}()
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		t.Fatal(err)
	}
	t.Run("Testar conexão e create", func(t *testing.T) {
		os.Setenv("DATABASE_URL", connStr)
		ConectarBanco()
		DB.AutoMigrate(&Livro{})
		disponivel := true
		livroTeste := Livro{Titulo: "Teste", Autor: "autorTeste", Preco: 999, Ano: 1500, Quantidade: 150, Disponivel: &disponivel}
		err := adicionarLivro(livroTeste)
		if err != nil {
			t.Error("Livro não salvo")
		}
		listaLivros, err := carregarDados()
		recebido := listaLivros[0]
		esperado := livroTeste
		recebido.ID = 0
		if *recebido.Disponivel != *esperado.Disponivel {
			t.Error("Disponibilidade errada")
		}
		recebido.Disponivel = esperado.Disponivel
		if recebido != esperado {
			t.Errorf("Ainda há diferenças! Recebido %+v, Esperado: %+v", recebido, esperado)
		}
		if err != nil {
			t.Error("log: erro interno ao carregar dados:", err)
		}
		if len(listaLivros) == 0 {
			t.Error("A lista está vazia, esperava 1 livro.")
		}

		livroTeste = Livro{Titulo: "Teste2", Autor: "autorTeste2", Preco: 9992, Ano: 1502, Quantidade: 152, Disponivel: &disponivel}
		err = atualizarLivro(1, livroTeste)
		if err != nil {
			t.Errorf("log: erro ao atualizar livro, %v", err)
		}
		listaLivros, _ = carregarDados()
		recebido = listaLivros[0]
		esperado = livroTeste
		if esperado.Titulo != recebido.Titulo {
			t.Error("log: titulo nao editado")
		}
		if esperado.Autor != recebido.Autor {
			t.Error("log: autor nao editado")
		}
		if esperado.Ano != recebido.Ano {
			t.Error("log: ano não editado")
		}
		if esperado.Quantidade != recebido.Quantidade {
			t.Error("log: quantidade não editada")
		}
		listaID := listarID()
		if len(listaID) != 1 {
			t.Error("log: quantidade de IDs inesperada, esperado: %i, atual: %i", 1, len(listaID))
		}
		listaLivros, err = buscarLivroTitulo(recebido.Titulo)
		if err != nil {
			t.Error("log: erro interno na função buscarLivroTitulo")
		}
		if listaLivros[0].Titulo != recebido.Titulo {
			t.Error("log: busca errada por título")
		}
		listaLivros, err = buscarLivroAutor(recebido.Autor)
		if err != nil {
			t.Errorf("log: erro interno na função buscarLivroAutor %v", err)
		}
		if listaLivros[0].Titulo != recebido.Titulo {
			t.Error("log: busca errada por autor")
		}

		err = deletarLivro(1)
		if err != nil {
			t.Error("Erro interno!")
		}
		listaLivros, _ = carregarDados()
		if len(listaLivros) != 0 {
			t.Error("Livro não foi apagado!")
		}
	})
}

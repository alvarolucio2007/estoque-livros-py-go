package main

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	"github.com/testcontainers/testcontainers-go"
	"github.com/testcontainers/testcontainers-go/modules/postgres"
	"github.com/testcontainers/testcontainers-go/wait"
)

func TestMain(m *testing.M) {
	ctx := context.Background()
	dbName := "db_estoque_test"
	dbUser := "user"
	dbPass := "password"

	// 1. Sobe o container (Aumentei o timeout para 10s para dar fôlego ao Docker)
	container, err := postgres.RunContainer(ctx, testcontainers.WithImage("postgres:18-alpine"),
		postgres.WithDatabase(dbName),
		postgres.WithUsername(dbUser),
		postgres.WithPassword(dbPass),
		testcontainers.WithWaitStrategy(
			wait.ForLog("database system is ready to accept connections").
				WithOccurrence(2).
				WithStartupTimeout(10*time.Second),
		),
	)
	if err != nil {
		log.Fatal("Falha ao subir o container:", err)
	}

	// 2. Pega a string de conexão
	connStr, err := container.ConnectionString(ctx, "sslmode=disable")
	if err != nil {
		log.Fatal(err)
	}

	// --- IMPORTANTE: CONFIGURAR ANTES DO RUN ---

	// 3. Define a variável de ambiente
	err = os.Setenv("DATABASE_URL", connStr)
	if err != nil {
		log.Fatal(err)
	}

	// 4. Conecta o seu banco global e cria as tabelas
	ConectarBanco()
	err = DB.AutoMigrate(&Livro{})
	if err != nil {
		log.Fatal(err)
	}

	// 5. AGORA SIM, RODA OS TESTES
	// Ele vai procurar em todos os outros arquivos _test.go
	code := m.Run()

	// 6. LIMPEZA MANUAL (Melhor que o defer aqui)
	if err := container.Terminate(ctx); err != nil {
		log.Printf("Erro ao parar o container: %s", err)
	}

	// 7. FINALIZAÇÃO
	os.Exit(code)
}

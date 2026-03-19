package test

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/handlers"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/helpers"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
	"github.com/gin-gonic/gin"
)

func TestAPI(t *testing.T) {
	t.Run("apiListarLivros", func(t *testing.T) {
		TestHandlerListarLivros(t)
	})
	t.Run("apiCadastrarLivros", func(t *testing.T) {
		TestHandlerCadastrarLivros(t)
	})
	t.Run("apiBuscarID", func(t *testing.T) {
		TestHandlerBuscarID(t)
	})
}

func ChecarCodigo(t *testing.T, recebido int, esperado int) {
	t.Helper()
	if recebido != esperado {
		t.Errorf("Esperava %d ,recebi %d", esperado, recebido)
	}
}

func criarLivroTesteAux(t *testing.T) models.Livro {
	livro := models.Livro{Titulo: "Livro de Teste", Autor: "Autor Teste", Preco: 120, Quantidade: 100, Ano: 1984}

	if err := database.DB.Create(&livro).Error; err != nil {
		t.Fatalf("Erro ao criar livro teste: %v", err)
	}

	// A partir daqui, qualquer teste que chamar essa função
	// será limpo automaticamente no final.
	t.Cleanup(func() {
		database.DB.Unscoped().Delete(&livro)
	})

	return livro
}

func setupRouter() *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/livros", handlers.HandlerListarLivros)
	r.GET("/livros/listar_id", handlers.HandlerListarID)
	r.GET("/livros/relatorio", handlers.HandlerListarRelatorio)
	r.GET("/livros/:livro_id", handlers.HandlerBuscarID)
	r.GET("/livros/titulo/:titulo", handlers.HandlerBuscarTitulo)
	r.GET("/livros/autor/:autor", handlers.HandlerBuscarAutor)
	r.POST("/livros", handlers.HandlerCadastrarLivro)
	r.PUT("/livros/:id", handlers.HandlerAtualizarLivro)
	r.DELETE("/livros/:id", handlers.HandlerDeletarLivro)
	return r
}

func TestHandlerListarLivros(t *testing.T) {
	r := setupRouter()
	criarLivroTeste(t)
	req, _ := http.NewRequest("GET", "/livros", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("Esperava 200, recebi %d", w.Code)
	}
}

func TestHandlerListarID(t *testing.T) {
	r := setupRouter()
	criarLivroTesteAux(t)
	req, _ := http.NewRequest("GET", "/livros/listar_id", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	ChecarCodigo(t, w.Code, http.StatusOK)
}

func TestHandlerRelatorio(t *testing.T) {
	r := setupRouter()

	req, _ := http.NewRequest("GET", "/livros/relatorio", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	ChecarCodigo(t, w.Code, http.StatusOK)
}

func TestHandlerBuscarID(t *testing.T) {
	r := setupRouter()
	livro := helpers.CriarLivroTesteAux(t)
	url := fmt.Sprintf("/livros/%d", livro.ID)
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	ChecarCodigo(t, w.Code, http.StatusOK)
	var resultado models.Livro
	_ = json.Unmarshal(w.Body.Bytes(), &resultado)
	if resultado.Titulo != livro.Titulo {
		t.Errorf("Esperava título %s, mas veio %s", livro.Titulo, resultado.Titulo)
	}
}

func TestHandlerCadastrarLivros(t *testing.T) {
	r := setupRouter()
	body := `{"titulo": "O Senhor dos Anéis", "autor": "Tolkien", "preco": 59.90, "ano": 1954, "quantidade": 10}`
	req, _ := http.NewRequest("POST", "/livros", strings.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	ChecarCodigo(t, w.Code, http.StatusCreated)
}

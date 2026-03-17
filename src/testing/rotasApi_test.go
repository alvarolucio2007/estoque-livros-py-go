package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/handlers"
	"github.com/gin-gonic/gin"
)

func TestAPI(t *testing.T) {
	t.Run("apiListarLivros", func(t *testing.T) {
		testarHandlerCarregarLivros(t)
	})
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
func TestHandlerCadastrarLivros(t *testing.T){
	r:=setupRouter()
	body := `{"titulo": "O Senhor dos Anéis", "autor": "Tolkien", "preco": 59.90, "ano": 1954, "quantidade": 10}`
}

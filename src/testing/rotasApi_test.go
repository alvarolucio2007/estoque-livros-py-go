package test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
)

func TestAPI(t *testing.T) {
	t.Run("apiListarLivros", func(t *testing.T) {
		testarHandlerCarregarLivros(t)
	})
}

// 1. Defina o Handler como uma função separada
func HandlerListarLivros(c *gin.Context) {
	livros, err := carregarDados()
	if err != nil {
		// Se o erro for "não há livro", talvez você queira dar um 404 ou 200 vazio
		if err.Error() == "não há livro" {
			c.JSON(http.StatusOK, []Livro{}) // Retorna lista vazia, status 200
			return
		}

		// Erro real de banco (conexão, etc)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro interno no servidor"})
		return
	}

	c.JSON(http.StatusOK, livros)
}

func testarHandlerCarregarLivros(t *testing.T) {
	gin.SetMode(gin.TestMode)
	r := gin.Default()
	r.GET("/livros", HandlerListarLivros)
	criarLivroTeste(t)
	url := "/livros"
	req, _ := http.NewRequest("GET", url, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	if w.Code != http.StatusOK {
		t.Errorf("esperava status 200,recebi %d", w.Code)
	}
}

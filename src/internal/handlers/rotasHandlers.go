// Package handlers "monta" as rotas e tals da API pra produção e testes.
package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/verifiers"
	"github.com/gin-gonic/gin"
)

func SendError(c *gin.Context, err error) {
	status, resposta := ErrorHandler(err)
	c.JSON(status, resposta)
}

func ErrorHandler(err error) (int, gin.H) {
	if errors.Is(err, models.ErrTituloVazio) || errors.Is(err, models.ErrPrecoInvalido) || errors.Is(err, models.ErrIDNulo) {
		return http.StatusBadRequest, gin.H{"error": err.Error()}
	}
	if errors.Is(err, models.ErrLivroNaoEncontrado) {
		return http.StatusNotFound, gin.H{"error": err.Error()}
	}
	if errors.Is(err, models.ErrConexaoFatalDB) || errors.Is(err, models.ErrInternoDBFatalDB) {
		return http.StatusInternalServerError, gin.H{"error": "erro interno na DB"}
	}
	var (
		syntaxErrJSON    *json.SyntaxError
		unmarshalErrJSON *json.UnmarshalTypeError
	)
	if errors.As(err, &syntaxErrJSON) {
		return http.StatusBadRequest, gin.H{"error": models.ErrSintaxeJSONAPI.Error()}
	}
	if errors.As(err, &unmarshalErrJSON) {
		return http.StatusBadRequest, gin.H{"error": models.ErrTipagemJSONAPI.Error()}
	}
	return http.StatusInternalServerError, gin.H{"error": "erro catastrófico, contate o dev!"}
}

func HandlerListarLivros(c *gin.Context) {
	livros, err := database.CarregarDados()
	if err != nil {
		SendError(c, err)
		return
	}

	c.JSON(http.StatusOK, livros)
}

func HandlerListarID(c *gin.Context) {
	listaID, err := database.ListarID()
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, listaID)
}

func HandlerListarRelatorio(c *gin.Context) {
	relatorio, err := database.GerarRelatorio()
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, relatorio)
}

func HandlerBuscarID(c *gin.Context) {
	idStr := c.Param("livro_id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": models.ErrIDNulo.Error()})
		return
	}
	resultado, err := verifiers.ServicoBuscarLivroID(uint(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func HandlerBuscarTitulo(c *gin.Context) {
	tituloStr := c.Param("titulo")
	resultado, err := verifiers.ServicoBuscarLivroTitulo(tituloStr)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func HandlerBuscarAutor(c *gin.Context) {
	autorStr := c.Param("autor")
	resultado, err := verifiers.ServicoBuscarLivroAutor(autorStr)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusOK, resultado)
}

func HandlerCadastrarLivro(c *gin.Context) {
	var novoLivro models.LivroCadastrar
	if err := c.ShouldBindJSON(&novoLivro); err != nil {
		SendError(c, err)
		return
	}
	err := verifiers.ServicoAdicionarLivro(novoLivro.Titulo, novoLivro.Autor, novoLivro.Preco, novoLivro.Ano, novoLivro.Quantidade)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusCreated, "criado com sucesso")
}

func HandlerAtualizarLivro(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		SendError(c, models.ErrIDNulo)
		return
	}
	var dadosAtualizados models.Livro
	if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
		SendError(c, err)
		return
	}
	err = verifiers.ServicoAtualizarLivro(uint(idUint), dadosAtualizados)
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusAccepted, dadosAtualizados)
}

func HandlerDeletarLivro(c *gin.Context) {
	idStr := c.Param("id")
	idUint, err := strconv.ParseUint(idStr, 10, 64)
	if err != nil {
		SendError(c, models.ErrIDNulo)
		return
	}
	err = verifiers.ServicoDeletarLivro(uint(idUint))
	if err != nil {
		SendError(c, err)
		return
	}
	c.JSON(http.StatusAccepted, "livro deletado com sucesso")
}

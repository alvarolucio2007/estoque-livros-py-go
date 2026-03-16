package routes

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
	"github.com/gin-gonic/gin"
)

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

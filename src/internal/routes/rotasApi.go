// Package routes faz as rotas da API para conectar ao Streamlit Python
package routes

import (
	"strconv"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/handlers"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/verifiers"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func SetupAPI() {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	config := cors.Config{
		AllowOrigins:     []string{"http://localhost:8501"},
		AllowMethods:     []string{"GET", "POST", "PATCH", "DELETE", "PUT"},
		AllowHeaders:     []string{"Origin", "Content-Type"},
		AllowCredentials: true,
	}
	r.Use(cors.New(config))

	r.GET("/livros", func(c *gin.Context) {
		livros, err := database.CarregarDados()
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, livros)
	})
	r.GET("/livros/listar_id", func(c *gin.Context) {
		resultado, err := database.ListarID()
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
		}
		c.JSON(200, resultado)
	})
	r.GET("/livros/relatorio", func(c *gin.Context) {
		resultado, err := database.GerarRelatorio()
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, resultado)
	})
	r.GET("/livros/:livro_id", func(c *gin.Context) {
		idStr := c.Param("livro_id")
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		resultado, err := verifiers.ServicoBuscarLivroID(uint(idUint))
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, resultado)
	})
	r.GET("/livros/titulo/:titulo", func(c *gin.Context) {
		tituloStr := c.Param("titulo")
		resultado, err := verifiers.ServicoBuscarLivroTitulo(tituloStr)
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, resultado)
	})
	r.GET("/livros/autor/:autor", func(c *gin.Context) {
		tituloStr := c.Param("autor")
		resultado, err := verifiers.ServicoBuscarLivroAutor(tituloStr)
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, resultado)
	})
	r.POST("/livros", func(c *gin.Context) {
		var novoLivro models.LivroCadastrar
		if err := c.ShouldBindJSON(&novoLivro); err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		err := verifiers.ServicoAdicionarLivro(novoLivro.Titulo, novoLivro.Autor, novoLivro.Preco, novoLivro.Ano, novoLivro.Quantidade)
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(201, "criado com sucesso")
	})
	r.PUT("/livros/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		var dadosAtualizados models.Livro
		if err := c.ShouldBindJSON(&dadosAtualizados); err != nil {
			c.JSON(404, models.RespostaErro{Mensagem: err.Error()})
			return
		}
		err = verifiers.ServicoAtualizarLivro(uint(idUint), dadosAtualizados)
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, dadosAtualizados)
	})
	r.DELETE("/livros/:id", func(c *gin.Context) {
		idStr := c.Param("id")
		idUint, err := strconv.ParseUint(idStr, 10, 64)
		if err != nil {
			c.JSON(400, gin.H{"error": err.Error()})
			return
		}
		err = verifiers.ServicoDeletarLivro(uint(idUint))
		if err != nil {
			status, resposta := handlers.ErrorHandler(err)
			c.JSON(status, resposta)
			return
		}
		c.JSON(200, "livro deletado com sucesso")
	})

	err := r.Run(":8000")
	if err != nil {
		return
	}
}

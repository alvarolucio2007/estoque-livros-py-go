// Package routes faz as rotas da API para conectar ao Streamlit Python
package routes

import (
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/handlers"
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
	r.GET("/livros", handlers.HandlerListarLivros)
	r.GET("/livros/listar_id", handlers.HandlerListarID)
	r.GET("/livros/relatorio", handlers.HandlerListarRelatorio)
	r.GET("/livros/:livro_id", handlers.HandlerBuscarID)
	r.GET("/livros/titulo/:titulo", handlers.HandlerBuscarTitulo)
	r.GET("/livros/autor/:autor", handlers.HandlerBuscarAutor)
	r.POST("/livros", handlers.HandlerCadastrarLivro)
	r.PUT("/livros/:id", handlers.HandlerAtualizarLivro)
	r.DELETE("/livros/:id", handlers.HandlerDeletarLivro)
	if err := r.Run(":8000"); err != nil {
		panic("falha ao iniciar o servidor: " + err.Error())
	}
}

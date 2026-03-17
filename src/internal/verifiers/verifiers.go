// Package verifiers verifica regras de negócio.
package verifiers

import (
	"fmt"
	"time"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
)

func ServicoAdicionarLivro(titulo string, autor string, preco float64, ano int, quantidade uint) error {
	if titulo == "" {
		return models.ErrTituloVazio
	}
	if autor == "" {
		return models.ErrAutorVazio
	}
	if preco < 0 {
		return models.ErrPrecoInvalido
	}
	if ano > time.Now().Year()+1 {
		return models.ErrAnoInvalido
	}
	disponivel := quantidade > 0

	novoLivro := models.Livro{
		Titulo:     titulo,
		Autor:      autor,
		Preco:      preco,
		Ano:        ano,
		Quantidade: quantidade,
		Disponivel: &disponivel,
		// O ID você não precisa passar se ele for AutoIncrement
	}
	return database.AdicionarLivro(novoLivro)
}

func ServicoDeletarLivro(id uint) error {
	if id == 0 {
		return models.ErrIDNulo
	}
	return database.DeletarLivro(id)
}

func ServicoAtualizarLivro(id uint, novosDados models.Livro) error {
	livroExistente, err := database.BuscarPorID(id)
	if err != nil {
		return fmt.Errorf("erro ao consultar o banco: %w", err)
	}
	if livroExistente == nil {
		return models.ErrLivroNaoEncontrado
	}
	if novosDados.Titulo == "" {
		return models.ErrTituloVazio
	}
	if novosDados.Autor == "" {
		return models.ErrAutorVazio
	}
	if novosDados.Preco < 0 {
		return models.ErrPrecoInvalido
	}
	if novosDados.Ano > time.Now().Year()+1 {
		return models.ErrAnoInvalido
	}

	status := novosDados.Quantidade > 0
	novosDados.Disponivel = &status

	err = database.AtualizarLivro(id, novosDados)
	if err != nil {
		return err // o erro já vem limpo!
	}
	return nil
}

func ServicoBuscarLivroTitulo(titulo string) ([]models.Livro, error) {
	if titulo == "" {
		return nil, models.ErrTituloVazio
	}
	return database.BuscarLivroTitulo(titulo) // buscarLivroTitulo já retorna []Livro,error então não precisa colocar nil
}

func ServicoBuscarLivroAutor(autor string) ([]models.Livro, error) {
	if autor == "" {
		return nil, models.ErrAutorVazio
	}
	return database.BuscarLivroAutor(autor) // buscarLivroAutor já retorna []Livro,nil então não precisa colocar nil
}

func ServicoBuscarLivroID(id uint) (*models.Livro, error) {
	if id == 0 {
		return nil, models.ErrIDNulo
	}
	return database.BuscarPorID(id)
}

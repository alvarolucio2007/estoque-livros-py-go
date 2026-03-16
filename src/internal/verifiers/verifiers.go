// Verifica regras de negócio.
package verifiers

import (
	"errors"
	"fmt"
	"time"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/database"
	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"
)

func ServicoAdicionarLivro(titulo string, autor string, preco float64, ano int, quantidade uint) error {
	if titulo == "" {
		return errors.New("título não pode ficar em branco")
	}
	if autor == "" {
		return errors.New("autor não pode ficar em branco")
	}
	if preco < 0 {
		return errors.New("o preço tem que ser maior que 0")
	}
	if ano > time.Now().Year()+1 {
		return fmt.Errorf("o ano %d excede o limite de lançamentos futuros (%d)", ano, time.Now().Year()+1)
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
		return errors.New("id não pode ser 0")
	}
	return database.DeletarLivro(id)
}

func ServicoAtualizarLivro(id uint, novosDados models.Livro) error {
	livroExistente, err := database.BuscarPorID(id)
	if err != nil {
		return fmt.Errorf("erro ao consultar o banco: %w", err)
	}
	if livroExistente == nil {
		return errors.New("livro não encontrado")
	}
	if novosDados.Titulo == "" {
		return errors.New("título não pode ficar vazio")
	}
	if novosDados.Autor == "" {
		return errors.New("autor não pode ficar vazio")
	}
	if novosDados.Preco < 0 {
		return errors.New("o preco nao pode ser negativo")
	}
	if novosDados.Ano > time.Now().Year()+1 {
		return errors.New("ano inválido")
	}

	status := novosDados.Quantidade > 0
	novosDados.Disponivel = &status

	err = database.AtualizarLivro(id, novosDados)
	if err != nil {
		return fmt.Errorf("falha ao atualizar: %w", err)
	}
	return nil
}

func ServicoBuscarLivroTitulo(titulo string) ([]models.Livro, error) {
	if titulo == "" {
		return nil, errors.New("título não pode ficar vazio")
	}
	return database.BuscarLivroTitulo(titulo) // buscarLivroTitulo já retorna []Livro,error então não precisa colocar nil
}

func servicoBuscarLivroAutor(autor string) ([]models.Livro, error) {
	if autor == "" {
		return nil, errors.New("autor não pode ficar vazio")
	}
	return database.BuscarLivroAutor(autor) // buscarLivroAutor já retorna []Livro,nil então não precisa colocar nil
}

func ServicoBuscarLivroID(id uint) (*models.Livro, error) {
	if id == 0 {
		return nil, errors.New("id não pode ser 0")
	}
	return database.BuscarPorID(id)
}

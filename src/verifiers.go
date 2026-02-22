package main

import (
	"errors"
	"fmt"
	"slices"
	"time"
)

func verificarID(id uint) bool {
	return slices.Contains(listarID(), id)
}

func cadastrarLivro(titulo string, autor string, preco float64, ano int, quantidade uint) error {
	if titulo == "" {
		return errors.New("título não pode ficar em branco")
	}
	if autor == "" {
		return errors.New("autor não pode ficar em branco")
	}
	if preco <= 0 {
		return errors.New("o preço tem que ser maior que 0")
	}
	if ano > time.Now().Year()+1 {
		return fmt.Errorf("o ano %d excede o limite de lançamentos futuros (%d)", ano, time.Now().Year()+1)
	}
	disponivel := quantidade > 0

	novoLivro := Livro{
		Titulo:     titulo,
		Autor:      autor,
		Preco:      preco,
		Ano:        ano,
		Quantidade: quantidade,
		Disponivel: disponivel,
		// O ID você não precisa passar se ele for AutoIncrement

	}
	if err := DB.Create(&novoLivro).Error; err != nil {
		return fmt.Errorf("erro ao inserir no banco: %w", err)
	}
	return nil
}

func servicoAtualizarLivro(id uint, novosDados Livro) error {
	livroExistente, err := buscarPorID(id)
	if err != nil {
		return fmt.Errorf("erro ao consultar o banco: %w", err)
	}
	if livroExistente == nil {
		return errors.New("livro não encontrado")
	}
	if novosDados.Titulo == "" {
		return errors.New("título não pode ficar vazio")
	}
	if novosDados.Preco < 0 {
		return errors.New("o preco nao pode ser negativo")
	}
	if novosDados.Ano > time.Now().Year()+1 {
		return errors.New("ano inválido")
	}
	novosDados.Disponivel = novosDados.Quantidade > 0
	err = atualizarLivro(id, novosDados)
	if err != nil {
		return fmt.Errorf("falha ao atualizar: %w", err)
	}
	return nil
}

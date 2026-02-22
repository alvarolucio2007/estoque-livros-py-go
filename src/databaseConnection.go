package main

import (
	"errors"
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBanco() {
	dsn := "host=localhost user=user_projeto password=password_projeto dbname=db_estoque port=5432 sslmode=disable"
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Não foi possível conectar ao Postgres! ", err)
	}
	fmt.Println("Banco conectado com sucesso!")
	err = DB.AutoMigrate(&Livro{})
	if err != nil {
		log.Fatal("Erro interno no servidor!", err)
	}
}

func checarResultado(resultado *gorm.DB) error {
	if resultado.Error != nil {
		return resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return errors.New("nenhum livro encontrado com esse ID")
	}
	return nil
}

func adicionarLivro(dados Livro) error {
	novoLivro := Livro{
		Titulo:     dados.Titulo,
		Autor:      dados.Autor,
		Preco:      dados.Preco,
		Ano:        dados.Ano,
		Quantidade: dados.Quantidade,
		Disponivel: dados.Disponivel,
	}
	resultado := DB.Create(&novoLivro)
	if resultado.Error != nil {
		return resultado.Error
	}
	return nil
}

func carregarDados() ([]Livro, error) {
	var livros []Livro
	resultado := DB.Find(&livros)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	return livros, nil
}

func deletarLivro(id uint) error {
	return checarResultado(DB.Delete(&Livro{}, id))
}

func atualizarLivro(id uint, dados Livro) error {
	return checarResultado(DB.Model(&Livro{}).Where("id=?", id).Updates(dados))
}

func buscarLivroTitulo(titulo string) ([]Livro, error) {
	var livrosEncontados []Livro
	res := DB.Where("titulo LIKE ?", "%"+titulo+"%").Find(&livrosEncontados)
	err := checarResultado(res)
	if err != nil {
		return nil, err
	}
	return livrosEncontados, nil
}

func buscarLivroAutor(autor string) ([]Livro, error) {
	var livrosEncontados []Livro
	res := DB.Where("autor LIKE ?", "%"+autor+"%").Find(&livrosEncontados)
	err := checarResultado(res)
	if err != nil {
		return nil, err
	}
	return livrosEncontados, nil
}
func gerarRelatorio()

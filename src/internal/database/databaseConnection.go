// Package database se conecta ao banco de dados e faz todas as operações com ele.
package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/alvarolucio2007/estoque-livros-py/src/internal/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConectarBanco() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		dsn = "host=db user=user password=password dbname=estoque_db port=5432 sslmode=disable"
	}
	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Não foi possível conectar ao Postgres! ", err)
	}
	fmt.Println("Banco conectado com sucesso!")
	err = DB.AutoMigrate(&models.Livro{})
	if err != nil {
		log.Fatal("Erro interno no servidor!", err)
	}
}

func ChecarResultado(resultado *gorm.DB) error {
	if resultado.Error != nil {
		return resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return errors.New("nenhum livro encontrado com esse ID")
	}
	return nil
}

func AdicionarLivro(dados models.Livro) error {
	resultado := DB.Create(&dados)
	if resultado.Error != nil {
		return resultado.Error
	}
	return nil
}

func CarregarDados() ([]models.Livro, error) {
	var livros []models.Livro
	resultado := DB.Find(&livros)
	if resultado.Error != nil {
		return nil, resultado.Error
	}
	return livros, nil
}

func DeletarLivro(id uint) error {
	resultado := DB.Delete(&models.Livro{}, id)
	if resultado.Error != nil {
		return resultado.Error
	}
	if resultado.RowsAffected == 0 {
		return errors.New("livro não encontrado")
	}
	return nil
}

func AtualizarLivro(id uint, dados models.Livro) error {
	dados.ID = id
	return ChecarResultado(DB.Model(&models.Livro{}).Where("id=?", id).Omit("id").Select("*").Updates(dados))
}

func BuscarLivroTitulo(titulo string) ([]models.Livro, error) {
	var livrosEncontados []models.Livro
	res := DB.Where("titulo LIKE ?", "%"+titulo+"%").Find(&livrosEncontados)
	err := ChecarResultado(res)
	if err != nil {
		return nil, err
	}
	if len(livrosEncontados) == 0 {
		return nil, errors.New("não há livros")
	}
	return livrosEncontados, nil
}

func BuscarLivroAutor(autor string) ([]models.Livro, error) {
	var livrosEncontados []models.Livro
	res := DB.Where("autor LIKE ?", "%"+autor+"%").Find(&livrosEncontados)
	err := ChecarResultado(res)
	if err != nil {
		return nil, err
	}
	if len(livrosEncontados) == 0 {
		return nil, errors.New("não há livros")
	}
	return livrosEncontados, nil
}

func GerarRelatorio() (map[string]any, error) {
	var total, disponiveis, indisponiveis int64
	var valorTotal float64
	DB.Model(&models.Livro{}).Count(&total)
	DB.Model(&models.Livro{}).Where("disponivel=?", true).Count(&disponiveis)
	DB.Model(&models.Livro{}).Where("disponivel=?", false).Count(&indisponiveis)
	DB.Model(&models.Livro{}).Select("SUM(preco*quantidade)").Scan(&valorTotal)
	return map[string]any{
		"total_livros":         total,
		"livros_disponiveis":   disponiveis,
		"livros_indisponiveis": indisponiveis,
		"valor_total_estoque":  valorTotal,
	}, nil
}

func ListarID() ([]uint, error) {
	var ids []uint
	err := DB.Model(&models.Livro{}).Distinct().Pluck("id", &ids).Error
	if err != nil {
		return nil, models.ErrInternoDBFatalDB
	}
	if len(ids) == 0 {
		return nil, models.ErrLivroNaoEncontrado
	}
	return ids, nil
}

func BuscarPorID(id uint) (*models.Livro, error) {
	var livro models.Livro
	if err := DB.First(&livro, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("livro não encontrado")
		}
		return nil, err
	}
	return &livro, nil
}

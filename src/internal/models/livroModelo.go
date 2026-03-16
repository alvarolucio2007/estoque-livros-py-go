// Package models cria modelos que serão utilizados em todo o código.
package models

import "gorm.io/gorm"

type Livro struct {
	ID         uint    `json:"id" gorm:"primaryKey;autoIncrement"`
	Titulo     string  `json:"titulo" gorm:"not null"`
	Autor      string  `json:"autor" gorm:"not null"`
	Preco      float64 `json:"preco" gorm:"type:decimal(10,2)"`
	Ano        int     `json:"ano"`
	Quantidade uint    `json:"quantidade"`
	Disponivel *bool   `json:"disponivel" gorm:"default:true"`
}

func (l *Livro) BeforeSave(tx *gorm.DB) (err error) {
	status := l.Quantidade > 0
	l.Disponivel = &status
	return nil
}

type LivroCadastrar struct {
	Titulo     string  `json:"titulo" gorm:"not null"`
	Autor      string  `json:"autor" gorm:"not null"`
	Preco      float64 `json:"preco" gorm:"type:decimal(10,2)"`
	Ano        int     `json:"ano"`
	Quantidade uint    `json:"quantidade"`
}
type RespostaErro struct {
	Mensagem string `json:"mensagem"`
	Detalhe  string `json:"detalhe,omitempty"`
}

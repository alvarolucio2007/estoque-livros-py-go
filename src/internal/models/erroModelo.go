package models

import (
	"errors"
)

// Erros exclusivos Database+conexão postgres
var (
	ErrConexaoFatalDB   = errors.New("não foi possível conectar ao postgres")
	ErrInternoDBFatalDB = errors.New("erro interno na base de dados")
)

// Erros exclusivos verifiers

// Erros exclusivos API
var (
	ErrGerarRelatorioAPI = errors.New("erro ao gerar relatorio")
	ErrSintaxeJSONAPI    = errors.New("erro de sintaxe no JSON da api")
	ErrTipagemJSONAPI    = errors.New("erro de tipagem no JSON da api")
	ErrValidacaoJSONAPI  = errors.New("erro de validacao no JSON da api")
)

// Erros Universais
var (
	ErrTituloVazio        = errors.New("o título do livro é obrigatório")
	ErrAutorVazio         = errors.New("o nome do autor é obrigatório")
	ErrPrecoInvalido      = errors.New("o preço não pode ser 0 ou negativo")
	ErrAnoInvalido        = errors.New("o ano é inválido")
	ErrEstoquevazio       = errors.New("quantidade deve ser informada")
	ErrLivroNaoEncontrado = errors.New("nenhum livro encontrado")
	ErrIDNulo             = errors.New("o ID precisa ser válido")
)

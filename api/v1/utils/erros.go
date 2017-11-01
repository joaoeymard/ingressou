package utils

import (
	"errors"
	"fmt"
	"strings"
)

var (
	// Errors Conteudo de erros
	Errors = map[string]error{
		"NOT_FOUND": errors.New(`{
			"error": {
				"msg": "Registro não encontrado!",
				"code": 1
			}
		}`),
		"METHOD_DEFAULT": errors.New(`{
			"error": {
				"msg": "O método recebido não foi encontrado!",
				"code": 2
			}
		}`),
	}
)

// ValueInvalidos Tratamento de atributo com tipo diferente
func ValueInvalidos(atributo, tipo string) error {
	return fmt.Errorf(`{
		"error": {
			"msg": "Atributo %v: É do tipo %v!",
			"code": 1000
		}
	}`, atributo, tipo)
}

// ValueRequired Tratamento de atributo obrigatórios
func ValueRequired(atributo string) error {
	return fmt.Errorf(`{
		"error": {
			"msg": "Atributo %v: É obrigatório!",
			"code": 1001
		}
	}`, atributo)
}

// ValueMinino Tratamento de atributo com valor minino
func ValueMinino(atributo string, valor int16) error {
	return fmt.Errorf(`{
		"error": {
			"msg": "Atributo %v: Tem valor mínimo %v!",
			"code": 1002
		}
	}`, atributo, valor)
}

// ValueMaximo Tratamento de atributo com valor maximo
func ValueMaximo(atributo string, valor int16) error {
	return fmt.Errorf(`{
		"error": {
			"msg": "Atributo %v: Tem valor máximo %v!",
			"code": 1003
		}
	}`, atributo, valor)
}

// BancoDados Tratamento do erro quando é com o Banco de Dados
func BancoDados(err error) error {
	if strings.Contains(err.Error(), "pq") {
		retorno := strings.Split(err.Error(), ":")[1][1:]
		retorno = strings.Replace(retorno, "\"", "'", -1)
		err = fmt.Errorf(`{
			"error": {
				"msg": "postgres: %v!",
				"code": 2000
			}
		}`, retorno)
	}
	return err
}

// BancoDadosMethod Tratamento do erro para os metodos do Banco de Dados.
func BancoDadosMethod(method string) error {
	err := fmt.Errorf(`{
		"error": {
			"msg": "postgres: Erro ao %v um registro da tabela!",
			"code": 2001
		}
	}`, method)
	return err
}

// FormatError Responsavel por fomartar o JSON de retorno do erro
func FormatError(err error) error {
	err = fmt.Errorf(`{
		"error": {
			"msg": "%v",
			"code": 3000
		}
	}`, err)
	return err
}

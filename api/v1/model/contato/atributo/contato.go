package atributo

import (
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
)

// ValidValues Responsavel por validar os atributos recebidos
func ValidValues(params map[string]interface{}) error {

	var (
		valid = map[string]func(interface{}) error{
			"endereco":            endereco,
			"complemento":         complemento,
			"referencia":          referencia,
			"bairro":              bairro,
			"cep":                 cep,
			"cidade":              cidade,
			"uf":                  uf,
			"telefone_principal":  telefonePrincipal,
			"telefone_secundario": telefoneSecundario,
			"telefone_terciario":  telefoneTerciario,
			"email":               email,
		}
	)

	for key, value := range params {
		if retorno := valid[key](value); retorno != nil {
			return retorno
		}
	}

	return nil

}

func endereco(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("endereco", "string")
	}

	if data == "" {
		return utils.ValueRequired("endereco")
	}

	if len(data) < 3 {
		return utils.ValueMinino("endereco", 3)
	}

	if len(data) > 100 {
		return utils.ValueMaximo("endereco", 100)
	}

	return nil

}

func complemento(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("complemento", "string")
	}

	if data == "" {
		return utils.ValueRequired("complemento")
	}

	if len(data) < 3 {
		return utils.ValueMinino("complemento", 3)
	}

	if len(data) > 150 {
		return utils.ValueMaximo("complemento", 150)
	}

	return nil

}

func referencia(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("referencia", "string")
	}

	if data == "" {
		return utils.ValueRequired("referencia")
	}

	if len(data) < 3 {
		return utils.ValueMinino("referencia", 3)
	}

	if len(data) > 200 {
		return utils.ValueMaximo("referencia", 200)
	}

	return nil

}

func bairro(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("bairro", "string")
	}

	if data == "" {
		return utils.ValueRequired("bairro")
	}

	if len(data) < 3 {
		return utils.ValueMinino("bairro", 3)
	}

	if len(data) > 80 {
		return utils.ValueMaximo("bairro", 80)
	}

	return nil

}

func cep(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("cep", "string")
	}

	if data == "" {
		return utils.ValueRequired("cep")
	}

	if len(data) < 7 {
		return utils.ValueMinino("cep", 7)
	}

	if len(data) > 15 {
		return utils.ValueMaximo("cep", 15)
	}

	return nil

}

func cidade(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("cidade", "string")
	}

	if data == "" {
		return utils.ValueRequired("cidade")
	}

	if len(data) < 3 {
		return utils.ValueMinino("cidade", 3)
	}

	if len(data) > 80 {
		return utils.ValueMaximo("cidade", 80)
	}

	return nil

}

func uf(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("uf", "string")
	}

	if data == "" {
		return utils.ValueRequired("uf")
	}

	if len(data) < 2 {
		return utils.ValueMinino("uf", 2)
	}

	if len(data) > 2 {
		return utils.ValueMaximo("uf", 2)
	}

	return nil

}

func telefonePrincipal(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("telefone_principal", "string")
	}

	if data == "" {
		return utils.ValueRequired("telefone_principal")
	}

	if len(data) < 3 {
		return utils.ValueMinino("telefone_principal", 3)
	}

	if len(data) > 16 {
		return utils.ValueMaximo("telefone_principal", 16)
	}

	return nil

}

func telefoneSecundario(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("telefone_secundario", "string")
	}

	if data == "" {
		return utils.ValueRequired("telefone_secundario")
	}

	if len(data) < 3 {
		return utils.ValueMinino("telefone_secundario", 3)
	}

	if len(data) > 16 {
		return utils.ValueMaximo("telefone_secundario", 16)
	}

	return nil

}

func telefoneTerciario(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("telefone_terciario", "string")
	}

	if data == "" {
		return utils.ValueRequired("telefone_terciario")
	}

	if len(data) < 3 {
		return utils.ValueMinino("telefone_terciario", 3)
	}

	if len(data) > 16 {
		return utils.ValueMaximo("telefone_terciario", 16)
	}

	return nil

}

func email(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("email", "string")
	}

	if data == "" {
		return utils.ValueRequired("email")
	}

	if len(data) < 3 {
		return utils.ValueMinino("email", 3)
	}

	if len(data) > 50 {
		return utils.ValueMaximo("email", 50)
	}

	return nil

}

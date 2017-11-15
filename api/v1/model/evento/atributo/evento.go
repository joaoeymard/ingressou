package atributo

import (
	"github.com/JoaoEymard/ingressou/api/v1/utils"
)

var (
	// Filtros lista de filtro para consulta
	Filtros = map[string]string{
		"eventoLink": "TIEVENTO.link = '%v'",
	}
)

// ValidValues Responsavel por validar os atributos recebidos
func ValidValues(params map[string]interface{}) error {

	// titulo character varying(100) NOT NULL,
	// imagem text NOT NULL,
	// cidade character varying(80) NOT NULL,
	// uf character varying(2) NOT NULL,
	// localidade character varying(100),
	// taxa double precision NOT NULL,
	// mapa text,
	// descricao text,
	// link character varying(100) NOT NULL,
	// data_criacao timestamp without time zone NOT NULL,
	// status boolean DEFAULT false,

	var (
		valid = map[string]func(interface{}) error{
			"titulo":       titulo,
			"imagem":       imagem,
			"cidade":       cidade,
			"uf":           uf,
			"localidade":   localidade,
			"taxa":         taxa,
			"mapa":         mapa,
			"descricao":    descricao,
			"link":         link,
			"data_criacao": dataCriacao,
			"status":       status,
		}
	)

	for key, value := range params {
		if retorno := valid[key](value); retorno != nil {
			return retorno
		}
	}

	return nil

}

func titulo(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("nome", "string")
	}

	if data == "" {
		return utils.ValueRequired("nome")
	}

	if len(data) < 3 {
		return utils.ValueMinino("nome", 3)
	}

	if len(data) > 50 {
		return utils.ValueMaximo("nome", 50)
	}

	return nil
}

func imagem(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("senha", "string")
	}

	if data == "" {
		return utils.ValueRequired("senha")
	}

	if len(data) < 6 {
		return utils.ValueMinino("senha", 6)
	}

	if len(data) > 128 {
		return utils.ValueMaximo("senha", 128)
	}

	return nil
}

func cidade(value interface{}) error {

	_, valueType := value.(bool)
	if !valueType {
		return utils.ValueInvalidos("ativo", "booleano")
	}

	return nil
}

func uf(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("cpf", "string")
	}

	if data == "" {
		return utils.ValueRequired("cpf")
	}

	if len(data) < 11 {
		return utils.ValueMinino("cpf", 11)
	}

	if len(data) > 15 {
		return utils.ValueMaximo("cpf", 15)
	}

	return nil
}

func localidade(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("data_nascimento", "string")
	}

	if data == "" {
		return utils.ValueRequired("data_nascimento")
	}

	return nil
}

func taxa(value interface{}) error {

	data, valueType := value.(string)
	if !valueType {
		return utils.ValueInvalidos("sexo", "string")
	}

	if data == "" {
		return utils.ValueRequired("sexo")
	}

	if len(data) < 3 {
		return utils.ValueMinino("sexo", 3)
	}

	if len(data) > 15 {
		return utils.ValueMaximo("sexo", 15)
	}

	return nil
}

func mapa(value interface{}) error {

	data, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nivel", "float64")
	}

	if data < 1 {
		return utils.ValueMinino("nivel", 1)
	}

	return nil
}

func descricao(value interface{}) error {

	data, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nivel", "float64")
	}

	if data < 1 {
		return utils.ValueMinino("nivel", 1)
	}

	return nil
}

func link(value interface{}) error {

	data, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nivel", "float64")
	}

	if data < 1 {
		return utils.ValueMinino("nivel", 1)
	}

	return nil
}

func dataCriacao(value interface{}) error {

	data, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nivel", "float64")
	}

	if data < 1 {
		return utils.ValueMinino("nivel", 1)
	}

	return nil
}

func status(value interface{}) error {

	data, valueType := value.(float64)
	if !valueType {
		return utils.ValueInvalidos("nivel", "float64")
	}

	if data < 1 {
		return utils.ValueMinino("nivel", 1)
	}

	return nil
}

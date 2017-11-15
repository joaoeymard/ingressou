package usuario

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JoaoEymard/ingressou/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressou/api/v1/model/usuario/atributo"
	"github.com/JoaoEymard/ingressou/api/v1/utils"
)

const (
	// Tabela referente ao usuario
	tableUsuario = "t_ingressou_usuario"
	// Tabela referente ao contato
	tableUsuarioContato = "t_ingressou_usuario_contato"
)

// Insert Adiciona um registro
func Insert(contentBody io.ReadCloser) ([]byte, int, error) {

	var contentJSON map[string]interface{}

	content, err := ioutil.ReadAll(contentBody)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(content, &contentJSON)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	if err = atributo.ValidValues(contentJSON); err != nil {
		return nil, http.StatusBadRequest, err
	}

	values := map[string]interface{}{
		"nome":            contentJSON["nome"],
		"senha":           contentJSON["senha"],
		"ativo":           contentJSON["ativo"],
		"cpf":             contentJSON["cpf"],
		"data_nascimento": contentJSON["data_nascimento"],
		"sexo":            contentJSON["sexo"],
		"nivel":           contentJSON["nivel"],
	}

	rows, err := postgres.InsertOne(tableUsuario, values)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusCreated, nil

}

// Find Retorna os eventos via json
func Find(params url.Values) ([]byte, int, error) {

	// Tratamento dos paramentros e filtro recebidos pela URL
	filter, order, limit, offset, err := postgres.SetParams(params, atributo.Filtros)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Consulta para saber o total de registro
	sqlTotal := fmt.Sprintf(`SELECT COUNT(TIUSUARIO.id) AS total
	FROM %v TIUSUARIO
	%v`, tableUsuario, filter)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rowTotal, err := postgres.SelectOne(sqlTotal)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rowTotal == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	// Consulta para coletar os registro
	sql := fmt.Sprintf(`SELECT TIUSUARIO.id AS id, TIUSUARIO.nome, TIUSUARIO.ultimo_acesso, TIUSUARIO.ativo, TIUSUARIO.cpf, TIUSUARIO.data_nascimento, TIUSUARIO.sexo, TIUSUARIO.nivel,
	(
		SELECT array_to_json (array_agg (row_to_json(dados_contatos.*) ) )
		FROM (
			SELECT TIUCONTATO.id AS contato_id, TIUCONTATO.endereco, TIUCONTATO.complemento, TIUCONTATO.referencia, TIUCONTATO.bairro, TIUCONTATO.cep, TIUCONTATO.cidade, TIUCONTATO.uf, TIUCONTATO.telefone_principal, TIUCONTATO.telefone_secundario, TIUCONTATO.telefone_terciario, TIUCONTATO.email
			FROM %v TIUCONTATO
			WHERE TIUCONTATO.id_usuario = TIUSUARIO.id
		) AS dados_contatos
	) AS contatos
	FROM %v TIUSUARIO
	%v %v %v %v`, tableUsuarioContato, tableUsuario, filter, order, limit, offset)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rows, err := postgres.Select(sql)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	err = nil
	// Montar o json de retorno
	for i, row := range rows {
		if row["contatos"] != nil {
			var jsonContato []map[string]interface{}
			// Converte a estrutura para json
			err = json.Unmarshal([]byte(row["contatos"].(string)), &jsonContato)
			if err != nil {
				break
			}
			rows[i]["contatos"] = jsonContato
		}
	}
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	// Monta a estrutura de retorno
	dados := map[string]interface{}{
		"dados": rows,
		"total": rowTotal["total"],
	}

	// Converte a estrutura para json
	retorno, err := json.Marshal(dados)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusOK, nil
}

// Update Adiciona um registro
func Update(contentBody io.ReadCloser, params url.Values) ([]byte, int, error) {

	var contentJSON map[string]interface{}

	if params.Get("usuarioID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("id")
	}

	content, err := ioutil.ReadAll(contentBody)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(content, &contentJSON)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	if err = atributo.ValidValues(contentJSON); err != nil {
		return nil, http.StatusBadRequest, err
	}

	values := map[string]interface{}{
		"nome":            contentJSON["nome"],
		"senha":           contentJSON["senha"],
		"ativo":           contentJSON["ativo"],
		"cpf":             contentJSON["cpf"],
		"data_nascimento": contentJSON["data_nascimento"],
		"sexo":            contentJSON["sexo"],
		"nivel":           contentJSON["nivel"],
	}

	where := fmt.Sprintf("id = %v", params.Get("usuarioID"))

	rows, err := postgres.UpdateOne(tableUsuario, values, where)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusNoContent, nil

}

// Delete Adiciona um registro
func Delete(params url.Values) ([]byte, int, error) {

	if params.Get("usuarioID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("id")
	}

	where := fmt.Sprintf("id = %v", params.Get("usuarioID"))

	rows, err := postgres.DeleteOne(tableUsuario, where)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusNoContent, nil

}

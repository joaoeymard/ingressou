package usuario

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JoaoEymard/ingressoscariri/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/usuario/atributo"
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
)

const (
	// Tabela referente ao usuario
	tUsuario = "t_ingressoscariri_usuario"
	// Tabela referente ao contato para o inner JOIN
	tUsuarioContato = "t_ingressoscariri_usuario_contato"
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
		return nil, http.StatusBadRequest, err
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

	rows, err := postgres.InsertOne(tUsuario, values)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, err
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
	sqlTotal := fmt.Sprintf(`SELECT COUNT(USUARIO.id) AS total
	FROM %v USUARIO
	%v`, tUsuario, filter)

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
	sql := fmt.Sprintf(`SELECT USUARIO.id AS id, USUARIO.nome, USUARIO.ultimo_acesso, USUARIO.ativo, USUARIO.cpf, USUARIO.data_nascimento, USUARIO.sexo, USUARIO.nivel,
	(
		SELECT array_to_json (array_agg (row_to_json(dados_contatos.*) ) )
		FROM (
			SELECT USUARIO_CONTATO.id AS contato_id, USUARIO_CONTATO.endereco, USUARIO_CONTATO.complemento, USUARIO_CONTATO.referencia, USUARIO_CONTATO.bairro, USUARIO_CONTATO.cep, USUARIO_CONTATO.cidade, USUARIO_CONTATO.uf, USUARIO_CONTATO.telefone_principal, USUARIO_CONTATO.telefone_secundario, USUARIO_CONTATO.telefone_terciario, USUARIO_CONTATO.email
			FROM %v USUARIO_CONTATO
			WHERE USUARIO_CONTATO.id_usuario = USUARIO.id
		) AS dados_contatos
	) AS contatos
	FROM %v USUARIO
	%v %v %v %v`, tUsuarioContato, tUsuario, filter, order, limit, offset)

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
		return nil, http.StatusBadRequest, err
	}

	return retorno, http.StatusOK, nil
}

// Update Adiciona um registro
func Update(contentBody io.ReadCloser, params url.Values) ([]byte, int, error) {

	var contentJSON map[string]interface{}

	if params.Get("usuarioID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("id")
	}

	urlParams := url.Values{
		"filtro": []string{`{
			"usuarioID":` + params.Get("usuarioID") + `
			}`},
	}

	_, statusCode, err := Find(urlParams)
	if err != nil {
		return nil, statusCode, utils.BancoDados(err)
	}

	content, err := ioutil.ReadAll(contentBody)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(content, &contentJSON)
	if err != nil {
		return nil, http.StatusBadRequest, err
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

	rows, err := postgres.UpdateOne(tUsuario, values, where)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retorno, http.StatusNoContent, nil

}

// Delete Adiciona um registro
func Delete(params url.Values) ([]byte, int, error) {

	if params.Get("usuarioID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("id")
	}

	urlParams := url.Values{
		"filtro": []string{`{
			"usuarioID":` + params.Get("usuarioID") + `
			}`},
	}

	_, statusCode, err := Find(urlParams)
	if err != nil {
		return nil, statusCode, utils.BancoDados(err)
	}

	where := fmt.Sprintf("id = %v", params.Get("usuarioID"))

	rows, err := postgres.DeleteOne(tUsuario, where)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	return retorno, http.StatusNoContent, nil

}

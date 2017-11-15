package evento

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoEymard/ingressou/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressou/api/v1/model/evento/atributo"
	"github.com/JoaoEymard/ingressou/api/v1/utils"
)

const (
	// Tabela referente ao evento
	tableEvento = "t_ingressou_evento"
	// Tabela referente ao periodo
	tableEventoPeriodo = "t_ingressou_evento_periodo"
	// Tabela referente ao categoria
	tablePeriodoCategoria = "t_ingressou_periodo_categoria"
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

	rows, err := postgres.InsertOne(tableEvento, values)
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
	filter, _, _, _, err := postgres.SetParams(params, atributo.Filtros)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Consulta para coletar os registro
	sql := fmt.Sprintf(`SELECT TIEVENTO.id, TIEVENTO.titulo, TIEVENTO.imagem, TIEVENTO.cidade, TIEVENTO.uf, TIEVENTO.localidade, TIEVENTO.taxa, TIEVENTO.mapa, TIEVENTO.descricao, TIEVENTO.link, TIEVENTO.data_criacao, TIEVENTO.status, 
	(
		SELECT array_to_json (array_agg (row_to_json(dados_periodo.*) ) )
		FROM (
			SELECT TIEPERIODO.id, TIEPERIODO.data_periodo AS data, TIEPERIODO.atracao,
			(
				SELECT array_to_json (array_agg (row_to_json(dados_categoria.*) ) )
				FROM (
					SELECT TIPCATEGORIA.id, TIPCATEGORIA.nome, TIPCATEGORIA.valor, TIPCATEGORIA.quantidade, TIPCATEGORIA.lote
					FROM %v TIPCATEGORIA
					WHERE TIPCATEGORIA.id_periodo = TIEPERIODO.id
				) AS dados_categoria
			) AS categorias
			FROM %v TIEPERIODO
			WHERE TIEPERIODO.id_evento = TIEVENTO.id
		) AS dados_periodo
	) AS periodos
	FROM %v TIEVENTO
	%v`, tablePeriodoCategoria, tableEventoPeriodo, tableEvento, filter)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	row, err := postgres.SelectOne(sql)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if row == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	// Montar o json de retorno
	if row["periodos"] != nil {
		var jsonPeriodo []map[string]interface{}
		// Converte a estrutura para json
		err := json.Unmarshal([]byte(row["periodos"].(string)), &jsonPeriodo)
		if err != nil {
			return nil, http.StatusBadRequest, utils.FormatError(err)
		}
		row["periodos"] = jsonPeriodo
	}

	// Converte a estrutura para json
	retorno, err := json.Marshal(row)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusOK, nil
}

// ListarEventos Lista os eventos com poucas informações
func ListarEventos(params url.Values) ([]byte, int, error) {
	// Tratamento dos paramentros e filtro recebidos pela URL
	filter, order, limit, offset, err := postgres.SetParams(params, atributo.Filtros)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	// Consulta para saber o total de registro
	sqlTotal := fmt.Sprintf(`SELECT COUNT(TIEVENTO.id) AS total
	FROM %v TIEVENTO
	%v`, tableEvento, filter)

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
	sql := fmt.Sprintf(`SELECT TIEVENTO.id, TIEVENTO.titulo, TIEVENTO.imagem, (TIEVENTO.cidade || ' - ' || TIEVENTO.uf) AS cidade, TIEVENTO.link,
	(
		SELECT TIEPERIODO_AUX.data_periodo
		FROM %v TIEPERIODO_AUX
		WHERE TIEPERIODO_AUX.id_evento = TIEVENTO.id
		ORDER BY TIEPERIODO_AUX.data_periodo ASC
		LIMIT 1
	)
	FROM %v TIEVENTO
	%v %v %v %v`, tableEventoPeriodo, tableEvento, filter, order, limit, offset)

	// Retorna um []map com as colunas e valores vindo do banco de dados
	rows, err := postgres.Select(sql)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	// Verifica se o retorno está nulo
	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	fmt.Printf("%v\n", rows[0]["data_periodo"].(time.Time).Format("2006_01_02_15_04_05"))
	fmt.Printf("%v\n", rows[1]["data_periodo"].(time.Time).Format("2006_01_02_15_04_05"))

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

	rows, err := postgres.UpdateOne(tableEvento, values, where)
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

	rows, err := postgres.DeleteOne(tableEvento, where)
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

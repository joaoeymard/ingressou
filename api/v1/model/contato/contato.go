package contato

import (
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JoaoEymard/ingressou/api/utils/database/postgres"
	"github.com/JoaoEymard/ingressou/api/v1/model/contato/atributo"
	"github.com/JoaoEymard/ingressou/api/v1/utils"
)

const (
	// Tabela referente ao contato
	tableUsuarioContato = "t_ingressou_usuario_contato"
)

// Insert Adiciona um registro
func Insert(contentBody io.ReadCloser, params url.Values) ([]byte, int, error) {

	var contentJSON map[string]interface{}

	if params.Get("usuarioID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("usuarioID")
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
		"id_usuario":          params.Get("usuarioID"),
		"endereco":            contentJSON["endereco"],
		"complemento":         contentJSON["complemento"],
		"referencia":          contentJSON["referencia"],
		"bairro":              contentJSON["bairro"],
		"cep":                 contentJSON["cep"],
		"cidade":              contentJSON["cidade"],
		"uf":                  contentJSON["uf"],
		"telefone_principal":  contentJSON["telefone_principal"],
		"telefone_secundario": contentJSON["telefone_secundario"],
		"telefone_terciario":  contentJSON["telefone_terciario"],
		"email":               contentJSON["email"],
	}

	rows, err := postgres.InsertOne(tableUsuarioContato, values)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	return retorno, http.StatusCreated, nil

}

// Update Adiciona um registro
func Update(contentBody io.ReadCloser, params url.Values) ([]byte, int, error) {

	var contentJSON map[string]interface{}

	if params.Get("contatoID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("contatoID")
	}

	content, err := ioutil.ReadAll(contentBody)
	if err != nil {
		return nil, http.StatusBadRequest, err
	}

	err = json.Unmarshal(content, &contentJSON)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	values := map[string]interface{}{
		"cep":                 contentJSON["cep"],
		"bairro":              contentJSON["bairro"],
		"cidade":              contentJSON["cidade"],
		"email":               contentJSON["email"],
		"endereco":            contentJSON["endereco"],
		"complemento":         contentJSON["complemento"],
		"referencia":          contentJSON["referencia"],
		"telefone_principal":  contentJSON["telefone_principal"],
		"telefone_secundario": contentJSON["telefone_secundario"],
		"telefone_terciario":  contentJSON["telefone_terciario"],
		"uf":                  contentJSON["uf"],
	}

	where := fmt.Sprintf("id = %v", params.Get("contatoID"))

	rows, err := postgres.UpdateOne(tableUsuarioContato, values, where)
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

	if params.Get("contatoID") == "" {
		return nil, http.StatusBadRequest, utils.ValueRequired("contatoID")
	}

	where := fmt.Sprintf("id = %v", params.Get("contatoID"))

	rows, err := postgres.DeleteOne(tableUsuarioContato, where)
	if err != nil {
		return nil, http.StatusBadRequest, utils.BancoDados(err)
	}

	retorno, err := json.Marshal(rows)
	if err != nil {
		return nil, http.StatusBadRequest, utils.FormatError(err)
	}

	if rows == nil {
		return nil, http.StatusNotFound, utils.Errors["NOT_FOUND"]
	}

	return retorno, http.StatusNoContent, nil

}

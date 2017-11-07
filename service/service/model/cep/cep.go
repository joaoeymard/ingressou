package cep

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"os"
	"sync"
	"time"

	"strconv"

	iris "gopkg.in/kataras/iris.v6"
)

var (
	mutex sync.Mutex

	localFile = "service/model/cep/cache/"

	hostname = []string{
		"https://viacep.com.br/ws/",
	}

	paramFormat = []string{
		"/json",
	}
)

// Find Retorna as informações do CEP especifico
func Find(param int64) (iris.Map, error) {

	cep := strconv.FormatInt(param, 10)

	cache, status := getCache(cep)
	if status == true {
		return cache, nil
	}

	jsonCep, err := requestURL(cep)

	if err != nil {
		return nil, err
	}

	return jsonCep, nil
}

// Retorna o arquivo da pasta cache com as informações do CEP
func getCache(cep string) (iris.Map, bool) {
	var (
		jsonCep iris.Map
	)

	if expiredFile(cep) {
		jsonFile, err := ioutil.ReadFile(localFile + cep + ".json")
		if err != nil {
			return nil, false
		}
		json.Unmarshal(jsonFile, &jsonCep)

		return jsonCep, true
	}

	return nil, false
}

// Retorna se o arquivo está desatualizado ou não
func expiredFile(cep string) bool {
	file, err := os.Stat(localFile + cep + ".json")

	if err != nil {
		return false
	}

	expireTime := file.ModTime().Add(24 * time.Hour)

	return !time.Now().After(expireTime)
}

// Retorna o json com as informações do CEP via requisição ao ViaCep
func requestURL(cep string) (iris.Map, error) {
	var (
		response *http.Response
		err      error
		jsonCep  iris.Map
	)

	transport := &http.Transport{
		MaxIdleConns:       10,
		IdleConnTimeout:    10 * time.Second,
		DisableCompression: true,
	}

	client := &http.Client{Transport: transport}

	for i := 0; i < len(hostname); i++ {
		response, err = client.Get(hostname[i] + cep + paramFormat[i])
		if err == nil {
			break
		} else if i+1 == len(hostname) {
			return nil, errors.New("CEP não encontrado")
		}
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.New("Ler conteúdo do body")
	}

	json.Unmarshal(body, &jsonCep)

	if jsonCep == nil || jsonCep["erro"] != nil {
		return nil, errors.New("CEP não encontrado")
	}

	saveCache(cep, body)

	return jsonCep, nil

}

// Salva o arquivo na pasta cache
func saveCache(cep string, content []byte) error {
	err := ioutil.WriteFile(localFile+cep+".json", content, 0644)

	if err != nil {
		return errors.New("Arquivo nao encontrado")
	}

	return nil
}

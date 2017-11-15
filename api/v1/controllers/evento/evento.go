package evento

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoEymard/ingressou/api/utils/logger"
	"github.com/JoaoEymard/ingressou/api/v1/model/evento"
	"github.com/JoaoEymard/ingressou/api/v1/utils"
	"github.com/gorilla/mux"
)

// Methods Divisão de rotas para gerenciar o controller
func Methods(res http.ResponseWriter, req *http.Request) {

	var (
		retornoDados []byte
		statusCode   int
		err          error
	)

	begin := time.Now().UTC()

	res.Header().Set("Content-Type", "application/json")

	switch req.Method {

	case "POST":
		retornoDados, statusCode, err = evento.Insert(req.Body)

	case "GET":
		var urlParams url.Values

		urlParams = url.Values{
			"filtro": []string{`{
					"eventoLink": "` + mux.Vars(req)["link"] + `"
					}`},
		}

		retornoDados, statusCode, err = evento.Find(urlParams)

	case "PUT":
		urlParams := url.Values{
			"eventoID": []string{mux.Vars(req)["eventoID"]},
		}

		retornoDados, statusCode, err = evento.Update(req.Body, urlParams)

	case "DELETE":
		urlParams := url.Values{
			"eventoID": []string{mux.Vars(req)["eventoID"]},
		}

		retornoDados, statusCode, err = evento.Delete(urlParams)

	default:
		retornoDados, statusCode, err = nil, http.StatusNotFound, utils.Errors["METHOD_DEFAULT"]

	}

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		logger.Warnln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
		return
	}

	res.Write(retornoDados)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))

}

// ListarEventos Rota para retornar apenas informações básicas
func ListarEventos(res http.ResponseWriter, req *http.Request) {

	var (
		retornoDados []byte
		statusCode   int
		err          error
	)

	begin := time.Now().UTC()

	res.Header().Set("Content-Type", "application/json")

	retornoDados, statusCode, err = evento.ListarEventos(req.URL.Query())

	res.WriteHeader(statusCode)
	if err != nil {
		res.Write([]byte(err.Error()))
		logger.Warnln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
		return
	}

	res.Write(retornoDados)

	logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))

}

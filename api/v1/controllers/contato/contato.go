package contato

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoEymard/ingressou/api/utils/logger"
	"github.com/JoaoEymard/ingressou/api/v1/model/contato"
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
		urlParams := url.Values{
			"usuarioID": []string{mux.Vars(req)["usuarioID"]},
		}

		retornoDados, statusCode, err = contato.Insert(req.Body, urlParams)

	case "PUT":
		urlParams := url.Values{
			"usuarioID": []string{mux.Vars(req)["usuarioID"]},
			"contatoID": []string{mux.Vars(req)["contatoID"]},
		}

		retornoDados, statusCode, err = contato.Update(req.Body, urlParams)

	case "DELETE":
		urlParams := url.Values{
			"usuarioID": []string{mux.Vars(req)["usuarioID"]},
			"contatoID": []string{mux.Vars(req)["contatoID"]},
		}

		retornoDados, statusCode, err = contato.Delete(urlParams)

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

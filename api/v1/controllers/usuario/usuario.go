package usuario

import (
	"fmt"
	"net/http"
	"net/url"
	"time"

	"github.com/JoaoEymard/ingressoscariri/api/utils/logger"
	"github.com/JoaoEymard/ingressoscariri/api/v1/model/usuario"
	"github.com/JoaoEymard/ingressoscariri/api/v1/utils"
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
		retornoDados, statusCode, err = usuario.Insert(req.Body)

	case "GET":
		var urlParams url.Values

		if mux.Vars(req)["usuarioID"] == "" {
			urlParams = req.URL.Query()
		} else {
			urlParams = url.Values{
				"filtro": []string{`{
					"usuarioID":` + mux.Vars(req)["usuarioID"] + `
					}`},
			}
		}

		retornoDados, statusCode, err = usuario.Find(urlParams)

	case "PUT":
		urlParams := url.Values{
			"usuarioID": []string{mux.Vars(req)["usuarioID"]},
		}

		retornoDados, statusCode, err = usuario.Update(req.Body, urlParams)

	case "DELETE":
		urlParams := url.Values{
			"usuarioID": []string{mux.Vars(req)["usuarioID"]},
		}

		retornoDados, statusCode, err = usuario.Delete(urlParams)

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

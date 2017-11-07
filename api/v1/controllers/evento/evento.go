package evento

import (
	"fmt"
	"net/http"
	"time"

	"github.com/JoaoEymard/ingressou/api/utils/logger"
	"github.com/JoaoEymard/ingressou/api/v1/utils"
)

// Methods Divisão de rotas para gerenciar o controller
func Methods(res http.ResponseWriter, req *http.Request) {

	var (
		retornoDados []byte
		statusCode   int
		err          error
	)

	begin := time.Now().UTC()

	switch req.Method {

	case "POST":
		// retornoDados, statusCode, err = evento.Insert(req.Body)

	case "GET":
		// var urlParams url.Values

		// if mux.Vars(req)["id"] == "" {
		// 	urlParams = req.URL.Query()
		// } else {
		// 	urlParams = url.Values{
		// 		"filtro": []string{`{"id":` + mux.Vars(req)["id"] + `}`},
		// 	}
		// }

		// retornoDados, statusCode, err = evento.Find(urlParams)

	case "PUT":
		// urlParams := url.Values{
		// 	"id": []string{mux.Vars(req)["id"]},
		// }

		// retornoDados, statusCode, err = evento.Update(req.Body, urlParams)

	case "DELETE":
		// urlParams := url.Values{
		// 	"id": []string{mux.Vars(req)["id"]},
		// }

		// retornoDados, statusCode, err = evento.Delete(urlParams)

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

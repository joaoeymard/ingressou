package api

import (
	"github.com/JoaoEymard/ingressou/api/v1"
	"github.com/gorilla/mux"
)

// Routes pacotes das rotas
func Routes(app *mux.Router) {

	apiParty := app.PathPrefix("/api").Subrouter()
	{

		v1Party := apiParty.PathPrefix("/v1").Subrouter()
		{
			v1.ConfigRoutes(v1Party)
		}

	}

}

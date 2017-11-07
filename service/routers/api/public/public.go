package public

import (
	ctrlCep "github.com/JoaoEymard/ingressou/service/controllers/cep"
	ctrlCidades "github.com/JoaoEymard/ingressou/service/controllers/cidades"
	ctrlMapa "github.com/JoaoEymard/ingressou/service/controllers/mapa"

	utils "github.com/JoaoEymard/ingressou/service/utils"
	iris "gopkg.in/kataras/iris.v6"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(router *iris.Router) {

	// Mapa
	router.Get("/map", ctrlMapa.Find)

	// CEP
	router.Get("/cep/{cep:"+utils.Regex["integer"]+"}", ctrlCep.Find)

	router.Get("/cidades", ctrlCidades.FindAll)

}

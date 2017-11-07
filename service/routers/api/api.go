/*
package api

import (
	ctrlAgendamento "github.com/Brisanet/site_brisanet_go/controllers/agendamento"
	ctrlCep "github.com/Brisanet/site_brisanet_go/controllers/cep"
	ctrlCidades "github.com/Brisanet/site_brisanet_go/controllers/cidades"
	ctrlCombo "github.com/Brisanet/site_brisanet_go/controllers/combo"
	ctrlMapa "github.com/Brisanet/site_brisanet_go/controllers/mapa"
	ctrlServicos "github.com/Brisanet/site_brisanet_go/controllers/servicos"
	ctrlTv "github.com/Brisanet/site_brisanet_go/controllers/tv"
	ctrlVendas "github.com/Brisanet/site_brisanet_go/controllers/vendas"
	utils "github.com/Brisanet/site_brisanet_go/utils"

	iris "gopkg.in/kataras/iris.v6"
)

// Routes Rotas publicas
func Routes(router *iris.Router) {
	// Cidades
	router.Post("/cidades", ctrlCidades.Create)
	router.Put("/cidades/{id:"+utils.Regex["integer"]+"}", ctrlCidades.Modify)
	router.Get("/cidades/{id:"+utils.Regex["integer"]+"}/servicos/:type", ctrlCidades.Find) // CORRIGIR
	router.Get("/cidades", ctrlCidades.FindAll)

	// Mapa
	router.Get("/map", ctrlMapa.Find)

	//Combo
	router.Get("/combo/planos-telefone", ctrlCombo.FindAllTelefone)
	router.Get("/combo/plano-telefone/detalhes/{planoID:"+utils.Regex["integer"]+"}", ctrlCombo.FindTelefone)
	router.Get("/combo/valores/{internetID:"+utils.Regex["integer"]+"}", ctrlCombo.FindValores)

	//Tv
	router.Get("/tv/canais/:plano", ctrlTv.FindCanais) // CORRIGIR

	//Vendas
	router.Post("/venda", ctrlVendas.Create)
	router.Put("/vendas/:venda/mudar-status/:status", ctrlVendas.Modify) // CORRIGIR
	router.Get("/vendas", ctrlVendas.Find)

	// Serviços
	router.Post("/servicos", ctrlServicos.Create)
	router.Put("/servicos/{id:"+utils.Regex["integer"]+"}", ctrlServicos.Modify)
	router.Get("/servico/{servicoID:"+utils.Regex["integer"]+"}/info", ctrlServicos.Find)

	router.Get("/servicos/internet/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindInternet)
	router.Get("/servicos/internet/plano/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindInternetPlano)
	router.Get("/servicos/telefone/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindTelefone)
	router.Get("/servicos/telefone/plano/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindTelefonePlano)
	router.Get("/servicos/tv/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindTv)
	router.Get("/servicos/tv/plano/{id:"+utils.Regex["integer"]+"}", ctrlServicos.FindTvPlano)

	// Agendamento
	router.Post("/agendamentos", ctrlAgendamento.Create)
	// router.Post("/agendamentos_test", ctrlAgendamento.CreateTeste)

	// Cep
	router.Get("/cep/{cep:"+utils.Regex["integer"]+"}", ctrlCep.Find)
}
*/
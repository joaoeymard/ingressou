package v1

import (
	ctrlAuth "github.com/JoaoEymard/ingressoscariri/api/v1/controllers/auth"
	"github.com/JoaoEymard/ingressoscariri/api/v1/controllers/contato"
	"github.com/JoaoEymard/ingressoscariri/api/v1/controllers/evento"
	"github.com/JoaoEymard/ingressoscariri/api/v1/controllers/usuario"
	"github.com/gorilla/mux"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(route *mux.Router) {

	// Usuarios
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}", usuario.Methods).Methods("GET", "PUT", "DELETE")
	route.HandleFunc("/usuario/", usuario.Methods).Methods("POST", "GET")
	route.HandleFunc("/usuario", usuario.Methods).Methods("POST", "GET")

	// Contatos
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}/contato/{contatoID:[0-9]+}", contato.Methods).Methods("GET", "PUT", "DELETE")
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}/contato/", contato.Methods).Methods("POST", "GET")
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}/contato", contato.Methods).Methods("POST", "GET")

	// Eventos
	route.HandleFunc("/evento", evento.Methods).Methods("POST", "GET")
	route.HandleFunc("/evento/{link:[a-zA-Z0-9_]+?}", evento.Methods).Methods("GET", "PUT", "DELETE")

	// Teste
	route.HandleFunc("/withAuth", ctrlAuth.Check).Methods("GET")

}

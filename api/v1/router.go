package v1

import (
	"net/http"

	ctrlAuth "github.com/JoaoEymard/ingressou/api/v1/controllers/auth"
	"github.com/JoaoEymard/ingressou/api/v1/controllers/contato"
	"github.com/JoaoEymard/ingressou/api/v1/controllers/evento"
	"github.com/JoaoEymard/ingressou/api/v1/controllers/usuario"
	"github.com/gorilla/mux"
)

// ConfigRoutes Tratamento das Rotas publicas
func ConfigRoutes(route *mux.Router) {

	// Usuarios
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}", usuario.Methods).Methods("GET", "PUT", "DELETE")
	route.HandleFunc("/usuario", usuario.Methods).Methods("POST", "GET")

	// Contatos
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}/contato/{contatoID:[0-9]+}", contato.Methods).Methods("GET", "PUT", "DELETE")
	route.HandleFunc("/usuario/{usuarioID:[0-9]+}/contato", contato.Methods).Methods("POST", "GET")

	// Eventos
	route.HandleFunc("/evento", evento.Methods).Methods("POST")
	route.HandleFunc("/evento/lista", evento.ListarEventos).Methods("GET")
	route.HandleFunc("/evento/{link:[a-zA-Z0-9_]+?}", evento.Methods).Methods("GET")

	// Teste
	route.HandleFunc("/withAuth", ctrlAuth.Check).Methods("GET")

	// Imagens
	route.PathPrefix("/image/").Handler(http.StripPrefix("/api/v1/image/", http.FileServer(http.Dir("api/utils/database/image/"))))

}

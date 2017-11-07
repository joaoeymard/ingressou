package authentication

import (
	auth "github.com/JoaoEymard/ingressou/service/controllers/authentication"
	"gopkg.in/kataras/iris.v6"
)

func ConfigRoutes(router *iris.Router) {
	router.Post("/login", auth.Login)
	router.Get("/check", auth.CheckLogin)
	router.Put("/logout", auth.Logout)
}

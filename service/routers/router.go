package routers

import (
	"github.com/JoaoEymard/ingressou/service/routers/api/admin"
	"github.com/JoaoEymard/ingressou/service/routers/api/public"
	"gopkg.in/kataras/iris.v6"
)

// Routes pacotes das rotas
func Routes(app *iris.Framework) {

	apiParty := *app.Party("/api")
	{
		publicParty := apiParty.Party("/public")
		public.ConfigRoutes(publicParty)

		adminParty := apiParty.Party("/admin")
		admin.ConfigRoutes(adminParty)

		apiParty.Get("/", func(ctx *iris.Context) {
			ctx.JSON(200, map[string]string{"api": "teste"})
		})
	}

}

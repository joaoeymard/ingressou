package auth

import (
	"net/http"
)

func Check(res http.ResponseWriter, req *http.Request) {
	// sess := session.GetSession().Start(ctx)
	// value, _ := sess.GetBoolean("sessionCript")
	// if value {
	// 	ctx.Next()
	// }
	// ctx.StatusCode(iris.StatusForbidden)

	// begin := time.Now().UTC()

	// store, err := session.GetSession().Get(req, "teste")
	// if err != nil {
	// 	logger.Errorln(err)
	// }
	// logger.Debugln(store.ID)

	// res.Write([]byte("Teste"))

	// logger.Infoln(logger.Status(fmt.Sprintf("%+v\n", res)), req.RemoteAddr, req.Method, req.URL, time.Now().UTC().Sub(begin))
}

func Login(res http.ResponseWriter, req *http.Request) {

	// if ctx.Params().Get("user") == "c019" && ctx.Params().Get("passw") == "admin" {
	// 	sess := session.GetSession().Start(ctx)
	// 	sess.Set("sessionCript", true)
	// 	ctx.JSON(map[string]interface{}{"Login": "OK"})
	// 	return
	// }
	// ctx.JSON(map[string]interface{}{"Login": "Incorreto"})

}

func Logout(res http.ResponseWriter, req *http.Request) {
	// sess := session.GetSession().Start(ctx)
	// if sess.Get("sessionCript") != nil {
	// 	sess.Delete("sessionCript")
	// 	ctx.JSON(map[string]interface{}{"Logout": "OK"})
	// 	return
	// }
	// ctx.JSON(map[string]interface{}{"Logout": "Nil"})
}

package middleware

import (
	"net/http"
)

func Check(res http.ResponseWriter, req *http.Request) {
	// if ctrlAuth.Check(ctx) {
	// 	ctx.Next()
	// }
	// ctx.StatusCode(iris.StatusForbidden)
	// return
}

// func Cors(res http.ResponseWriter, req *http.Request) {

// 	ctx.Header("Access-Control-Allow-Origin", "*")
// 	ctx.Header("Access-Control-Allow-Headers", "Content-Type, X-Requested-With")
// 	ctx.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

// 	ctx.Next()
// }

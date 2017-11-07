package session

import (
	"github.com/JoaoEymard/ingressou/api/utils/settings"
	"github.com/gorilla/sessions"
)

var (
	cookieName = []byte(settings.GetSettings().CookieName)
	hashKey    = []byte(settings.GetSettings().HashKey)
	blockKey   = []byte(settings.GetSettings().BlockKey)

	store *sessions.CookieStore
)

func init() {

	store = sessions.NewCookieStore(cookieName, hashKey, blockKey)
}

// SetSession Adiciona os valores para uma sessão
func SetSession(domain, path string, maxAge int, httpOnly, secure bool) {
	store.Options = &sessions.Options{
		Domain:   "",
		Path:     "/",
		MaxAge:   5,
		HttpOnly: true,
		Secure:   true,
	}
}

// GetSession Retorna a sessions criadas
func GetSession() *sessions.CookieStore {
	return store
}

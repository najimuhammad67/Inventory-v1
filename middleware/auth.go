package middleware

import (
	"github.com/gorilla/sessions"
	"net/http"
)

var Store = sessions.NewCookieStore([]byte("super-secret-key"))

func RequireLogin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		if session.Values["username"] == nil {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

func RequireAdmin(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session, _ := Store.Get(r, "session")
		if session.Values["username"] == nil || session.Values["role"] != "admin" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

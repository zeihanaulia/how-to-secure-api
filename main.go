package main

import (
	"net/http"

	"fmt"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

func main() {
	r := chi.NewRouter()

	tokenAuth := jwtauth.New("HS256", []byte("secret"), nil)

	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(tokenAuth))
		r.Use(jwtauth.Authenticator)
		r.Get("/admin", func(w http.ResponseWriter, r *http.Request) {
			_, claims, _ := jwtauth.FromContext(r.Context())
			w.Write([]byte(fmt.Sprintf("protected area. hi %v", claims["user_id"])))
		})
	})

	r.Group(func(r chi.Router) {
		r.Get("/", func(w http.ResponseWriter, r *http.Request) {
			_, tokenString, _ := tokenAuth.Encode(jwtauth.Claims{"user_id": 123})
			w.Write([]byte(fmt.Sprintf("proof of concept securing with jwt - public api, token: %s", tokenString)))
		})
	})
	http.ListenAndServe(":3001", r)
}

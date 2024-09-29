package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"rachitmishra.com/pastebin/cmd/web/create"
	"rachitmishra.com/pastebin/cmd/web/home"
	"rachitmishra.com/pastebin/cmd/web/login"
	"rachitmishra.com/pastebin/cmd/web/view"
)

func (a *Application) routeHandler() http.Handler {
	mux := func() *httprouter.Router {
		// mux := http.NewServeMux()
		// fileServer := http.FileServer(
		// 	http.Dir("./ui/static/"))
		// mux.Handle("/static/", http.StripPrefix("/static", fileServer))
		// mux.HandleFunc("/view", view.Handler(a))
		// mux.HandleFunc("/create", create.Handler(a))

		router := httprouter.New()
		router.NotFound = http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				a.NotFound(w)
			})
		router.HandlerFunc(http.MethodGet, "/", home.Handler(a))
		router.HandlerFunc(http.MethodGet, "/paste/:id", view.Handler(a))
		router.HandlerFunc(http.MethodPost, "/pastes/new", create.Handler(a))
		router.HandlerFunc(http.MethodGet, "/pastes", home.Handler(a))
		router.HandlerFunc(http.MethodGet, "/user", login.Handler(a))
		router.HandlerFunc(http.MethodPost, "/session/new", login.Handler(a))
		router.HandlerFunc(http.MethodDelete, "/session", login.Handler(a))
		return router
	}()
	return a.recoverPanic(a.logRequest(secureHeaders(mux)))
}

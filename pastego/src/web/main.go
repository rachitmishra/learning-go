package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-playground/form/v4"
	"github.com/redis/go-redis/v9"
	"rachitmishra.com/pastebin/internal/models"
)

func newApplication() *Application {
	formDecoder := form.NewDecoder()
	errorLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	return &Application{
		loggerE: errorLogger,
		loggerI: infoLogger,
		db: models.NewDB(
			redis.NewClient(&redis.Options{
				Addr:     "localhost:6379",
				Password: "", // no password set
				DB:       0,  // use default DB
			}),
		),
		formDecoder: formDecoder,
	}
}

func (a *Application) newServer(addr string) *http.Server {
	return &http.Server{
		Addr:     addr,
		ErrorLog: a.loggerE,
		Handler:  a.routeHandler(),
	}
}

func main() {
	port := flag.String("addr", "4000", "Http network address")
	app := newApplication()
	flag.Parse()
	addr := fmt.Sprintf(":%s", *port)
	srv := app.newServer(addr)
	app.LogI(fmt.Sprintf("Starting server on %s", addr))
	err := srv.ListenAndServe()
	app.loggerE.Fatal(err)
}

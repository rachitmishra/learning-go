package main

import (
	"crypto/tls"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/go-playground/form/v4"
	"github.com/redis/go-redis/v9"
	"rachitmishra.com/pastebin/internal/models"
)

// setp traefic + go
// set p1 redis + postgres
func newApplication() *Application {
	formDecoder := form.NewDecoder()
	errorLogger := log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	infoLogger := log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)

	redis := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	// sessionManager := scs.New()
	// sessionManager.Store = redisstore.New(redis)

	return &Application{
		loggerE:     errorLogger,
		loggerI:     infoLogger,
		db:          models.NewDB(redis),
		formDecoder: formDecoder,
	}
}

func (a *Application) newServer(addr string) *http.Server {
	return &http.Server{
		Addr:         addr,
		ErrorLog:     a.loggerE,
		Handler:      a.routeHandler(),
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		TLSConfig: &tls.Config{
			MinVersion:               tls.VersionTLS13,
			PreferServerCipherSuites: true,
		},
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

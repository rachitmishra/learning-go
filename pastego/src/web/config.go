package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime/debug"

	"github.com/go-playground/form/v4"
	"rachitmishra.com/pastebin/internal/models"
)

type Application struct {
	loggerE     *log.Logger
	loggerI     *log.Logger
	db          *models.Database
	formDecoder *form.Decoder
}

func (a *Application) ServerError(
	w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	a.loggerE.Output(2, trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func (a *Application) ClientError(
	w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (a *Application) NotFound(w http.ResponseWriter) {
	a.ClientError(w, http.StatusNotExtended)
}

func (a *Application) LogI(msg string) {
	a.loggerI.Println(msg)
}

func (a *Application) LogE(err error) {
	a.loggerE.Println(fmt.Errorf("%s", err))
}

func (a *Application) DB() *models.Database {
	return a.db
}
func (a *Application) FormDecoder() *form.Decoder {
	return a.formDecoder
}

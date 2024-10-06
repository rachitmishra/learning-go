package shared

import (
	"net/http"

	"github.com/go-playground/form/v4"
	"rachitmishra.com/pastebin/internal/models"
)

type ViewDeps interface {
	ServerError(w http.ResponseWriter, err error)
	ClientError(w http.ResponseWriter, status int)
	LogI(msg string)
	LogE(err error)
}

type ServiceDeps interface {
	ServerError(w http.ResponseWriter, err error)
	ClientError(w http.ResponseWriter, status int)
	LogI(msg string)
	LogE(err error)
	DB() *models.Database
	FormDecoder() *form.Decoder
}

type Service interface {
}

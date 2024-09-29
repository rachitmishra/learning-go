package handlers

import (
	"encoding/json"
	"log/slog"
	"net/http"

	"rachitmishra.com/learning-go/microservices/types"
)

type Products struct {
}

func NewProducts() *Products {
	return &Products{}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		p.postProducts(rw, r)
		return
	}

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}
	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) postProducts(rw http.ResponseWriter, r *http.Request) {
	product := types.Product{}
	body := json.NewDecoder(r.Body)
	err := body.Decode(&product)
	if err != nil {
		http.Error(rw, "Failed", http.StatusInternalServerError)
		return
	}
	slog.Info(product.Roast)
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	EncodeAndServe(rw)
}

func EncodeAndServe(rw http.ResponseWriter) {
	pl := types.GetProductList()
	err := pl.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed", http.StatusInternalServerError)
		return
	}
}

func MarshalAndServe(rw http.ResponseWriter) {
	v, err := json.Marshal(types.GetProductList())
	if err != nil {
		http.Error(rw, "Failed", http.StatusInternalServerError)
		return
	}
	rw.Write(v)
}

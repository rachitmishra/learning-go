package handlers

import (
	"encoding/json"
	"fmt"
	"log/slog"
	"net/http"
	"regexp"
	"strconv"

	"rachitmishra.com/learning-go/microservices-with-go/data"
)

type Products struct {
	l *slog.Logger
}

func NewProducts(l *slog.Logger) *Products {
	return &Products{l}
}

func (p *Products) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPut {
		p.updateProduct(rw, r)
		return
	}

	if r.Method == http.MethodPost {
		p.createProduct(rw, r)
		return
	}

	if r.Method == http.MethodGet {
		p.getProducts(rw, r)
		return
	}

	rw.WriteHeader(http.StatusMethodNotAllowed)
}

func (p *Products) updateProduct(rw http.ResponseWriter, r *http.Request) {
	rg := regexp.MustCompile(`/([0-9)]+)`)
	cg := rg.FindAllStringSubmatch(r.URL.Path, -1)
	if len(cg) != 1 {
		p.l.Error("Invalid URI more than one id")
		http.Error(rw, "Invalid URI more than one id", http.StatusBadRequest)
		return
	}

	if len(cg[0]) != 2 {
		p.l.Error("Invalid URI more than one capture group")
		http.Error(rw, "Invalid URI more than one capture group", http.StatusBadRequest)
		return
	}

	idString := cg[0][1]
	id, cErr := strconv.Atoi(idString)
	if cErr != nil {
		p.l.Error("Invalid URI unable to parse number")
		http.Error(rw, "Invalid URI unable to parse number", http.StatusBadRequest)
		return
	}

	p.l.Info(fmt.Sprintf("%d", id))

	pd := data.Product{}
	pErr := pd.FromJSON(r.Body)
	if pErr != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}

	res, uErr := data.UpdateProduct(int32(id), &pd)
	if uErr != nil && uErr == data.ErrorProductNotFound {
		http.Error(rw, "product not found", http.StatusNotFound)
		return
	}
	p.EncodeAndServe(rw, res)
}

func (p *Products) createProduct(rw http.ResponseWriter, r *http.Request) {
	pd := data.Product{}
	err := pd.FromJSON(r.Body)
	if err != nil {
		http.Error(rw, "unable to unmarshal json", http.StatusBadRequest)
		return
	}
	p.EncodeAndServe(rw, data.AddProduct(&pd))
}

func (p *Products) getProducts(rw http.ResponseWriter, r *http.Request) {
	p.EncodeAndServe(rw, data.GetProducts())
}

func (p *Products) EncodeAndServe(rw http.ResponseWriter, pd data.Products) {
	err := pd.ToJSON(rw)
	if err != nil {
		http.Error(rw, "Failed", http.StatusInternalServerError)
		return
	}
}

func (p *Products) MarshalAndServe(rw http.ResponseWriter) {
	v, err := json.Marshal(data.GetProducts())
	if err != nil {
		http.Error(rw, "Failed", http.StatusInternalServerError)
		return
	}
	rw.Write(v)
}

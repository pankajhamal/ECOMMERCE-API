package products

import (
	"log"
	"net/http"

	"github.com/pankajhamal/ECOMMERCE-API/internal/json"
)

type handler struct {
	service Service
}

// This is constructor
func NewHandler(service Service) *handler{
	return &handler{
		service: service,
	}
}

func (h *handler) ListProducts(w http.ResponseWriter, r *http.Request) {
	//1. CALL the service --> ListProduct
	err := h.service.ListProducts(r.Context())
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	//2. Return JSON in an HTTP response
	products := []string{"Hellow", "World"}
	
	json.Write(w, http.StatusOK, products)
}
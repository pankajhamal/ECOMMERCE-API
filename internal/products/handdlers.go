package products

import (
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
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
	products, err := h.service.ListProducts(r.Context())
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	
	json.Write(w, http.StatusOK, products)
}

func (h *handler) FindProductByID(w http.ResponseWriter, r *http.Request){

	idStr := chi.URLParam(r, "id")

	id, err := strconv.ParseInt(idStr, 10, 64)
    if err != nil {
        http.Error(w, "invalid product id", http.StatusBadRequest)
        return
    }
		 
	// Call the service for find product by id
	product, err := h.service.FindProductByID(r.Context(), id)
	if err != nil{
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	json.Write(w, http.StatusOK, product)
}
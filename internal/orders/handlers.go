package orders

import (
	"log"
	"net/http"

	"github.com/pankajhamal/ECOMMERCE-API/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) PlaceOrder(w http.ResponseWriter, r *http.Request){

	var tempOrder CreateOrderParams
	if err := json.Read(r, &tempOrder); err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	createdOrder, err := h.service.PlaceOrder(r.Context(), tempOrder)
	if err != nil{
		log.Println(err)
		if err == ErrProductNotFound{
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return

		}
	}

	// h.service.PlaceOrder(r.Context(), tempOrder)
	json.Write(w, http.StatusCreated, createdOrder)
}
package usecase

import (
	"log"
	"net/http"

	"github.com/nicnuyster/Effective_Mobile_Junior_GO/internal/json"
)

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListAll(w http.ResponseWriter, r *http.Request) {
	err := h.service.ListAll(r.Context())

	if err != nil {
		log.Println(err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	subscribtions := struct {
		sub_name string
	}{}

	json.write(w, http.StatusOK, subscribtions)

}

package usecase

import "net/http"

type handler struct {
	service Service
}

func NewHandler(service Service) *handler {
	return &handler{
		service: service,
	}
}

func (h *handler) ListAll(w http.ResponseWriter, r *http.Request) {

}

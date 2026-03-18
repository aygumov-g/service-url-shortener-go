package get_link

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	get_linkUC GetLinkUsecase
	domain     string
}

func NewHandler(get_linkUC GetLinkUsecase, domain string) *handler {
	return &handler{
		get_linkUC: get_linkUC,
		domain:     domain,
	}
}

func (h *handler) Execute(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	link, err := h.get_linkUC.Execute(r.Context(), code)
	if err != nil {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	var resp response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.toResponse(link, code, h.domain))
}

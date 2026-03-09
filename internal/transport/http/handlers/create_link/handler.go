package create_link

import (
	"encoding/json"
	"errors"
	"net/http"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type handler struct {
	create_linkUC CreateLinkUsecase
	domain        string
}

func NewHandler(create_linkUC CreateLinkUsecase, domain string) *handler {
	return &handler{
		create_linkUC: create_linkUC,
		domain:        domain,
	}
}

func (h *handler) Execute(w http.ResponseWriter, r *http.Request) {
	var req request
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "invalid request", http.StatusBadRequest)
		return
	}

	code, err := h.create_linkUC.Execute(req.URL)
	if err != nil {
		switch {
		case errors.Is(err, link_d.ErrCannotShortenLink):
			http.Error(w, "cannot shorten link", http.StatusConflict)
		case errors.Is(err, link_d.ErrLinkNotFound):
			http.Error(w, "link not found", http.StatusNotFound)
		case errors.Is(err, link_d.ErrUrlToLong):
			http.Error(w, "url to long", http.StatusBadRequest)
		default:
			http.Error(w, "internal error", http.StatusInternalServerError)
		}

		return
	}

	var resp response
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.toResponse(code, h.domain))
}

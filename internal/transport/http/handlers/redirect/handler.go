package redirect

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	redirectUC RootUsecase
}

func NewHandler(redirectUC RootUsecase) *handler {
	return &handler{redirectUC: redirectUC}
}

func (h *handler) Execute(w http.ResponseWriter, r *http.Request) {
	link, err := h.redirectUC.Execute(chi.URLParam(r, "code"))
	if err != nil {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}

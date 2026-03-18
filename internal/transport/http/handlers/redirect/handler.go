package redirect

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type handler struct {
	get_linkUC    GetLinkUsecase
	update_linkUC UpdateLinkUsecase
}

func NewHandler(get_linkUC GetLinkUsecase, update_linkUC UpdateLinkUsecase) *handler {
	return &handler{
		get_linkUC:    get_linkUC,
		update_linkUC: update_linkUC,
	}
}

func (h *handler) Execute(w http.ResponseWriter, r *http.Request) {
	code := chi.URLParam(r, "code")
	link, err := h.get_linkUC.Execute(r.Context(), code)
	if err != nil {
		http.Error(w, "link not found", http.StatusNotFound)
		return
	}

	if err := h.update_linkUC.Execute(r.Context(), link.ID); err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, link.OriginalURL, http.StatusFound)
}

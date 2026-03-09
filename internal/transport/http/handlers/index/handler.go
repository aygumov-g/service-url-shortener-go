package index

import (
	"io/fs"
	"net/http"

	"github.com/aygumov-g/service-url-shortener-go/web/embed"
)

type handler struct{}

func NewHandler() *handler {
	return &handler{}
}

func (h *handler) Execute(w http.ResponseWriter, r *http.Request) {
	data, err := fs.ReadFile(embed.Public, "html/index.html")
	if err != nil {
		http.Error(w, "internal error", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	w.Write(data)
}

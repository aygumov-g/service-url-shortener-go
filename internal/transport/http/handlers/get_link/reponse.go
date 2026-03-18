package get_link

import (
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type response struct {
	OriginalURL    string  `json:"original_url"`
	ShortURL       string  `json:"short_url"`
	ShortCode      string  `json:"short_code"`
	CustomCode     *string `json:"custom_code"`
	Domain         string  `json:"domain"`
	ClickCount     int64   `json:"clicks"`
	LastAccessedAt *int64  `json:"last_accessed_at"`
	CreatedAt      int64   `json:"created_at"`
}

func (r response) toResponse(link *link_d.Link, code string, domain string) *response {
	return &response{
		OriginalURL:    link.OriginalURL,
		ShortURL:       "https://" + domain + "/" + code,
		ShortCode:      code,
		CustomCode:     link.CustomCode,
		Domain:         domain,
		ClickCount:     link.ClickCount,
		LastAccessedAt: toUnixPtr(link.LastAccessedAt),
		CreatedAt:      link.CreatedAt.Unix(),
	}
}

func toUnixPtr(t *time.Time) *int64 {
	if t == nil {
		return nil
	}

	v := t.Unix()
	return &v
}

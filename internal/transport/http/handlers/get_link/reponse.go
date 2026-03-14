package get_link

import (
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type response struct {
	OriginalURL    string     `json:"original_url"`
	ShortURL       string     `json:"short_url"`
	ShortCode      string     `json:"short_code"`
	CustomCode     *string    `json:"custom_code"`
	ClickCount     int64      `json:"clicks"`
	LastAccessedAt *time.Time `json:"last_accessed_at"`
	CreatedAt      time.Time  `json:"created_at"`
}

func (r response) toResponse(link *link_d.Link, code string, domain string) *response {
	return &response{
		OriginalURL:    link.OriginalURL,
		ShortURL:       "https://" + domain + "/" + code,
		ShortCode:      code,
		CustomCode:     link.CustomCode,
		ClickCount:     link.ClickCount,
		LastAccessedAt: link.LastAccessedAt,
		CreatedAt:      link.CreatedAt,
	}
}

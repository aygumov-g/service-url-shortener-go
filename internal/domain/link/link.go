package link

import "time"

type Link struct {
	ID             int64
	OriginalURL    string
	CustomCode     *string
	ClickCount     int64
	LastAccessedAt *time.Time
	CreatedAt      time.Time
}

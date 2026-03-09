package link

import "time"

type Link struct {
	ID             int64   `gorm:"primaryKey"`
	OriginalURL    string  `gorm:"not null"`
	CustomCode     *string `gorm:"uniqueIndex"`
	ClickCount     int64   `gorm:"not null;default:0"`
	LastAccessedAt *time.Time
	CreatedAt      time.Time
}

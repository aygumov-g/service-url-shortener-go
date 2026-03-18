package link

import (
	"context"
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type LinkRepository interface {
	Create(ctx context.Context, link *link_d.Link) error
	GetByID(ctx context.Context, id int64) (*link_d.Link, error)
}

type Generator interface {
	Encode(id int64) (string, error)
}

type Clock interface {
	Now() time.Time
}

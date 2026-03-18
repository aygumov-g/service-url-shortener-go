package get_link

import (
	"context"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type LinkRepository interface {
	GetByID(ctx context.Context, id int64) (*link_d.Link, error)
	GetByCustomCode(ctx context.Context, code string) (*link_d.Link, error)
}

type Generator interface {
	Decode(code string) (int64, error)
}

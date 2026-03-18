package get_link

import (
	"context"
	"time"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type GetLinkUsecase interface {
	Execute(ctx context.Context, code string) (*link_d.Link, error)
}

type Clock interface {
	Now() time.Time
}

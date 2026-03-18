package redirect

import (
	"context"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type GetLinkUsecase interface {
	Execute(ctx context.Context, code string) (*link_d.Link, error)
}

type UpdateLinkUsecase interface {
	Execute(ctx context.Context, id int64) error
}

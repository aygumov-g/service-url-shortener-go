package get_link

import (
	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type GetLinkUsecase interface {
	Execute(code string) (*link_d.Link, error)
}

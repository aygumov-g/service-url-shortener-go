package get_link

import (
	"context"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type GetLink struct {
	linksRepo LinkRepository
	gen       Generator
}

func NewGetLink(
	linksRepo LinkRepository,
	gen Generator,
) *GetLink {
	return &GetLink{
		linksRepo: linksRepo,
		gen:       gen,
	}
}

func (uc *GetLink) Execute(ctx context.Context, code string) (*link_d.Link, error) {
	id, err := uc.gen.Decode(code)
	if err == nil {
		link, err := uc.linksRepo.GetByID(ctx, id)
		if err == nil && link != nil {
			return link, nil
		}
	}

	return uc.linksRepo.GetByCustomCode(ctx, code)
}

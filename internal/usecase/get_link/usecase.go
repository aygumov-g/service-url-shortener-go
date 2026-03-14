package get_link

import (
	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type GetLink struct {
	linksRepo LinkRepository
	gen       Generator
	clk       Clock
}

func NewGetLink(
	linksRepo LinkRepository,
	gen Generator,
	clk Clock,
) *GetLink {
	return &GetLink{
		linksRepo: linksRepo,
		gen:       gen,
		clk:       clk,
	}
}

func (uc *GetLink) Execute(code string) (*link_d.Link, error) {
	id, err := uc.gen.Decode(code)
	if err == nil {
		link, err := uc.linksRepo.GetByID(id)
		if err == nil && link != nil {
			return link, nil
		}
	}

	return uc.linksRepo.GetByCustomCode(code)
}

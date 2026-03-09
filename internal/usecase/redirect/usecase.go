package redirect

import (
	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"
)

type Redirect struct {
	linksRepo LinkRepository
	gen       Generator
	clk       Clock
}

func NewRedirect(
	linksRepo LinkRepository,
	gen Generator,
	clk Clock,
) *Redirect {
	return &Redirect{
		linksRepo: linksRepo,
		gen:       gen,
		clk:       clk,
	}
}

func (uc *Redirect) Execute(code string) (*link_d.Link, error) {
	id, err := uc.gen.Decode(code)
	if err == nil {
		link, err := uc.linksRepo.GetByID(id)
		if err == nil && link != nil {
			return link, nil
		}
	}

	return uc.linksRepo.GetByCustomCode(code)
}

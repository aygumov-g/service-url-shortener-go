package update_link

import "context"

type UpdateLink struct {
	linksRepo LinkRepository
	clk       Clock
}

func NewUpdateLink(
	linksRepo LinkRepository,
	clk Clock,
) *UpdateLink {
	return &UpdateLink{
		linksRepo: linksRepo,
		clk:       clk,
	}
}

func (uc *UpdateLink) Execute(ctx context.Context, id int64) error {
	return uc.linksRepo.Update(ctx, id, uc.clk.Now())
}

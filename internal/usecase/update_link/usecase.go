package update_link

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

func (uc *UpdateLink) Execute(id int64) error {
	return uc.linksRepo.Update(id, uc.clk.Now())
}

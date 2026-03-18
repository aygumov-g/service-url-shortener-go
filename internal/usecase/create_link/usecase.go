package link

import (
	"context"
	"net/url"
	"strings"

	link_d "github.com/aygumov-g/service-url-shortener-go/internal/domain/link"

	"golang.org/x/net/idna"
)

type CreateLink struct {
	linksRepo LinkRepository
	gen       Generator
	clk       Clock
	domain    string
}

func NewCreateLink(
	linksRepo LinkRepository,
	gen Generator,
	clk Clock,
	domain string,
) *CreateLink {
	return &CreateLink{
		linksRepo: linksRepo,
		gen:       gen,
		clk:       clk,
		domain:    domain,
	}
}

func (uc *CreateLink) Execute(ctx context.Context, original string) (string, error) {
	if len(original) > 5000 {
		return "", link_d.ErrUrlToLong
	}

	if !strings.HasPrefix(original, "http://") && !strings.HasPrefix(original, "https://") {
		original = "https://" + original
	}

	parsed, err := url.Parse(original)
	if err != nil {
		return "", err
	}

	domainASCII, err := idna.ToASCII(parsed.Hostname())
	if err != nil {
		return "", err
	}

	ucDomainASCII, err := idna.ToASCII(uc.domain)
	if err != nil {
		return "", err
	}

	if domainASCII == ucDomainASCII {
		return "", link_d.ErrCannotShortenLink
	}

	link := &link_d.Link{
		OriginalURL: original,
		ClickCount:  0,
		CreatedAt:   uc.clk.Now(),
	}

	if err := uc.linksRepo.Create(ctx, link); err != nil {
		return "", err
	}

	return uc.gen.Encode(link.ID)
}

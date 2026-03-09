package create_link

type response struct {
	ShortURL string `json:"short_url"`
}

func (r response) toResponse(code string, domain string) *response {
	return &response{
		ShortURL: "https://" + domain + "/" + code,
	}
}

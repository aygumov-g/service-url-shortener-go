package create_link

type response struct {
	ShortURL  string `json:"short_url"`
	ShortCode string `json:"short_code"`
}

func (r response) toResponse(code string, domain string) *response {
	return &response{
		ShortURL:  "https://" + domain + "/" + code,
		ShortCode: code,
	}
}

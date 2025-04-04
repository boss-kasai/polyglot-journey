package requests

type CreateCompanyRequest struct {
	Name        string   `json:"name" binding:"required"`
	URL         []string `json:"url"`
	PhoneNumber string   `json:"phone_number"`
	PostalCode  string   `json:"postal_code"`
	Address     string   `json:"address"`
	Tags        []string `json:"tags"`
}

type SearchCompanyRequest struct {
	Name        string   `json:"name"`
	URL         []string `json:"url"`
	PhoneNumber string   `json:"phone_number"`
	PostalCode  string   `json:"postal_code"`
	Address     string   `json:"address"`
	Tags        []string `json:"tags"`
	Limit       int      `json:"limit"`
}

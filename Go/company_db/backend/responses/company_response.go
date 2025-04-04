package responses

type CreateCompanyResponse struct {
	Message string `json:"message"`
}

type CreateCompanyErrorResponse struct {
	Error string `json:"error"`
}

type CompanyData struct {
	Name        string   `json:"name"`
	URL         []string `json:"url"`
	PhoneNumber string   `json:"phone_number"`
	PostalCode  string   `json:"postal_code"`
	Address     string   `json:"address"`
	Tags        []string `json:"tags"`
}

type CompanyResponse struct {
	Num     int           `json:"num"`
	Company []CompanyData `json:"company"`
}

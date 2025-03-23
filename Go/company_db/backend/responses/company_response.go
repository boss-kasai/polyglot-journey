package responses

type CreateCompanyResponse struct {
	Message string `json:"message"`
}

type CreateCompanyErrorResponse struct {
	Error string `json:"error"`
}

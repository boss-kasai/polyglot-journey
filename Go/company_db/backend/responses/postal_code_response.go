package responses

import "company_db/backend/models"

type CreatePostalCodeResponse struct {
	Message    string            `json:"message"`
	PostalCode models.PostalCode `json:"postal_code"`
}

type CreatePostalCodeErrorResponse struct {
	Error string `json:"error"`
}

type CreatePostalCodeDuplicationResponse struct {
	Message    string            `json:"message"`
	PostalCode models.PostalCode `json:"postal_code"`
}

type SearchPostalCodeResponse struct {
	Message     string              `json:"message"`
	PostalCodes []models.PostalCode `json:"postal_codes"`
}

type SearchPostalCodeErrorResponse struct {
	Error string `json:"error"`
}

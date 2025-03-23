package requests

type CreatePostalCodeRequest struct {
	PostalCode string `json:"postal_code" binding:"required"`
	Address    string `json:"address" binding:"required"`
}

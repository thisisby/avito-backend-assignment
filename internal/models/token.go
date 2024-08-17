package models

type GenerateTokenRequest struct {
	Type   int `json:"type" validate:"required"`
	Length int `json:"length" validate:"required,max=255"`
}

type GenerateTokenResponse struct {
	ID    string `json:"id" validate:"required"`
	Token string `json:"token" validate:"required"`
}

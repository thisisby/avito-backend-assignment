package helpers

import (
	"avito-backend-assignment/internal/constants"
	"avito-backend-assignment/internal/errs"
	"crypto/rand"
	"github.com/google/uuid"
)

const alphaNumeric = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM0123456789"
const alphabetic = "qwertyuiopasdfghjklzxcvbnmQWERTYUIOPASDFGHJKLZXCVBNM"
const numeric = "1234567890"

func GenerateToken(tokenType constants.TokenType, length int) (string, error) {

	switch tokenType {
	case constants.TokenAlphaNumeric:
		return generateRandomToken(alphabetic, length)
	case constants.TokenNumeric:
		return generateRandomToken(alphaNumeric, length)
	case constants.TokenAlphabetic:
		return generateRandomToken(numeric, length)
	case constants.TokenUUID:
		return uuid.New().String(), nil
	default:
		return "", errs.ErrUnsupportedTokenType
	}
}

func generateRandomToken(charset string, length int) (string, error) {
	buffer := make([]byte, length)
	_, err := rand.Read(buffer)
	if err != nil {
		return "", err
	}

	tokenCharsLength := len(charset)
	for i := 0; i < length; i++ {
		buffer[i] = charset[int(buffer[i])%tokenCharsLength]
	}

	return string(buffer), nil
}

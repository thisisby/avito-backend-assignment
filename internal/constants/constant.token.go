package constants

type TokenType int

const (
	TokenAlphabetic TokenType = iota + 1
	TokenNumeric
	TokenAlphaNumeric
	TokenUUID
)

var (
	MapIntToToken = map[int]TokenType{
		1: TokenAlphabetic,
		2: TokenNumeric,
		3: TokenAlphaNumeric,
		4: TokenUUID,
	}
)

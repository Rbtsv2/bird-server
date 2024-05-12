package bird

type TokenType int

const (
	EOF         TokenType = iota
	IDENT                 // Identifiers
	NUMBER                // Numeric literals
	OPERATOR              // '+', '-', '*', '/', '='
	PUNCTUATION           // ',', ';'
	ILLEGAL               // Représente les tokens invalides ou inconnus
)

type Token struct {
	Type    TokenType
	Literal string
}

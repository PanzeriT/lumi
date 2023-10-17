package token

type Type string

type Token struct {
	Type    Type
	Literal string
}

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// identifiers + literals
	IDENT = "IDENT" // add, foobar, x, y, ...
	INT   = "INT"   // 1343456

	// operators
	ASSIGN = "="
	ADD    = "+"
	SUB    = "-"
	MUL    = "*"
	DIV    = "/"

	BANG = "!"

	LT = "<"
	GT = ">"

	// delimiters
	COMMA  = ","
	LPAREN = "("
	RPAREN = ")"

	// keywords
	FUNCTION = "FUNCTION"
	DO       = "DO"
	END      = "END"
	RETURN   = "RETURN"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	EQ       = "=="
	NE       = "!="
)

var keywords = map[string]Type{
	"fn":     FUNCTION,
	"do":     DO,
	"end":    END,
	"return": RETURN,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
}

// LookupIdentifier returns the token type for the given identifier
func LookupIdentifier(ident string) Type {
	tok, ok := keywords[ident]
	if ok {
		return tok
	}

	return IDENT
}

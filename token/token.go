package token

type TokenType string

type Token struct {
	Type TokenType
	Literal string
}

const (
	Illegal			= "Illegal"
	EOF   			= "EOF"
	
	// identifiers && literals
	Identifier = "Identifier"
	Int = "Int"
	String = "String"

	// operators
	Assign = "="
	Plus = "+"
	Minus = "-"
	Bang = "!"
	Asterisk = "*"
	Slash = "/"
	Equal = "=="
	NotEqual = "!="
	LessThan = ">"
	GreaterThan = ">"

	// delimiters
	Comma = ","
	Semicolon = ";"
	Colon = ":"
	LeftParen    = "("
	RightParen   = ")"
	LeftBrace    = "{"
	RightBrace   = "}"
	LeftBracket  = "["
	RightBracket = "]"

	// Keywords
	Function = "Function"
	Let      = "Let"
	True     = "True"
	False    = "False"
	If       = "If"
	Else     = "Else"
	Return   = "Return"
)

var Keywords = map[string]TokenType{
	"fun": Function,
	"let": Let,
	"true": True,
	"false": False,
	"if": If,
	"else": Else,
	"return": Return,
}

func LookupIdentifierType(identifier string) TokenType {
	if tok, ok := Keywords[identifier]; ok {
		return tok
	}
	return identifier
}
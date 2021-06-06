package ast

import (
	"bytes"

	"github.com/Sable-20/monkey-golang/token"
)

type BlockStatement struct {
	Token      token.Token
	Statements []Statement
}

func (bs *BlockStatement) expressionNode() {}

func (bs *BlockStatement) TokenLiteral() string { return bs.Token.Literal }

func (bs *BlockStatement) String() string {
	var out bytes.Buffer

	for _, s := range bs.Statements {
		out.WriteString(s.String())
	}

	return out.String()
}

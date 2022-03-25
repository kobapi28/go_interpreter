package ast

import (
	"bytes"

	"github.com/shoma3571/go_interpreter/token"
)

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 識別子を保持するため
	Value Expression  // 値を生成する式を保持するため
}

// これらが Node, Statement インターフェースを満たす
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	var out bytes.Buffer

	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")

	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}

	out.WriteString(";")
	return out.String()
}

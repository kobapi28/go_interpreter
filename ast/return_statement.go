package ast

import (
	"bytes"

	"github.com/shoma3571/go_interpreter/token"
)

type ReturnStatement struct {
	Token       token.Token // token.RETURN
	ReturnValue Expression  // 値を生成する式を保持するため
}

func (rs *ReturnStatement) statementNode() {}
func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")
	return out.String()
}

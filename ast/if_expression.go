package ast

import (
	"bytes"

	"github.com/shoma3571/go_interpreter/token"
)

type IfExpression struct {
	Token       token.Token     // if トークン
	Condition   Expression      // 条件式
	Consequence *BlockStatement // trueのとき
	Alternative *BlockStatement // falseのとき
}

func (ie *IfExpression) expressionNode() {}
func (ie *IfExpression) TokenLiteral() string {
	return ie.Token.Literal
}

func (ie *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(ie.Condition.String())
	out.WriteString(" ")
	out.WriteString(ie.Consequence.String())

	if ie.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(ie.Alternative.String())
	}

	return out.String()
}

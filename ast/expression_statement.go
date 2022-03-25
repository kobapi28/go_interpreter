package ast

import "github.com/shoma3571/go_interpreter/token"

// 式文
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression  // 式を保持する
}

func (es *ExpressionStatement) statementNode() {}
func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

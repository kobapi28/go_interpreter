package ast

import "github.com/shoma3571/go_interpreter/token"

// ASTの全てのノードはNodeインターフェースを実装する必要がある
type Node interface {
	TokenLiteral() string // そのノードが関連づけられているトークンのリテラル値を返す
}

type Statement interface {
	Node
	statementNode() // コンパイラに情報を与えるためのもの
}

type Expression interface {
	Node
	expressionNode() // コンパイラに情報を与えるためのもの
}

// 構文解析器が生成する全てのASTのルートノードになるもの
type Program struct {
	Statements []Statement
}

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name *Identifier // 識別子を保持するため
	Value Expression // 値を生成する式を保持するため
}

// これらが Node, Statement インターフェースを満たす
func (ls *LetStatement) statementNode() {}
func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}


type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}
package ast

import (
	"bytes"

	"github.com/shoma3571/go_interpreter/token"
)

// ASTの全てのノードはNodeインターフェースを実装する必要がある
type Node interface {
	TokenLiteral() string // そのノードが関連づけられているトークンのリテラル値を返す
	String() string // デバッグのためなどに用いる
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

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

// バッファを作成し、それぞれの文のString() メソッドの返り値を書き込むだけ
func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()
}

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

type Identifier struct {
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {}
func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token token.Token // token.RETURN
	ReturnValue Expression // 値を生成する式を保持するため
}

func (rs *ReturnStatement) statementNode(){}
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

// 式文
type ExpressionStatement struct {
	Token token.Token // 式の最初のトークン
	Expression Expression // 式を保持する
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

// 式なのでExpressionインターフェースを満たす
type IntegerLiteral struct {
	Token token.Token 
	Value int64
}

func (il *IntegerLiteral) expressionNode() {}
func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

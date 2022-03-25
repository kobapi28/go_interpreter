package ast

import (
	"bytes"
)

// ASTの全てのノードはNodeインターフェースを実装する必要がある
type Node interface {
	TokenLiteral() string // そのノードが関連づけられているトークンのリテラル値を返す
	String() string       // デバッグのためなどに用いる
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















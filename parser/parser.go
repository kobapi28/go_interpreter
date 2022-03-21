package parser

import (
	"github.com/shoma3571/go_interpreter/ast"
	"github.com/shoma3571/go_interpreter/lexer"
	"github.com/shoma3571/go_interpreter/token"
)

type Parser struct {
	l *lexer.Lexer // 字句解析器インスタンスへのポインタ
	curToken token.Token // 現在のトークンを指し示す
	peekToken token.Token // 次のトークンを指し示す
}

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l}
	// 2つトークンを読み込む。curToken, peekTokenの両方がセットされる
	p.nextToken()
	p.nextToken()

	return p
}

func (p *Parser) nextToken() {
	p.curToken = p.peekToken
	p.peekToken = p.l.NextToken()
}

func (p *Parser) ParseProgram() *ast.Program {
	return nil
}
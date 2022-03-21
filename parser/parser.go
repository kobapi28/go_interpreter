package parser

import (
	"fmt"

	"github.com/shoma3571/go_interpreter/ast"
	"github.com/shoma3571/go_interpreter/lexer"
	"github.com/shoma3571/go_interpreter/token"
)

type Parser struct {
	l         *lexer.Lexer // 字句解析器インスタンスへのポインタ
	curToken  token.Token  // 現在のトークンを指し示す
	peekToken token.Token  // 次のトークンを指し示す
	errors    []string     // エラー
	prefixParseFns map[token.TokenType]prefixParseFn // 前置構文解析関数
	infixParseFns map[token.TokenType]infixParseFn // 中置構文解析関数
}

type (
	prefixParseFn func() ast.Expression // 前置構文解析関数
	// 中置構文解析関数
	// ここでの引数はまた別の ast.Expression で中置演算子の左側
	infixParseFn func(ast.Expression) ast.Expression 
)

func New(l *lexer.Lexer) *Parser {
	p := &Parser{l: l, errors: []string{}}
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
	// ASTのルートノードの生成
	program := &ast.Program{}
	program.Statements = []ast.Statement{}

	// token.EOFに達するまで入力のトークンを読む
	for !p.curTokenIs(token.EOF) {
		stmt := p.parseStatement()
		if stmt != nil {
			program.Statements = append(program.Statements, stmt)
		}
		// curToken, peekToken を進める
		p.nextToken()
	}
	return program
}

// 構文解析をする
// tokenTypeによって、呼ぶ関数を振り分け
func (p *Parser) parseStatement() ast.Statement {
	switch p.curToken.Type {
	case token.LET:
		return p.parseLetStatement()
	case token.RETURN:
		return p.parseReturnStatement()
	default:
		return nil
	}
}

func (p *Parser) parseLetStatement() *ast.LetStatement {
	// 現在見ているトークン(token.LET)に基づいて、*ast.LetStatement ノードの構築
	stmt := &ast.LetStatement{Token: p.curToken}

	// 次のトークンが IDENT(変数名) を期待する。
	// token.IDENT でなければ終了する
	if !p.expectPeek(token.IDENT) {
		return nil
	}

	// Identifier ノードの作成
	stmt.Name = &ast.Identifier{Token: p.curToken, Value: p.curToken.Literal}

	// 次のトークンが ASSIGN (イコール) を期待する
	if !p.expectPeek(token.ASSIGN) {
		return nil
	}

	// TODO: セミコロンに遭遇するまで値を読み飛ばしてしまっている
	// 後で実装する
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) parseReturnStatement() *ast.ReturnStatement {
	stmt := &ast.ReturnStatement{Token: p.curToken}

	// 構文解析器を後続する式の位置へと移動させる
	p.nextToken()

	// TODO: セミコロンに遭遇するまで値を読み飛ばしてしまっている
	for !p.curTokenIs(token.SEMICOLON) {
		p.nextToken()
	}

	return stmt
}

func (p *Parser) curTokenIs(t token.TokenType) bool {
	return p.curToken.Type == t
}

func (p *Parser) peekTokenIs(t token.TokenType) bool {
	return p.peekToken.Type == t
}

// peerTokenの方をチェックし、それが正しい場合に限ってnextTokenを呼びトークンを進める
func (p *Parser) expectPeek(t token.TokenType) bool {
	if p.peekTokenIs(t) {
		p.nextToken()
		return true
	} else {
		// errorsに格納する
		p.peekError(t)
		return false
	}
}

func (p *Parser) Errors() []string {
	return p.errors
}

func (p *Parser) peekError(t token.TokenType) {
	msg := fmt.Sprintf("expected next token to be %s, got %s instead", t, p.peekToken.Type)
	p.errors = append(p.errors, msg)
}


// 前置構文解析関数のmapにエントリを追加する
func (p *Parser) registerPrefix(tokenType token.TokenType, fn prefixParseFn) {
	p.prefixParseFns[tokenType] = fn
}

// 中置構文解析関数のmapにエントリを追加する
func (p *Parser) registerInfix(tokenType token.TokenType, fn infixParseFn) {
	p.infixParseFns[tokenType] = fn
}
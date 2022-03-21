package lexer

import "github.com/shoma3571/go_interpreter/token"

type Lexer struct {
	input        string
	position     int  // 入力における現在の位置(現在の位置を指し示す)
	readPosition int  // これから読み込む位置(現在の文字の次)
	ch           byte // 現在検査中の文字
}

func New(input string) *Lexer {
	// inputのみを定義して、他はゼロ値で設定
	l := &Lexer{input: input}
	// とりあえず最初の文字を読んでおく
	l.readChar()
	return l
}

// ポインタメソッド
// 次の一文字を読んで、現在位置を進める
func (l *Lexer) readChar() {
	// 次に読み込むものがあるかないかを判定
	if l.readPosition >= len(l.input) {
		// 終端に到達した場合 0 にする
		l.ch = 0
	} else {
		l.ch = l.input[l.readPosition]
	}
	// 位置の更新
	l.position = l.readPosition
	l.readPosition += 1
}

// 現在検査中の文字 l.ch を見て、それに応じてトークンを返す
// 返す前に入力のポインタを進めて、次に読んだ時に位置が更新されるようにする
func (l *Lexer) NextToken() token.Token {
	var tok token.Token

	l.skipWhitespace()

	switch l.ch {
	case '=':
		if l.peekChar() == '=' {
			// == の場合
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.EQ, Literal: literal}
		} else {
			tok = newToken(token.ASSIGN, l.ch)
		}
	case '+':
		tok = newToken(token.PLUS, l.ch)
	case '-':
		tok = newToken(token.MINUS, l.ch)
	case '!':
		if l.peekChar() == '=' {
			// != の場合
			ch := l.ch
			l.readChar()
			literal := string(ch) + string(l.ch)
			tok = token.Token{Type: token.NOT_EQ, Literal: literal}
		} else {
			tok = newToken(token.BANG, l.ch)
		}
	case '*':
		tok = newToken(token.ASTERISK, l.ch)
	case '/':
		tok = newToken(token.SLASH, l.ch)
	case '<':
		tok = newToken(token.LT, l.ch)
	case '>':
		tok = newToken(token.GT, l.ch)
	case ';':
		tok = newToken(token.SEMICOLON, l.ch)
	case '(':
		tok = newToken(token.LPAREN, l.ch)
	case ')':
		tok = newToken(token.RPAREN, l.ch)
	case ',':
		tok = newToken(token.COMMA, l.ch)
	case '{':
		tok = newToken(token.LBRACE, l.ch)
	case '}':
		tok = newToken(token.RBRACE, l.ch)
	case 0:
		tok.Literal = ""
		tok.Type = token.EOF
	default:
		// 現在読んでいるものが英字(letter)かどうか
		if isLetter(l.ch) {
			// 英字であるなら、英字でないものが出てくるまで残りを読み進める
			tok.Literal = l.readIdentifier()
			// 返ってきた英文字列がキーワードかどうかを確認し、Typeに入れる
			tok.Type = token.LookupIdent(tok.Literal)
			// 早期の脱出が必要なのは、readIdentifierで現在の識別子の最後の文字を過ぎたところまで進んでいるから
			return tok
		} else if isDigit(l.ch) {
			tok.Type = token.INT
			tok.Literal = l.readNumber()
			return tok
		} else {
			tok = newToken(token.ILLEGAL, l.ch)
		}
	}

	l.readChar()
	return tok
}

// token初期化の役割を果たす
func newToken(tokenType token.TokenType, ch byte) token.Token {
	return token.Token{Type: tokenType, Literal: string(ch)}
}

func (l *Lexer) readIdentifier() string {
	position := l.position
	// 英字はひたすら読んで、positionを進める
	for isLetter(l.ch) {
		l.readChar()
	}
	// 読み初めから、終わりまでをスライスして返す
	return l.input[position:l.position]
}

func (l *Lexer) readNumber() string {
	position := l.position
	for isDigit(l.ch) {
		l.readChar()
	}
	return l.input[position:l.position]
}

// 英字かどうかの判定
func isLetter(ch byte) bool {
	return 'a' <= ch && ch <= 'z' || 'A' <= ch && ch <= 'Z' || ch == '_'
}

// 数字かどうかの判定
func isDigit(ch byte) bool {
	return '0' <= ch && ch <= '9'
}

// ホワイトスペースを読み飛ばす
func (l *Lexer) skipWhitespace() {
	for l.ch == ' ' || l.ch == '\t' || l.ch == '\n' || l.ch == '\r' {
		l.readChar()
	}
}

// 次の値を返す
// 次の値を覗き見したいだけなので、readChar で進めることはしない
func (l *Lexer) peekChar() byte {
	if l.readPosition >= len(l.input) {
		return 0
	} else {
		return l.input[l.readPosition]
	}
}

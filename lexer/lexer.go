package lexer

type lexer struct {
	input string
	position int // 入力における現在の位置(現在の位置を指し示す)
	readPosition int // これから読み込む位置(現在の文字の次)
	ch byte // 現在検査中の文字
}

func New(input string) *lexer {
	// inputのみを定義して、他はゼロ値で設定
	l := &lexer{input: input}
	l.readChar()
	return l
}

// ポインタメソッド
// 次の一文字を読んで、現在位置を進める
func (l *lexer) readChar() {
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
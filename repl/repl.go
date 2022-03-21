package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shoma3571/go_interpreter/lexer"
	"github.com/shoma3571/go_interpreter/token"
)

const PROMPT = ">> "

// 改行が来るまで入力ソースから読み込み、読み込んだ行を取り出して、
// 字句解析器のインスタンスに渡し、最後に全てのトークンを表示する
func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Printf(PROMPT)
		scanned := scanner.Scan()
		if !scanned {
			return
		}

		line := scanner.Text()
		l := lexer.New(line)

		for tok := l.NextToken(); tok.Type != token.EOF; tok = l.NextToken() {
			fmt.Printf("%+v\n", tok)
		}
	}
}

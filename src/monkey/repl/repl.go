package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shoebilyas123/monkeylang/monkey/lexer"
	"github.com/shoebilyas123/monkeylang/monkey/token"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	PROMPT := ">> "
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

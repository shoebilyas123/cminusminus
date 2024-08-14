package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/shoebilyas123/cminusminus/cmm/eval"
	"github.com/shoebilyas123/cminusminus/cmm/lexer"
	"github.com/shoebilyas123/cminusminus/cmm/object"
	"github.com/shoebilyas123/cminusminus/cmm/parser"
)

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	PROMPT := ">> "
	environment := object.NewEnvironment()

	for {
		fmt.Printf(PROMPT)

		scanned := scanner.Scan()

		if !scanned {
			return
		}

		line := scanner.Text()

		if line == "exit()" {
			break
		}

		l := lexer.New(line)
		p := parser.New(l)

		program := p.ParseProgram()

		if len(p.Errors()) != 0 {
			printParserErrors(out, p.Errors())
			continue
		}

		evaluated := eval.Eval(program, environment)

		if evaluated != nil {
			io.WriteString(out, evaluated.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

func printParserErrors(out io.Writer, errors []string) {
	for _, msg := range errors {
		io.WriteString(out, "\t"+msg+"\n")
	}
}

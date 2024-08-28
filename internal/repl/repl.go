package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rytsh/deg/internal/parser"
)

const PROMPT = ">> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	for {
		fmt.Fprint(out, PROMPT)
		if !scanner.Scan() {
			break
		}

		line := scanner.Text()
		if line == "exit" {
			break
		}

		l := parser.NewLexer(line)
		for token := l.NextToken(); token.Type != parser.EOF; token = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", token)
		}
	}
}

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/rytsh/deg/internal/lexer"
	"github.com/rytsh/deg/internal/token"
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

		l := lexer.New(line)
		for t := l.NextToken(); t.Type != token.EOF; t = l.NextToken() {
			fmt.Fprintf(out, "%+v\n", t)
		}
	}
}

package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/indeedhat/monkey-lang/lexer"
	"github.com/indeedhat/monkey-lang/token"
)

const PromptString = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)

	for {
		fmt.Fprintf(out, PromptString)

		if !scanner.Scan() {
			return
		}

		text := scanner.Text()
		if text == "exit" || text == "quit" {
			fmt.Fprintln(out, "Goodbye :)")
			return
		}

		lex := lexer.New(text)

		for {
			tok := lex.Next()
			if tok.Type == token.Eof {
				break
			}

			fmt.Fprintf(out, "%+v\n", tok)
		}
	}
}

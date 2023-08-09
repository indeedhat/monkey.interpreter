package repl

import (
	"bufio"
	"fmt"
	"io"

	"github.com/indeedhat/monkey-lang/evaluator"
	"github.com/indeedhat/monkey-lang/evaluator/object"
	"github.com/indeedhat/monkey-lang/lexer"
	"github.com/indeedhat/monkey-lang/parser"
)

const PromptString = "> "

func Start(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	env := object.NewEnvironment()

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
		p := parser.New(lex)

		program := p.ParseProgram()
		if len(p.Errors()) != 0 {
			for _, err := range p.Errors() {
				fmt.Fprintf(out, "parser error: %s\n", err.Error())
			}
		}

		evald := evaluator.Eval(program, env)
		if evald != nil {
			io.WriteString(out, evald.Inspect())
			io.WriteString(out, "\n")
		}
	}
}

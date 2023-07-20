package main

import (
	"fmt"
	"log"
	"os"
	"os/user"

	"github.com/indeedhat/monkey-lang/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Hello %s! Welcome to the Monkey Lang REPL!\n", user.Name)
	repl.Start(os.Stdin, os.Stdout)
}

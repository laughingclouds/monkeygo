package main

import (
	"fmt"
	"os"
	"os/user"

	"github.com/laughingclouds/monkeygo/repl"
)

func main() {
	user, err := user.Current()
	if err != nil {
		panic(err)
	}

	fmt.Printf("Welcome to the wondrous world of Monkeygo %s!\n", user.Username)
	repl.Start(os.Stdin, os.Stdout)
}

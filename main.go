package main

import (
	"fmt"
	"github.com/feralonsomaccari/my-cli-tools/commands"
	"os"
)

func main() {

	arguments := os.Args

	if len(arguments) < 2 {
		fmt.Println("Usage: ftools <tool>")
	}

	switch arguments[1] {
	case "openrep":
		fmt.Println("OPEJN THE REP")
		commands.OpenRepository()
	case "openpr":
		fmt.Println("OPEN THE PR!!")
		commands.OpenPR()
	default:
		fmt.Println("Uknown command:", arguments[1])
	}

}

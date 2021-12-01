package main

import (
	"eth_service/internal/cli"
	"os"
)

func main(){
	if !cli.Run(os.Args) {
		os.Exit(1)
	}
}

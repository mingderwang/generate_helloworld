package main

import (
	"github.com/mingderwang/generate_helloworld/gen"
	"github.com/mingderwang/generate_helloworld/parse"
	"os"
)

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		if fileName == "" {
			return
		}
		hello := parse.Scan("", fileName)
		gen.GenMain(hello)
	}
}

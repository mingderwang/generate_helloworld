/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package main

import (
	"fmt"
	"github.com/mingderwang/generate_helloworld/gen"
	"github.com/mingderwnag/generate_helloworld/parse"
	"os"
)

// don't print
type Person2 struct {
	FirstName string
	LastName  string
	HairColor string
}

// @ginger test
type Person struct {
	FirstName string
	LastName  string
	HairColor string
}

func main() {
	if len(os.Args) > 1 {
		fileName := os.Args[1]
		if fileName == "" {
			return
		}
		hello := parse.Scan("", fileName)
		fmt.Println(hello)
		gen.GenMain(hello)
	}
}

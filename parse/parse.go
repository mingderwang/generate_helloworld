/*
 * Copyright 2015 Ming-der Wang <ming@log4analytics.com> All rights reserved.
 * Licensed under MIT License.
 */

package parse

import (
	//"fmt"
	//	"github.com/davecgh/go-spew/spew"
	"go/ast"
	"go/parser"
	"go/token"
	"strings"
)

func Scan(src string, path string) string {
	fset := token.NewFileSet()
	var f *ast.File
	var err error
	if src == "" {
		f, err = parser.ParseFile(fset, path, nil, parser.ParseComments)
	} else {
		f, err = parser.ParseFile(fset, path, src, parser.ParseComments)
	}
	if err != nil {
		panic(err)
	}
	//ast.Print(fset, f)
	for _, decl := range f.Decls {
		if genDecl, ok := decl.(*ast.GenDecl); ok {
			if genDecl.Doc == nil {
			} else {
				//spew.Dump(genDecl.Doc)
				for _, comment := range genDecl.Doc.List {
					if strings.Contains(comment.Text, "@generate_helloworld") {
						for _, spec := range genDecl.Specs {
							if typeSpec, ok := spec.(*ast.ValueSpec); ok {
								//spew.Dump(genDecl.Specs)
								if typeSpec.Values != nil {
									for _, hello := range typeSpec.Values {
										if basicLit, ok := hello.(*ast.BasicLit); ok {
											//fmt.Println(basicLit.Value)
											return basicLit.Value
										}
									}
								}
							}
						}
					}
				}
			}
		}
	}
	return ""
}

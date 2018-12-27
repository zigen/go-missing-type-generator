package main

import (
	"bytes"
	"fmt"
	"github.com/zigen/go-missing-type-generator/generator"
	"os"
)

func main() {
	basePath := "./testdata/go-sample-api"
	fileName := basePath + "/generated_missing_types.go"
	g := generator.NewGenerator(basePath)
	err := g.Parse()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	g.Check()
	buf := &bytes.Buffer{}
	g.GenerateNeededTypes(buf)
	file, err := os.OpenFile(fileName, os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		fmt.Printf("Error: failed to open file\n")
		os.Exit(1)
	}
	defer file.Close()
	n, err := file.WriteString(buf.String())
	fmt.Println(n, err)
	//fmt.Printf("Generated Code: \n%s\n", buf.String())
	fmt.Println(g.Errors)
}

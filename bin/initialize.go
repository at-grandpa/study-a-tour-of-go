package main

import (
	"flag"
	"fmt"
	"os"
	"path"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("Too many arguments. expected: 1")
		os.Exit(1)
	}

	filePath := args[0]
	dirName := path.Dir(filePath)

	if err := os.MkdirAll(dirName, 0755); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	fp, err := os.OpenFile(filePath, os.O_RDWR|os.O_CREATE|os.O_TRUNC, 0755)
	defer fp.Close()

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

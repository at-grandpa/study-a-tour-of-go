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

	file, err := os.OpenFile(filePath, os.O_CREATE|os.O_WRONLY, 0755)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

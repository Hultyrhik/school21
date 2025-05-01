package main

import (
	"day01/s21_common"
	"flag"
	"fmt"
)

func main() {

	wordPtr := flag.String("f", "", "Enter a filename")
	flag.Parse()

	err := s21_common.ParseFile(*wordPtr)
	if err != nil {
		fmt.Println(err)
	}
}

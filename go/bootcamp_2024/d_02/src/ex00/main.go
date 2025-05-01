package main

import (
	w "day02/myFind/walkdirs"
	"flag"
	"log"
	"os"
)

func main() {
	targetExtFlag := flag.String("ext", "", "works ONLY when -f is specified) for user to be able to print only files with a certain extension")

	slFlag := flag.Bool("sl", false, "Prints symlinks")
	dFlag := flag.Bool("d", false, "Prints directories")
	fFlag := flag.Bool("f", false, "Prints files")

	flag.Parse()

	var input w.AllInput
	err := input.ParseAllInput(*slFlag, *dFlag, *fFlag, *targetExtFlag, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}

	err = w.PrintDir(input)
	if err != nil {
		log.Fatal(err)
	}

}

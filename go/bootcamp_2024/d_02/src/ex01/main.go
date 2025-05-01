package main

import (
	"errors"
	"flag"
	"fmt"
	"log"
	"os"
	"strings"
	"sync"
)

type Flags struct {
	Lines   bool
	M_Chars bool
	Words   bool
	Files   []string
}

func main() {
	l_flag := flag.Bool("l", false, "-l for counting lines")
	m_flag := flag.Bool("m", false, "-m for counting characters")
	w_flag := flag.Bool("w", false, "-w for counting words")
	flag.Parse()

	var input Flags
	err := input.ParseArgs(*l_flag, *m_flag, *w_flag, os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	// fmt.Println(input)

	ParseFile(input)

}

func ParseFile(f Flags) {
	var wg sync.WaitGroup
	for _, file := range f.Files {
		wg.Add(1)
		go func() {
			defer wg.Done()

			text, err := os.ReadFile(file)
			if err != nil {
				fmt.Println(err)
				return
			}

			if f.Lines {
				lines := strings.Split(string(text), "\n")
				fmt.Printf("%d\t%s\n", len(lines), file)
			}

			if f.Words {
				words := strings.Split(string(text), " ")
				var count_w int
				for _, word := range words {
					if len(word) == 0 {
						continue
					} else {
						count_w++
					}
				}
				fmt.Printf("%d\t%s\n", count_w, file)
			}

			if f.M_Chars {
				chars := strings.Split(string(text), "")
				fmt.Printf("%d\t%s\n", len(chars), file)
			}
		}()
	}
	wg.Wait()
}

func ContainsString(file string, flags []string) bool {
	for _, flag := range flags {
		if "-"+flag == file {
			return true
		}
	}
	return false
}

func (f *Flags) ParseArgs(l, m, w bool, args []string) error {
	var count int
	if l {
		f.Lines = true
		count++
	}
	if m {
		f.M_Chars = true
		count++
	}
	if w {
		f.Words = true
		count++
	}

	if count > 1 {
		return errors.New("too many options specified")
	}

	if count == 0 {
		f.Words = true
	}

	// fmt.Println("len(args):", len(args))
	// fmt.Println("args:", args)

	flags := []string{"l", "m", "w"}
	for _, file := range args {
		if ContainsString(file, flags) {
			continue
		} else {
			f.Files = append(f.Files, file)
		}
	}

	if len(f.Files) == 0 {
		return errors.New("file is not specified")
	}

	return nil
}

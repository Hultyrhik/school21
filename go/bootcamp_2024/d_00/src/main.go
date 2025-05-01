package main

import (
	"flag"
	"fmt"
	"os"
	"sort"
)

func main() {
	optionsPtr := flag.Int("mode", 4, "Select options from 1 to 4 to to print one or up to 4 elements of Anscombe's quartet")
	flag.Parse()

	chosenOptions := mode_option{}
	modeError := parseModeOptions(&chosenOptions, *optionsPtr)
	if modeError != nil {
		fmt.Println(modeError)
	}

	numbers := parse_floats()

	if len(numbers) == 0 {
		fmt.Println("Zero valid numbers are entered")
		os.Exit(1)
	}

	sort.Float64s(numbers)

	if chosenOptions.is_mean {
		fmt.Println("Mean:", mean(numbers))
	}
	if chosenOptions.is_median {
		fmt.Println("Median:", median(numbers))
	}
	if chosenOptions.is_mode {
		modeResult, err := mode(numbers)
		if err == nil {
			fmt.Println("Mode:", modeResult)
		} else {
			fmt.Println(err)
		}

	}
	if chosenOptions.is_standardDeviation {
		fmt.Println("Standard deviation:", standardDeviation(numbers))
	}

}

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

type mode_option struct {
	is_mean              bool
	is_median            bool
	is_mode              bool
	is_standardDeviation bool
}

func roundFloat(num float64, precision uint) float64 {
	ratio := math.Pow10(int(precision))
	return math.Round(num*ratio) / ratio
}

func parse_floats() []float64 {
	fmt.Println("Please enter numbers between -100'000 and 100'000, separated by newlines")
	fmt.Println("To exit the loop please press 'q'")

	floatsSlice := make([]float64, 0)

	for {
		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		inputText := scanner.Text()

		if strings.ToLower(inputText) == "q" {
			break
		}

		num, err := strconv.ParseFloat(inputText, 64)
		if err == nil {
			num = roundFloat(num, 2)
			floatsSlice = append(floatsSlice, num)
		} else {
			fmt.Println(err)
		}
	}
	return floatsSlice
}

func parseModeOptions(opts *mode_option, mode int) (err error) {
	switch mode {
	default:
		err = fmt.Errorf("mode options is out of bound: %d is not a valid options. Using default option", mode)
		fallthrough
	case 4:
		opts.is_standardDeviation = true
		fallthrough
	case 3:
		opts.is_mode = true
		fallthrough
	case 2:
		opts.is_median = true
		fallthrough
	case 1:
		opts.is_mean = true
	}

	return err
}

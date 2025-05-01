package main

import (
	"fmt"
	"math"
)

func mean(nums []float64) float64 {

	var total float64

	for _, num := range nums {
		total += num
	}

	return roundFloat(total/float64(len(nums)), 2)
}

func median(nums []float64) float64 {
	var median_result float64
	if length := len(nums); length%2 == 0 {
		both_nums := nums[length/2] + nums[length/2-1]
		median_result = both_nums / 2
	} else {
		median_result = nums[length/2]
	}

	return roundFloat(median_result, 2)
}

func mode(nums []float64) (float64, error) {
	counter := make(map[float64]uint)

	for _, num := range nums {
		counter[num] += 1
	}

	var mostFrequentKey float64 = nums[0]
	var mostFrequentValue uint
	for key, value := range counter {
		if value > mostFrequentValue {
			mostFrequentValue = value
			mostFrequentKey = key
		} else if value == mostFrequentValue {
			if key < mostFrequentKey {
				mostFrequentKey = key
			}
		}
	}

	if mostFrequentValue == 1 {
		return 0.0, fmt.Errorf("there is no mode - all numbers are encountered once")
	}

	return roundFloat(mostFrequentKey, 2), nil
}

func standardDeviation(nums []float64) float64 {
	var total float64
	meanValue := mean(nums)
	length := len(nums)
	for _, num := range nums {
		difference := num - meanValue
		total += (difference * difference)
	}
	variance := total / float64(length)
	return roundFloat(math.Sqrt(variance), 2)
}

package main

import (
	"testing"
)

func TestAll_1(t *testing.T) {
	sampleSlice := []float64{1, 2, 3, 4}
	answer_mean := 2.5
	answer_median := 2.5
	answer_mode := 0.0
	answer_population_DS := 1.12

	mean_result := mean(sampleSlice)
	if mean_result != answer_mean {
		t.Errorf("Result for mean was incorrect, got: %f, want: %f.", mean_result, answer_mean)
	}

	median_result := median(sampleSlice)
	if median_result != answer_median {
		t.Errorf("Result for median was incorrect, got: %f, want: %f.", median_result, answer_median)
	}

	mode_result, err := mode(sampleSlice)
	if err != nil {
		// t.Errorf("No mode found")
	} else if mode_result != answer_mode {
		t.Errorf("Result for mode was incorrect, got: %f, want: %f.", median_result, answer_median)
	}

	population_DS_result := standardDeviation(sampleSlice)
	if population_DS_result != answer_population_DS {
		t.Errorf("Result for Standard Deviation was incorrect, got: %f, want: %f.", population_DS_result, answer_population_DS)
	}
}

func TestAll_2(t *testing.T) {
	sampleSlice := []float64{-3, -3, -1, 0, 1, 22, 543}
	answer_mean := 79.86
	answer_median := 0.0
	answer_mode := -3.0
	answer_population_DS := 189.25

	mean_result := mean(sampleSlice)
	if mean_result != answer_mean {
		t.Errorf("Result for mean was incorrect, got: %f, want: %f.", mean_result, answer_mean)
	}

	median_result := median(sampleSlice)
	if median_result != answer_median {
		t.Errorf("Result for median was incorrect, got: %f, want: %f.", median_result, answer_median)
	}

	mode_result, err := mode(sampleSlice)
	if err != nil {
		// t.Errorf("No mode found")
	} else if mode_result != answer_mode {
		t.Errorf("Result for mode was incorrect, got: %f, want: %f.", median_result, answer_median)
	}

	population_DS_result := standardDeviation(sampleSlice)
	if population_DS_result != answer_population_DS {
		t.Errorf("Result for Standard Deviation was incorrect, got: %f, want: %f.", population_DS_result, answer_population_DS)
	}
}

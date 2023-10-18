package main

import (
	"testing"
)

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func TestEvenSegregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6}, []string{"even"}, All)
	expectedResult := []int{2, 4, 6}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestOddSegregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6}, []string{"odd"}, All)
	expectedResult := []int{1, 3, 5}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestPrimeSegregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6}, []string{"prime"}, All)
	expectedResult := []int{1, 2, 3, 5}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestOddPrimeSegregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6}, []string{"prime", "odd"}, All)
	expectedResult := []int{1, 3, 5}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestEvenMult5Segregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6, 10, 15, 20}, []string{"even", "multiple of 5"}, All)
	expectedResult := []int{10, 20}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestOddMult3G10Segregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6, 9, 10, 15, 21, 25, 27}, []string{"odd", "multiple of 3", "greater than 10"}, All)
	expectedResult := []int{15, 21, 27}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

func TestOddOrMult3OrG10Segregator(t *testing.T) {
	result := NumberFinder([]int{1, 2, 3, 4, 5, 6, 9, 10, 15, 21, 25, 27}, []string{"multiple of 3", "greater than 10"}, Any)
	expectedResult := []int{3, 6, 9, 15, 21, 27, 25}
	if !Equal(result, expectedResult) {
		t.Error("Result was incorrect")
	}
}

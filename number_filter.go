package main

import (
	"math"
	"strconv"
	"strings"

	"golang.org/x/exp/slices"
)

type NumberType = int
type ConditionType = int

const (
	Even NumberType = iota
	Odd
	Prime
	OddPrime
	EvenMult5
	OddMult3G10
)

const (
	Any ConditionType = iota
	All
)

func NumberFinder(numbers []int, conditions []string, conditionType ConditionType) []int {
	if conditionType == All {
		for _, condition := range conditions {
			if len(numbers) == 0 {
				break
			}
			numbers = OperationEvaluator(numbers, condition)
		}
		return numbers
	} else {
		resultNumbers := make([]int, 0)
		for _, condition := range conditions {
			if len(numbers) == 0 {
				break
			}
			result := OperationEvaluator(numbers, condition)
			resultNumbers = append(resultNumbers, result...)
			numbers = remove(numbers, result)
		}
		return resultNumbers
	}
}

func remove(mainSlice []int, sliceToRemove []int) []int {
	result := make([]int, 0)
	for _, num := range mainSlice {
		if !slices.Contains(sliceToRemove, num) {
			result = append(result, num)
		}
	}
	return result
}

func OperationEvaluator(numbers []int, condition string) []int {

	baseCondition := ""
	numberStr := ""
	for _, c := range strings.Split(condition, "") {
		if c == " " {
			continue
		} else if _, err := strconv.Atoi(c); err != nil {
			baseCondition = baseCondition + c
		} else {
			numberStr = numberStr + c
		}
	}

	number, _ := strconv.Atoi(numberStr)

	switch strings.ToLower(baseCondition) {
	case "even":
		return divisibility(numbers, 2, 0)
	case "odd":
		return divisibility(numbers, 2, 1)
	case "prime":
		return prime(numbers)
	case "greaterthan":
		return greaterThan(numbers, number)
	case "lessthan":
		return lessThan(numbers, number)
	case "multipleof":
		return divisibility(numbers, number, 0)
	default:
		return numbers
	}
}

func NumberSegregator(numbers []int, numType NumberType) []int {
	switch numType {
	case Even:
		return divisibility(numbers, 2, 0)
	case Odd:
		return divisibility(numbers, 2, 1)
	case Prime:
		return prime(numbers)
	case OddPrime:
		return prime(NumberSegregator(numbers, Odd))
	case EvenMult5:
		return divisibility(numbers, 10, 0)
	case OddMult3G10:
		return greaterThan(divisibility(divisibility(numbers, 3, 0), 2, 1), 10)
	default:
		return divisibility(numbers, 2, 0)
	}
}

func greaterThan(numbers []int, compareWith int) []int {
	resultNumbers := make([]int, 0)
	for _, num := range numbers {
		if num > compareWith {
			resultNumbers = append(resultNumbers, num)
		}
	}
	return resultNumbers
}

func lessThan(numbers []int, compareWith int) []int {
	resultNumbers := make([]int, 0)
	for _, num := range numbers {
		if num < compareWith {
			resultNumbers = append(resultNumbers, num)
		}
	}
	return resultNumbers
}

func divisibility(numbers []int, divisor int, remainder int) []int {
	resultNumbers := make([]int, 0)
	for _, num := range numbers {
		if num%divisor == remainder {
			resultNumbers = append(resultNumbers, num)
		}
	}
	return resultNumbers
}

func prime(numbers []int) []int {
	resultNumbers := make([]int, 0)
	for _, num := range numbers {

		var isPrime bool = true
		for i := 2; i <= int(math.Sqrt(float64(num))); i++ {
			if num%i == 0 {
				isPrime = false
				break
			}
		}

		if isPrime {
			resultNumbers = append(resultNumbers, num)
		}
	}
	return resultNumbers
}

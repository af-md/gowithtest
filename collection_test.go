package main

import (
	"slices"
	"testing"
)

func TestSum(t *testing.T) {
	t.Run("sum fixed size array int", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}

		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

	t.Run("sum fixed size array int", func(t *testing.T) {
		numbers := []int{1, 2, 3}

		got := Sum(numbers)
		want := 6

		if got != want {
			t.Errorf("got %d want %d given, %v", got, want, numbers)
		}
	})

}

func Sum(numbers []int) int {
	var sum int
	for _, v := range numbers {
		sum += v
	}
	return sum
}

func TestSumAllTail(t *testing.T) {
	t.Run("sum the tail of 2 arrays together", func(t *testing.T) {
		firstArray := []int{1, 2}
		secondArray := []int{0, 9}

		got := SumAllTail(firstArray, secondArray)
		want := []int{2, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d given firstArray: %d   secondArray: %d", got, want, firstArray, secondArray)
		}
	})

	t.Run("sum the tail of empty array", func(t *testing.T) {
		firstArray := []int{}
		secondArray := []int{2, 4, 5}

		got := SumAllTail(firstArray, secondArray)
		want := []int{0, 9}

		if !slices.Equal(got, want) {
			t.Errorf("got %d want %d given firstArray: %d   secondArray: %d", got, want, firstArray, secondArray)
		}
	})
}

func SumAllTail(numbersToSum ...[]int) []int {
	var sums []int
	for _, numbers := range numbersToSum {
		if len(numbers) == 0 {
			sums = append(sums, 0)
		} else {
			tail := numbers[1:]
			sums = append(sums, Sum(tail))
		} 
	}
	return sums
}
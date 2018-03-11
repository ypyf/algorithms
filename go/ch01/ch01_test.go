package ch01

import (
	"reflect"
	"testing"
)

func TestNextSubset(t *testing.T) {
	expected := "0110"
	r := NextSubset("0101")
	if r != expected {
		t.Errorf("NextSubset was incorrect: got %s, expected %s\n", r, expected)
	}

	subset := "01010"
	for subset != "11111" {
		t.Log(subset)
		subset = NextSubset(subset)
	}
	t.Log(subset)
}

func TestSelectionSort(t *testing.T) {
	expected := []int{-24, 0, 3, 4, 6, 17, 24, 24, 321, 325}
	r := SelectionSort([]int{3, 321, 6, -24, 325, 24, 24, 17, 4, 0})
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("SelectionSort was incorrect: got %v, expected %v\n", r, expected)
	}
}

func TestBubbleSort(t *testing.T) {
	expected := []int{-24, 0, 3, 4, 6, 17, 24, 24, 321, 325}
	r := BubbleSort([]int{3, 321, 6, -24, 325, 24, 24, 17, 4, 0})
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("BubbleSort was incorrect: got %v, expectd %v\n", r, expected)
	}
}

func TestFibonacci(t *testing.T) {
	expected := []int{1, 1, 2, 3, 5, 8, 13, 21, 34, 55}
	r := []int{}
	for i := uint(0); i < 10; i++ {
		r = append(r, Fibonacci(i))
	}
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Fibonacci was incorrect: got %v, expectd %v\n", r, expected)
	}
}

func TestSieve(t *testing.T) {
	expected := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29}
	r := sieve(30)
	if !reflect.DeepEqual(r, expected) {
		t.Errorf("Sieve was incorrect: got %v, expectd %v\n", r, expected)
	}
}

func TestExtendedGcd(t *testing.T) {
	expected := struct {
		x int
		y int
		g int
	}{}
	expected.x = 7
	expected.y = -3
	expected.g = 1
	x, y, g := ExtendedGcd(10, 23)
	if x != expected.x || y != expected.y || g != expected.g {
		t.Errorf("ExtendedGcd was incorrect: got (%d, %d, %d), expectd (%d, %d, %d)\n", x, y, g, expected.x, expected.y, expected.g)
	}
}

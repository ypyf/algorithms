package ch01

import (
	"reflect"
	"testing"
)

func TestNextSubset(t *testing.T) {
	r := NextSubset("0101")
	if r != "0110" {
		t.Errorf("NextSubset was incorrect: got %s, expected %s\n", r, "0110")
	}

	subset := "01010"
	for subset != "11111" {
		t.Log(subset)
		subset = NextSubset(subset)
	}
	t.Log(subset)
}

func TestBubbleSort(t *testing.T) {
	arr := []int{3, 321, 6, -24, 325, 24, 24, 17, 4, 0}
	sorted := []int{-24, 0, 3, 4, 6, 17, 24, 24, 321, 325}
	r := BubbleSort(arr)
	if !reflect.DeepEqual(r, sorted) {
		t.Errorf("BubbleSort was incorrect: got %v, expectd %v\n", r, sorted)
	}
}

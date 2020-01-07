package array

import (
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {
	numbers := []int{1, 2, 3, 4, 5}

	got := Sum(numbers)
	expected := 15

	if got != expected {
		t.Errorf("got %v, expected %v, %v", got, expected, numbers)
	}
}

func TestSumAll(t *testing.T) {
	got := SumAll([]int{1, 2, 3, 4, 5}, []int{6, 7})

	want := []int{15, 13}

	// reflect.DeepEqual存在怪异行为 例如 字符串和 slice 竟然可以一起比较
	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, expected %v", got, want)
	}
}

func TestSumTail(t *testing.T) {

	checkSum := func(t *testing.T, got, want []int) {
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, expected %v", got, want)
		}
	}

	t.Run("make the sum of tails of", func(t *testing.T) {
		got := SumTail([]int{2, 3}, []int{1, 9, 3})
		want := []int{3, 12}
		checkSum(t, got, want)
	})

	t.Run("Safely empty slice", func(t *testing.T) {
		got := SumTail([]int{}, []int{1})
		want := []int{0, 0}
		checkSum(t, got, want)
	})
}

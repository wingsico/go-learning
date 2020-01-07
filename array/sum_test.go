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
	got := SumTail([]int{2, 3}, []int{1, 9, 3})
	want := []int{3, 12}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %v, expected %v", got, want)
	}

}

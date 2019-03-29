package helper

import (
	"testing"
)

func TestMax_1(t *testing.T) {
	v := []int{1, 2, 3}
	expected := 3

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_2(t *testing.T) {
	v := []int{1, 2, 4}
	expected := 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

func TestMax_3(t *testing.T) {
	v := []int32{1, 2, 4}
	var expected int32 = 4

	result := Max(v)

	if result != expected {
		t.Error("Value should be expected!")
	}
}

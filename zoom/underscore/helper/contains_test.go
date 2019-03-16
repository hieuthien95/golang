package helper

import "testing"

func TestContains2_1(t *testing.T) {
	list := []int{1, 2, 3}
	value := 3
	expected := true
	actual := Contains2(list, value)
	if actual != expected {
		t.Error("actual should be same exptected")
	}
}

func TestContains2_2(t *testing.T) {
	list := "Hello world"
	value := "world"
	expected := true
	actual := Contains2(list, value)
	if actual != expected {
		t.Error("actual should be same exptected")
	}
}

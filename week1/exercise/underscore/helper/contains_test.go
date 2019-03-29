package helper

import "testing"

func TestContains_1(t *testing.T) {
	list := []int{1, 2, 3}
	value := 3
	expected := true
	actual := Contains(list, value)
	if actual != expected {
		t.Error("actual should be same exptected")
	}
}

func TestContains_2(t *testing.T) {
	list := "Hello world"
	value := "world"
	expected := true
	actual := Contains(list, value)
	if actual != expected {
		t.Error("actual should be same exptected")
	}
}

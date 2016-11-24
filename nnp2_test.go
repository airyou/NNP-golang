package main

import "testing"

func TestGetLastButOne(t *testing.T) {
	list := []int{0, 1, 2, 3, 4, 5}
	ans := 4
	result, err := GetLastButOne(list)
	if err != nil {
		t.Errorf("Error1")
	}
	if result != ans {
		t.Errorf("Error2")
	}
}

func TestGetLastButOneError1(t *testing.T) {
	list := []int{0}
	ans := 0
	result, err := GetLastButOne(list)
	if err == nil {
		t.Errorf("Error3")
	}
	if err.Error() != "EmptyError" {
		t.Errorf("Error4")
	}
	if result != ans {
		t.Errorf("Error5")
	}
}

func TestGetLastButOneError2(t *testing.T) {
	ans := 0
	result, err := GetLastButOne(nil)
	if err == nil {
		t.Errorf("Error6")
	}
	if err.Error() != "nilError" {
		t.Errorf("Error7")
	}
	if result != ans {
		t.Errorf("Error8")
	}
}

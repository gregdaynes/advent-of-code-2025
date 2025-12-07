package main

import "testing"

var input = "123 328  51 64 \n 45 64  387 23 \n  6 98  215 314\n*   +   *   +  "
var want = 3263827

func TestDay06(t *testing.T) {
	got := Day06(input)
	if got != want {
		t.Errorf("wanted %v, got %v", want, got)
	}
}

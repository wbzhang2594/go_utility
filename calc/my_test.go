package calc_test

import (
	"testing"

	"github.com/wbzhang2594/go_utility/calc"
)

func TestAdd(t *testing.T) {
	var a1 int = 2
	var b1 int = 3

	if calc.Add(a1, b1) != 5 {
		t.Fatal("wrong :(")
	}
}

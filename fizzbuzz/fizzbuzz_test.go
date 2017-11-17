package fizzbuzz

import "testing"

func TestInputOneSayOne(t *testing.T) {
	fizzbuzz := Say(1)

	if fizzbuzz != "1" {
		t.Error("result != 1")
	}
}

func TestInputTwoSayTwo(t *testing.T) {
	fizzbuzz := Say(2)

	if fizzbuzz != "2" {
		t.Error("result != 2")
	}
}



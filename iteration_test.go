package main

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("expected %q but got %q", expected, repeated)
	}
}

const repeatCount = 5

func Repeat(c string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += c
	}
	return repeated
}

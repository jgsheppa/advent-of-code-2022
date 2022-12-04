package dayFour

import "testing"

func TestCalculateEntireOverlappingSections(t *testing.T) {
	got, err := CalculateEntireOverlappingSections("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateEntireOverlappingSections", err)
	}
	want := 2

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestCalculateOverlappingSections(t *testing.T) {
	got, err := CalculateOverlappingSections("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateOverlappingSections", err)
	}
	want := 4

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

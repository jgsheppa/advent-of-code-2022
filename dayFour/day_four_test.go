package dayFour

import "testing"

func TestCalculateOverlappingSections(t *testing.T) {
	got, err := CalculateOverlappingSections("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateOverlappingSections", err)
	}
	want := 3

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

package dayTwo

import "testing"

func TestCalculateRockPaperScissorsScore(t *testing.T) {
	got, err := CalculateRockPaperScissorsScore("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateRockPaperScissorsScore", err)
	}
	want := 15

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

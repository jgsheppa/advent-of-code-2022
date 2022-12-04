package dayTwo

import (
	"fmt"
	"testing"
)

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

func TestCalculateFinalRockPaperScissorsScore(t *testing.T) {
	got, err := CalculateFinalRockPaperScissorsScore("test_input_2.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateRockPaperScissorsScore", err)
	}
	want := 12

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func BenchmarkCalculateFinalRockPaperScissorsScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateFinalRockPaperScissorsScore("test_input_2.txt")
	}
}

func BenchmarkCalculateRockPaperScissorsScore(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateRockPaperScissorsScore("test_input.txt")
	}
}

func ExampleCalculateFinalRockPaperScissorsScore() {
	got, err := CalculateFinalRockPaperScissorsScore("test_input_2.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateRockPaperScissorsScore", err)
	}
	fmt.Println(got)
	// Output: 12
}

func ExampleCalculateRockPaperScissorsScore() {
	got, err := CalculateRockPaperScissorsScore("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateRockPaperScissorsScore", err)
	}
	fmt.Println(got)
	// Output: 15
}

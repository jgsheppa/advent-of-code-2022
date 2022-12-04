package dayThree

import (
	"fmt"
	"testing"
)

func TestCalculateSumOfDuplicateItemPriority(t *testing.T) {
	got, err := CalculateSumOfDuplicateItemPriority("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateSumOfDuplicateItemPriority", err)
	}
	want := 157

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestCalculateSumOfDuplicateElfBadges(t *testing.T) {
	got, err := CalculateSumOfDuplicateElfBadges("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateSumOfDuplicateElfBadges", err)
	}
	want := 70

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func BenchmarkCalculateSumOfDuplicateItemPriority(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSumOfDuplicateItemPriority("test_input.txt")
	}
}

func BenchmarkCalculateSumOfDuplicateElfBadges(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSumOfDuplicateElfBadges("test_input.txt")
	}
}

func ExampleCalculateSumOfDuplicateElfBadges() {
	got, err := CalculateSumOfDuplicateElfBadges("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateSumOfDuplicateElfBadges", err)
	}
	fmt.Println(got)
	// Output: 70
}

func ExampleCalculateSumOfDuplicateItemPriority() {
	got, err := CalculateSumOfDuplicateItemPriority("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateSumOfDuplicateItemPriority", err)
	}
	fmt.Println(got)
	// Output: 157
}

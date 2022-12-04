package dayOne

import (
	"fmt"
	"testing"
)

func TestFindMaxElfCalories(t *testing.T) {
	got, err := FindMaxElfCalories("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	want := 5000

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestFindSumTopThreeElfCalories(t *testing.T) {
	got, err := FindSumTopThreeElfTotalCalories("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	want := 10000

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func BenchmarkFindMaxElfCalories(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxElfCalories("test_input.txt")
	}
}

func BenchmarkFindSumTopThreeElfCalories(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindSumTopThreeElfTotalCalories("test_input.txt")
	}
}

func ExampleFindMaxElfCalories() {
	got, err := FindMaxElfCalories("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	fmt.Println(got)
	// Output: 10000
}

func ExampleFindSumTopThreeElfTotalCalories() {
	got, err := FindSumTopThreeElfTotalCalories("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	fmt.Println(got)
	// Output: 5000
}

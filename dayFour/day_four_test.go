package dayFour

import (
	"fmt"
	"testing"
)

func TestCalculateEntireOverlappingSections(t *testing.T) {
	got, err := CalculateSectionsContainingOtherSections("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateSectionsContainingOtherSections", err)
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

func BenchmarkCalculateEntireOverlappingSections(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateSectionsContainingOtherSections("test_input.txt")
	}
}

func BenchmarkCalculateOverlappingSections(b *testing.B) {
	for i := 0; i < b.N; i++ {
		CalculateOverlappingSections("test_input.txt")
	}
}

func ExampleCalculateOverlappingSections() {
	got, err := CalculateOverlappingSections("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateOverlappingSections", err)
	}
	fmt.Println(got)
	// Output: 4
}

func ExampleCalculateSectionsContainingOtherSections() {
	got, err := CalculateSectionsContainingOtherSections("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "CalculateSectionsContainingOtherSections", err)
	}
	fmt.Println(got)
	// Output: 2
}

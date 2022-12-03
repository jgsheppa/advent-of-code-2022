package dayThree

import "testing"

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

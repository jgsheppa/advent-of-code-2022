package dayFive

import "testing"

func TestFindTopCrates(t *testing.T) {
	got, err := FindTopCrates("test_input.txt", "test_map.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "CalculateSumOfDuplicateElfBadges", err)
	}
	want := "CMZ"

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

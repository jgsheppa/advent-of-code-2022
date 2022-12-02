package dayOne

import "testing"

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

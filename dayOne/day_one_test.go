package dayOne

import "testing"

func TestFindMaxElfCalories(t *testing.T) {
	got, err := FindMaxElfCalories("test_input_1.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	want := 5000

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

package daySeven

import "testing"

func TestFindDirectoryWithMostMemory(t *testing.T) {
	got, err := FindDirectoryWithMostMemory("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "FindDirectoryWithMostMemory", err)
	}
	want := 95437

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

func TestDeleteDirectory(t *testing.T) {
	got, err := DeleteDirectory("test_input.txt")
	if err != nil {
		t.Errorf("test: %v error: %v", "FindDirectoryWithMostMemory", err)
	}
	want := 24933642

	if got != want {
		t.Errorf("got = %v want = %v", got, want)
	}
}

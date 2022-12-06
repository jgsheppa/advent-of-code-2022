package daySix

import "testing"

func TestFindStartOfPacket(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one", "bvwbjplbgvbhsrlpgdmjqwftvncz", 5},
		{"two", "nppdvjthqldpwncqszvftbrmjlhg", 6},
		{"three", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 10},
		{"four", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 11},
		{"five", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 7},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FindStartOfPacket(tc.input)
			if err != nil {
				t.Errorf("got following error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %v -- want %v", got, tc.want)
			}
		})
	}
}

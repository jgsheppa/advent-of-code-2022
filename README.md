# Advent of Code 2022 using Golang & Test Driven Development

## Concept
For the past year I have been learning the Go, and I took a relatively scattered approach. I have 
done coding challenges, built a website, tinkered with a search engine, toyed around with CLI apps, 
and had stretches where I didn't write Go at all. There are so many directions, projects, and packages 
to learn about, but with the help of the great Go community and its teachers, I was able to learn 
how to think in Go.

One year is as good of a timespan as any to test one's knowledge, and Advent of Code seemed 
like the perfect environment to put some of my Go skills to the test. One commonality between Go
Advent of Code is both lend themselves well to Test Driven Development (TDD).

Here is an example of a test, an example test, and a benchmark from day 
one's challenge.
```Go
package main

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

func BenchmarkFindMaxElfCalories(b *testing.B) {
	for i := 0; i < b.N; i++ {
		FindMaxElfCalories("test_input.txt")
	}
}

func ExampleFindMaxElfCalories() {
	got, err := FindMaxElfCalories("test_input.txt")
	if err != nil {
		fmt.Errorf("test: %v error: %v", "TestFindMaxElfCalories", err)
	}
	fmt.Println(got)
	// Output: 5000
}
```
Here is an example of a table driven test from day six's challenge.
```Go
package main

import "testing"

func TestFindStartOfMessage(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"one", "mjqjpqmgbljsphdztnvjfqwrcgsmlb", 19},
		{"two", "bvwbjplbgvbhsrlpgdmjqwftvncz", 23},
		{"three", "nppdvjthqldpwncqszvftbrmjlhg", 23},
		{"four", "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", 29},
		{"five", "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw", 26},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			got, err := FindStartOfMessageOrSignal(tc.input, 14)
			if err != nil {
				t.Errorf("got following error: %v", err)
			}
			if got != tc.want {
				t.Errorf("got %v -- want %v", got, tc.want)
			}
		})
	}
}
```
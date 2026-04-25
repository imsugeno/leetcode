package main

import (
	"testing"
)

func TestNumIdenticalPairs(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"example1", []int{1, 2, 3, 1, 1, 3}, 4},
		{"example2", []int{1, 1, 1, 1}, 6},
		{"example3", []int{1, 2, 3}, 0},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := numIdenticalPairs(tt.nums)
			if got != tt.want {
				t.Fatalf("got %v, want %v", got, tt.want)
			}
		})
	}
}

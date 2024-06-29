package parser

import (
	"reflect"
	"testing"
)

// TestGenerateIntervals tests the generateIntervals method
func TestGenerateIntervals(t *testing.T) {
	parser := &parser{}

	tests := []struct {
		from, to, interval int
		expected           []int
	}{
		{0, 59, 15, []int{0, 15, 30, 45}},
		{0, 59, 10, []int{0, 10, 20, 30, 40, 50}},
		{1, 5, 1, []int{1, 2, 3, 4, 5}},
	}

	for _, tt := range tests {
		t.Run("", func(t *testing.T) {
			result := parser.generateIntervals(tt.from, tt.to, tt.interval)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}

// TestParse tests the Parse method
func TestParse(t *testing.T) {
	parser := &parser{floor:0, ceil: 59}

	tests := []struct {
		expression string
		expected   []int
		expectErr  bool
	}{
		{"*/15", []int{0, 15, 30, 45}, false},
		{"*/10", []int{0, 10, 20, 30, 40, 50}, false},
		{"*", []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27, 28, 29, 30, 31, 32, 33, 34, 35, 36, 37, 38, 39, 40, 41, 42, 43, 44, 45, 46, 47, 48, 49, 50, 51, 52, 53, 54, 55, 56, 57, 58, 59}, false},
		{"1,2,3", []int{1, 2, 3}, false},
		{"1-5", []int{1, 2, 3, 4, 5}, false},
		{"60", nil, true},
		{"*/invalid", nil, true},
		{"1-60", nil, true},
		{"1,invalid", nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.expression, func(t *testing.T) {
			result, err := parser.parse(tt.expression)
			if (err != nil) != tt.expectErr {
				t.Errorf("expected error: %v, got: %v", tt.expectErr, err)
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("expected: %v, got: %v", tt.expected, result)
			}
		})
	}
}


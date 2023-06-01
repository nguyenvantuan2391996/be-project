package helper

import (
	"math"
	"testing"
)

func TestFindMax(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty_slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "all_positive",
			input:    []int{1, 2, 3, 4, 5},
			expected: 5,
		},
		{
			name:     "all_negative",
			input:    []int{-5, -4, -3, -2, -1},
			expected: -1,
		},
		{
			name:     "mixed",
			input:    []int{-10, 5, 3, -20, 7, 9},
			expected: 9,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindMax(tc.input)
			if actual != tc.expected {
				t.Errorf("Test case %q failed: expected %v, but got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

func TestFindMin(t *testing.T) {
	testCases := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "empty_slice",
			input:    []int{},
			expected: 0,
		},
		{
			name:     "all_positive",
			input:    []int{1, 2, 3, 4, 5},
			expected: 1,
		},
		{
			name:     "all_negative",
			input:    []int{-5, -4, -3, -2, -1},
			expected: -5,
		},
		{
			name:     "mixed",
			input:    []int{-10, 5, 3, -20, 7, 9},
			expected: -20,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := FindMin(tc.input)
			if actual != tc.expected {
				t.Errorf("Test case %q failed: expected %v, but got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

func TestCalculateEuclideanDistance(t *testing.T) {
	testCases := []struct {
		name     string
		pointA   []float64
		pointB   []float64
		expected float64
	}{
		{
			name:     "same_point",
			pointA:   []float64{1.0, 2.0, 3.0},
			pointB:   []float64{1.0, 2.0, 3.0},
			expected: 0.0,
		},
		{
			name:     "different_dimensions",
			pointA:   []float64{1.0, 2.0, 3.0},
			pointB:   []float64{1.0, 2.0},
			expected: 0.0,
		},
		{
			name:     "different_points",
			pointA:   []float64{1.0, 2.0, 3.0},
			pointB:   []float64{4.0, 5.0, 6.0},
			expected: math.Sqrt(27.0),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call CalculateEuclideanDistance function with input points
			actual := CalculateEuclideanDistance(tc.pointA, tc.pointB)

			// Check if actual output matches expected output
			if math.Abs(actual-tc.expected) > 0.00001 {
				t.Errorf("Test case %q failed: expected %v, but got %v", tc.name, tc.expected, actual)
			}
		})
	}
}

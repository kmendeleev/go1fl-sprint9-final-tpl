package main

import "testing"

// Пишите тесты в этом файле

func TestGenerateRandomElements(t *testing.T) {
	for _, size := range []int{-1, 0, 1, 64} {
		data := generateRandomElements(size)
		want := size
		if want < 0 {
			want = 0
		}
		if len(data) != want {
			t.Fatalf("size %d: got %d elements, want %d", size, len(data), want)
		}
		for _, value := range data {
			if value <= 0 {
				t.Errorf("size %d: value %d is not positive", size, value)
			}
		}
	}
}

func TestMaximums(t *testing.T) {
	tests := []struct {
		name string
		data []int
		want int
	}{
		{"nil", nil, 0},
		{"empty", []int{}, 0},
		{"single", []int{7}, 7},
		{"first", []int{12, 3, 1}, 12},
		{"eight", []int{1, 3, 4, 2, 5, 6, 3, 5}, 6},
		{"remainder", []int{1, 2, 4, 1, 3, 6, 7, 4, 5, 11}, 11},
		{"negative", []int{-8, -7, -1, -6, -4, -1, -13, -6, -3}, -1},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := maximum(tc.data); got != tc.want {
				t.Errorf("maximum: got %d, want %d", got, tc.want)
			}
			if got := maxChunks(tc.data); got != tc.want {
				t.Errorf("maxChunks: got %d, want %d", got, tc.want)
			}
		})
	}
}

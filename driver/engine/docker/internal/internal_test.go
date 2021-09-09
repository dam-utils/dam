package internal

import (
	"reflect"
	"testing"
)

func TestRemoveDuplicates(t *testing.T) {
	tests := []struct {
		name     string
		elements []string
		want     []string
	}{
		{
			name:     "Nil test",
			elements: nil,
			want:     make([]string, 0),
		},
		{
			name:     "Empty test",
			elements: make([]string, 0),
			want:     make([]string, 0),
		},
		{
			name:     "Test with unique data",
			elements: []string{"Test 1", "2", "unique"},
			want:     []string{"Test 1", "2", "unique"},
		},
		{
			name: "Test with duplicate",
			elements: []string{"Test 1", "not unique element", "not unique element"},
			want:     []string{"Test 1", "not unique element"},
		},
		{
			name: "Test with only duplicates",
			elements: []string{"not unique element", "not unique element", "not unique element"},
			want:     []string{"not unique element"},
		},
	}

	for _, tt := range tests {
		t.Logf("Test: %s\n", tt.name)
		result := RemoveDuplicates(tt.elements)
		if !reflect.DeepEqual(tt.want, result) {
			t.Fatalf("Want array '%v' with result array '%v' not equal", tt.want, result)
		}
	}
}

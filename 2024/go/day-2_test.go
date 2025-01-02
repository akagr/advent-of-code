package main

import (
	"reflect"
	"testing"
)

func TestDay2(t *testing.T) {
	t.Run("parses input correctly", func(t *testing.T) {
		t.Parallel()
		input := `7 6 4 2 1
                          1 2 7 8 9`

		found := parseDay2Input(input)
		expected := [][]int{{7, 6, 4, 2, 1}, {1, 2, 7, 8, 9}}

		if !reflect.DeepEqual(found, expected) {
			t.Errorf("Found %v, expected: %v", found, expected)
		}
	})
}

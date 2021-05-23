package fflags

import (
	"fmt"
	"testing"
)

func TestIsEnabled(t *testing.T) {
	data := []struct {
		features []string
		active   []string
		want     []bool
	}{
		{[]string{"f1.remote", "f1.normal"}, []string{"f1.remote", "f1.normal", "f1.off"}, []bool{true, true, false}},
		{[]string{"f1"}, []string{"f1.remote", "f1.normal", "f1.off"}, []bool{true, true, true}},
	}
	for i, d := range data {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			// Define the features
			for _, feature := range d.features {
				Feature(feature).Enable()
			}
			// Check status
			for i, feature := range d.active {
				state := IsEnabled(Feature(feature))
				if state != d.want[i] {
					t.Errorf("Bad state for '%s': want %t, got %t", feature, d.want[i], state)
				}
			}
		})
	}
}

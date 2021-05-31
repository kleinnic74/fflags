package fflags

import (
	"fmt"
	"path/filepath"
	"testing"
)

const testdir = "testdata"

func TestLoadYaml(t *testing.T) {
	data := []struct {
		InputPath string
		Want      map[string]bool
	}{
		{
			InputPath: "basic.yaml",
			Want: map[string]bool{
				"utest.files":         false,
				"utest.network":       true,
				"debug":               true,
				"strings.utf8":        true,
				"strings.ascii":       true,
				"strings.ascii.upper": true,
			},
		},
	}
	for i, d := range data {
		t.Run(fmt.Sprintf("#%d", i), func(t *testing.T) {
			path := filepath.Join(testdir, d.InputPath)
			got, err := YamlFile(path).Load()
			if err != nil {
				t.Fatalf("Failed to load feature flags from '%s': %s", path, err)
			}
			for feature, expected := range d.Want {
				actual := got[feature]
				if actual != expected {
					t.Errorf("Bad state for feature %s: expected %t, got %t", feature, expected, actual)
					t.Logf("\tFeatures: %v", got)
				}
			}
		})
	}
}

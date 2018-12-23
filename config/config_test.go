package config

import (
	"testing"
)

// TestDefaultPath ... test for default path
func TestDefaultPath(t *testing.T) {
	dp, err := getConfigfilePath()
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("path #%v\n", dp)
}

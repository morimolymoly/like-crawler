package downloader

import (
	"testing"
)

// TestDefaultPath ... test for default path
func TestGetFilenameFromURl(t *testing.T) {
	fn, err := getFilenameFromURL("http://morimolymoly.com/files/fuck.png")
	if err != nil {
		t.Fatal(err)
	}
	if fn != "fuck.png" {
		t.Fatalf("Filename is not collect! %s", fn)
	}
}

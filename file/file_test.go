package file

import (
	"os"
	"strings"
	"testing"
)

func TestCopyR(t *testing.T) {
	os.Remove("./delete.me")
	os.Remove("/tmp/delete.me")
	os.MkdirAll("./delete.me/junk", os.ModePerm)
	data := []byte("\n\nData\n")
	os.WriteFile("./delete.me/junk/file", data, 0644)
	result, err := CopyR("./delete.me/junk/file", "/tmp")
	if err != nil {
		t.FailNow()
	}
	if result != "delete.me/junk/" {
		t.Errorf("got: %v expected: %v", result, "delete.me/junk/")
	}

	if !openRead("/tmp/delete.me/junk/file", "Data") {
		t.Errorf("file not copied")
	}

}

func openRead(file string, search string) bool {
	data, err := os.ReadFile(file)
	if err != nil {
		return false
	}
	if strings.Contains(string(data), search) {
		return true
	}
	return false

}

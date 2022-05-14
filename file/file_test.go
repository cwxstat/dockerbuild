package file

import (
	"os"
	"testing"
)

func TestCopyR(t *testing.T) {
	os.Remove("./delete.me")
	os.Remove("/tmp/delete.md")
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

}

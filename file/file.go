package file

import (
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

func CreateHomeDir(dir string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	err = os.Mkdir(home+"/"+dir, 0755)
	return err
}

func Copy(sourceFile, destinationFile string) error {
	input, err := ioutil.ReadFile(sourceFile)
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(destinationFile, input, 0644)
	if err != nil {

		return err
	}
	return nil
}

func CopyR(sourceFile, destinationDir string) (string, error) {
	dir, file := filepath.Split(sourceFile)
	clean := strings.ReplaceAll(dir, "../", "")
	clean = strings.ReplaceAll(clean, "//", "/")
	clean = strings.ReplaceAll(clean, "./", "")
	destDir := destinationDir + "/" + clean
	err := os.MkdirAll(destDir, os.ModePerm)
	if err != nil {
		return "", err
	}
	target := destDir + file
	target = strings.ReplaceAll(target, "//", "/")
	err = Copy(sourceFile, target)
	return clean, err

}

func Read(file string) (string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), err
}

func Handle(file string) (*os.File, error) {
	f, err := os.Open(file)
	return f, err
}

func HandleAppend(file string) (*os.File, error) {
	f, err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	return f, err
}

func ReadAll(f *os.File) (string, error) {
	bytesRead := 20000
	b := make([]byte, bytesRead)
	n, err := f.Read(b)
	if n >= bytesRead {
		log.Printf("file is huge: %v\n", n)
	}
	return string(b[0:n]), err
}

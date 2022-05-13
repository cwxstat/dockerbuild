package file

import (
	"io/fs"
	"log"
	"os"
)

type finfo struct {
	file string
	fi   fs.FileInfo
	err  error
}
type db struct {
	recs []finfo
}

func (d *db) Files() []string {
	result := []string{}
	for _, v := range d.recs {
		result = append(result, v.file)
	}
	return result
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

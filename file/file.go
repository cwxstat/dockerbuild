package file

import (
	"fmt"
	"io/fs"
	"log"
	"os"
	"strings"
)

type File struct {
	Name    string
	Pod     string
	Config  string
	Version int
	f       []*os.File
}

func NewFile() (*File, error) {
	return &File{}, nil
}

type finfo struct {
	file string
	fi   fs.FileInfo
	err  error
}
type db struct {
	recs []finfo
}
func (d *db)Files() []string {
	result := []string{}
	for _,v :=range d.recs {
		result = append(result,v.file)
	}
	return result
}

func (f *File) Exist(dockerFile ...string) *db {
	d := &db{}
	for _, v := range dockerFile {
		data, err := os.Stat(v)
		d.recs = append(d.recs, finfo{file: v, fi: data, err: err})

	}
	return d
}

func (f *File) Handle(dockerFiles ...string) {
	var files []string
	var err error

	if len(dockerFiles) == 0 {
		files = append(files, "Dockerfile")
	} else {
		for _, v := range dockerFiles {
			files = append(files, v)
		}
	}

	f.f[0], err = os.OpenFile(files[0], os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return
	}
}

func Read(file string) (string, error) {
	dat, err := os.ReadFile(file)
	if err != nil {
		return "", err
	}
	return string(dat), err
}

func Handle(file string) (*os.File, error) {
	f,err := os.Open(file)
	return f,err
}

func HandleAppend(file string) (*os.File, error) {
	f,err := os.OpenFile(file, os.O_APPEND|os.O_WRONLY, 0644)
	return f,err
}

func ReadAll(f *os.File) (string, error) {
	bytesRead := 20000
	b := make([]byte, bytesRead)
	n,err := f.Read(b)
	if n >= bytesRead{
		log.Printf("file is huge: %v\n",n)
	}
	return string(b[0:n]),err
}

func GrabTab(s string,tagBegin, tagEnd string) (string,string, error) {
	split := strings.Split(s, "\n")
	startLine := 0
	endLine := 0 
	for i,v := range split {
		if strings.Contains(v, tagBegin) {
			startLine = i
			break
		}
	}

	for i, v := range split {
		if strings.Contains(v, tagEnd) {
			endLine = i
			break
		}
	}

	if startLine >= endLine {
		return s,"",fmt.Errorf("Start > End")
	}
	sub := ""
	sep := ""
	for i:=0;i<startLine; i++ {
		sub += split[i]
		sub = sub + sep + split[i]
		sep = "\n"
	}

	tag := ""
	sep = ""
	for i:=startLine;i<=endLine; i++ {
		tag = tag + sep + split[i]
		sep = "\n"
	}

	return sub,tag,nil

}

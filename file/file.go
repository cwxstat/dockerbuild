package file

import (
	"io/fs"
	"os"
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

func (f *File) Exist(dockerFile ...string) []finfo {
	var db []finfo
	for _, v := range dockerFile {
		data, err := os.Stat(v)
		db = append(db, finfo{file: v, fi: data, err: err})

	}
	return db
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

package file

import (
	"fmt"
	"os"
	"testing"
)

func TestDev(t *testing.T) {
	f, _ := Handle("DockerTest")
	s, _ := ReadAll(f)
	fmt.Println(s)
}

func TestFile_Exist(t *testing.T) {
	type fields struct {
		Name    string
		Pod     string
		Config  string
		Version int
		f       []*os.File
	}
	type args struct {
		dockerFile []string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   []finfo
	}{
		// TODO: Add test cases.
		{
			name:   "File test",
			fields: fields{},
			args: args{
				dockerFile: []string{"Dockerfile"},
			},
			want: []finfo{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &File{
				Name:    tt.fields.Name,
				Pod:     tt.fields.Pod,
				Config:  tt.fields.Config,
				Version: tt.fields.Version,
				f:       tt.fields.f,
			}
			got := f.Exist(tt.args.dockerFile...)
			files := got.Files()
			if len(files) != 1 {
				t.Errorf("error: %v", files)
			}
		})
	}
}

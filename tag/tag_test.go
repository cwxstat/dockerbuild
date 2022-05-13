package tag

import (
	"os"
	"testing"

	"github.com/cwxstat/dopt/uyaml"
)

func Test_tag_writeTag(t *testing.T) {

	testFile := "DockerTest.delete"
	type args struct {
		filename string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Test Write",
			args: args{
				filename: testFile,
			},
			wantErr: false,
		},
	}
	err := os.WriteFile(testFile, []byte("FROM golang:latest AS build\n\n"), 0600)
	if err != nil {
		t.Error(err)
	}

	d := newTag()
	d.addTagIfNeeded(testFile)
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := newTag()
			if err := tr.writeTag(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", err, tt.wantErr)
			}
			dy := &uyaml.TopYaml{
				Image:   "spud2",
				Version: "v0.0.2",
				Config:  "",
				Spec:    uyaml.Spec{},
			}

			tr.UpdateDY(dy)
			if err := tr.writeTag(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", err, tt.wantErr)
			}

			dy.NextMinor()
			dy.NextMinor()
			tr.UpdateDY(dy)
			if err := tr.writeTag(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", err, tt.wantErr)
			}
			dy = tr.Dy()
			if dy.Version != "v0.0.4" {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", dy.Version, "v0.0.4")
			}
			dy.Version = "v0.0.10"
			tr.UpdateDY(dy)
			if err := tr.writeTag(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", err, tt.wantErr)
			}
			if err := tr.readTag(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", err, tt.wantErr)
			}
			dy = tr.Dy()
			if dy.Version != "v0.0.10" {
				t.Errorf("tag.writeTag() error = %v, wantErr %v", dy.Version, "v0.0.10")
			}

		})
	}
}

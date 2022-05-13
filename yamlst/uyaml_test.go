package yamlst

import (
	"os"
	"testing"
)

func TestMyTest(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "test",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MyTest()
		})
	}
}

func TestAddComments(t *testing.T) {

	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{
			name: "Quick test",
			args: args{},
			want: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := addComments(MyTest())
			os.WriteFile("sampleComment.yaml", []byte(got), 0644)
			got = removeComments(got)
			os.WriteFile("sampleUnComment.yaml", []byte(got), 0644)

		})
	}
}

func TestTopYaml_NextMinor(t *testing.T) {
	type fields struct {
		Image   string
		Version string
		Config  string
		Spec    Spec
	}
	tests := []struct {
		name    string
		fields  fields
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "Version test",
			fields: fields{
				Image:   "",
				Version: "v0.0.1",
				Config:  "",
				Spec:    Spec{},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dy := &TopYaml{
				Image:   tt.fields.Image,
				Version: tt.fields.Version,
				Config:  tt.fields.Config,
				Spec:    tt.fields.Spec,
			}
			if err := dy.NextMinor(); (err != nil) != tt.wantErr {
				t.Errorf("TopYaml.NextMinor() error = %v, wantErr %v", err, tt.wantErr)
			}
			if dy.Version != "v0.0.2" {
				t.Errorf("%s wanted %s", dy.Version, "v0.0.2")
			}
			if err := dy.NextMajor(); (err != nil) != tt.wantErr {
				t.Errorf("TopYaml.NextMinor() error = %v, wantErr %v", err, tt.wantErr)
			}
			if dy.Version != "v0.1.2" {
				t.Errorf("%s wanted %s", dy.Version, "v0.1.2")
			}

		})
	}
}

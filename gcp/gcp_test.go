package gcp

import (
	"fmt"
	"testing"
)

func TestMain(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if _, err := Repos("projects/mchirico/locations/us-central1"); err != nil {
				if err == ErrNewClient {
					return
				}
				t.Errorf("%v\n", err)

			}
		})
	}
}

func TestFiles(t *testing.T) {
	type args struct {
		parent string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				parent: "projects/mchirico/locations/us-central1/repositories/public",
			},
			want:    []string{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Files(tt.args.parent)
			if (err != nil) != tt.wantErr {
				if err == ErrNewClient {
					return
				}
				t.Errorf("Files() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			fmt.Println(got)
		})
	}
}

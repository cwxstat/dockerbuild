package gcp

import (
	"fmt"
	"google.golang.org/grpc/status"
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

func TestDelete(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name: "",
			args: args{
				name: "projects/mchirico/locations/us-central1/repositories/public/packages/septa/tags/test",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Delete(tt.args.name); (err != nil) != tt.wantErr {
				if err == ErrNewClient {
					return
				}
				if st, ok := status.FromError(err); ok {

					if err == ErrNewClient {
						return
					}
					if st.Message() == "Requested entity was not found." {
						return
					}
				}
				t.Errorf("Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

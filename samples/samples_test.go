package samples

import (
	"os"
	"testing"
)

func TestCreateSample(t *testing.T) {

	type args struct {
		file string
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
				file: "DockerfileTest",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		os.Remove("DockerfileTest")
		t.Run(tt.name, func(t *testing.T) {
			if err := CreateSample(tt.args.file); (err != nil) != tt.wantErr {
				t.Errorf("CreateSample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

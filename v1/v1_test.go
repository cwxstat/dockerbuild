package v1

import (
	"github.com/cwxstat/dopt/samples"
	"os"
	"testing"
)

func TestInit(t *testing.T) {
	type args struct {
		file string
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
		{
			name: "V1 test",
			args: args{
				file: "Dockerfile.golang.delete.me",
			},
		},
	}
	for _, tt := range tests {
		os.Remove(tt.args.file)
		samples.CreateSample(tt.args.file)
		t.Run(tt.name, func(t *testing.T) {
			Init(tt.args.file)
		})

	}
}

package uyaml

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
			got := AddComments(MyTest())
			os.WriteFile("sampleComment.yaml", []byte(got), 0644)
			got = RemoveComments(got)
			os.WriteFile("sampleUnComment.yaml", []byte(got), 0644)

		})
	}
}

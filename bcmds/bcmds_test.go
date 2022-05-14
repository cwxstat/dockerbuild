package bcmds

import (
	"testing"
)

func TestBash(t *testing.T) {
	cmd := "ls -l"
	Bash(cmd)
}

func TestExample(t *testing.T) {
	tests := []struct {
		name string
	}{
		// TODO: Add test cases.
		{
			name: "Smoke",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			Docker("ps", "-a")
		})
	}
}

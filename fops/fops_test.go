package fops

import "testing"

func TestFOPS_Sample(t *testing.T) {
	tests := []struct {
		name    string
		f       *FOPS
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			f:       &FOPS{},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FOPS{}
			if err := f.Sample(); (err != nil) != tt.wantErr {
				t.Errorf("FOPS.Sample() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

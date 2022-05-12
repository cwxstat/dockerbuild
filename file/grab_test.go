package file

import (
	"testing"
)

var testString = `
one
2
3
# <docb: >
# This is center
# </docb:>
		
`

func TestGrabTab(t *testing.T) {

	type args struct {
		s        string
		tagBegin string
		tagEnd   string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		want1   string
		wantErr bool
	}{
		// TODO: Add test cases.

		{
			name: "Smoke",
			args: args{
				s:        testString,
				tagBegin: "<docb:",
				tagEnd:   "</docb",
			},
			want:    "\none\n2\n3",
			want1:   "# <docb: >\n# This is center\n# </docb:>",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := GrabTab(tt.args.s, tt.args.tagBegin, tt.args.tagEnd)
			if (err != nil) != tt.wantErr {
				t.Errorf("GrabTab() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("GrabTab() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GrabTab() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestReadMiddle(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name    string
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
		{
			name:    "",
			args:    args{},
			want:    "# This is center",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, tag, err := GrabTab(testString, "<docb:", "</docb:")
			if err != nil {
				t.Errorf("ReadMiddle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			got, err := ReadMiddle(tag)
			if (err != nil) != tt.wantErr {
				t.Errorf("ReadMiddle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ReadMiddle() = %v, want %v", got, tt.want)
			}
		})
	}
}

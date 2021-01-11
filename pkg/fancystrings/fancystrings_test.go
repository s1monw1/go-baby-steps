package fancystrings

import "testing"

func TestEmojify(t *testing.T) {
	type args struct {
		original string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"unknown chars stay", args{"/&%"}, "/&%"},
		{"alphabetical chars are replaces", args{"abc"}, "😎🙃😍"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Emojify(tt.args.original); got != tt.want {
				t.Errorf("Emojify() = %v, want %v", got, tt.want)
			}
		})
	}
}

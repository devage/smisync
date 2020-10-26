// main of smisync
// 2013.03.21 ikpark@gmail.com

package main

import "testing"

func Test_smiSync(t *testing.T) {
	type args struct {
		line   string
		start  int
		offset int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
		{"If the input doesn't have '<SYNC ' tag, it must be identical to the output.", args{"plain text", 0, 0}, "plain text"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smiSync(tt.args.line, tt.args.start, tt.args.offset); got != tt.want {
				t.Errorf("smiSync() = %v, want %v", got, tt.want)
			}
		})
	}
}

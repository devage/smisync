// main of smisync
// 2013.03.21 ikpark@gmail.com

package main

import "testing"

func Test_parseOption(t *testing.T) {
	type args struct {
		option string
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{"'start+offset' 포맷을 파싱할 수 있어야 합니다.", args{"1+2"}, 1, 2},
		{"'start-offset' 포맷을 파싱할 수 있어야 합니다. `offset`은 음의 정수가 됩니다.", args{"1-2"}, 1, -2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := parseOption(tt.args.option)
			if got != tt.want {
				t.Errorf("parseOption() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("parseOption() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

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
		{"'<SYNC Start=' 태그가 없는 입력문은 수정 없이 출력돼야 합니다.", args{"plain text", 0, 0}, "plain text"},
		{"SYNC 태그의 시간은 `offset` 값만큼 더해져야 합니다.", args{"<SYNC Start=0>", 0, 1000}, "<SYNC Start=1000>"},
		{"SYNC 태그의 시간은 `offset` 값만큼 더해져야 합니다 (음수).", args{"<SYNC Start=1000>", 0, -1000}, "<SYNC Start=0>"},
		{"`start` 변수에 명시된 시간 이전의 SYNC 시간은 수정되지 않아야 합니다.", args{"<SYNC Start=100>", 101, 1000}, "<SYNC Start=100>"},
		{"'<SYNC Start=' 태그는 대소문자 구분을 하지 않습니다.", args{"<sync start=0>", 0, 1000}, "<sync start=1000>"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := smiSync(tt.args.line, tt.args.start, tt.args.offset); got != tt.want {
				t.Errorf("smiSync() = \"%v\", want \"%v\"", got, tt.want)
			}
		})
	}
}

package gotool

import "testing"

func TestComma(t *testing.T) {
	type args struct {
		v int64
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "", args: args{v: 999}, want: "999"},
		{name: "", args: args{v: 9999}, want: "9,999"},
		{name: "", args: args{v: 9999999}, want: "9,999,999"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Comma(tt.args.v); got != tt.want {
				t.Errorf("Comma() = %v, want %v", got, tt.want)
			}
		})
	}
}

package unsafehelper

import (
	"testing"
)

func TestByteInt8(t *testing.T) {
	type args struct {
		b byte
	}
	tests := []struct {
		name string
		args args
		want int8
	}{
		{name: "0",
			args: args{b: 0},
			want: 0,
		},
		{name: "1",
			args: args{b: 1},
			want: 1,
		},
		{name: "127",
			args: args{b: 127},
			want: 127,
		},
		{name: "128",
			args: args{b: 128},
			want: -128,
		},
		{name: "255",
			args: args{b: 255},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *ByteInt8(&tt.args.b); got != tt.want {
				t.Errorf("ByteInt8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt8Byte(t *testing.T) {
	type args struct {
		i8 int8
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		{name: "0",
			args: args{i8: 0},
			want: 0,
		},
		{name: "1",
			args: args{i8: 1},
			want: 1,
		},
		{name: "127",
			args: args{i8: 127},
			want: 127,
		},
		{name: "-128",
			args: args{i8: -128},
			want: 128,
		},
		{name: "-1",
			args: args{i8: -1},
			want: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int8Byte(&tt.args.i8); got != tt.want {
				t.Errorf("Int8Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

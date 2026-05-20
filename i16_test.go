package unsafehelper

import (
	"reflect"
	"testing"

	"github.com/solsw/errorhelper"
)

func TestInt16Uint16(t *testing.T) {
	type args struct {
		i16 int16
	}
	tests := []struct {
		name string
		args args
		want uint16
	}{
		{name: "0",
			args: args{i16: 0},
			want: 0,
		},
		{name: "1",
			args: args{i16: 1},
			want: 1,
		},
		{name: "-1",
			args: args{i16: -1},
			want: 0b_1111111111111111,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int16Uint16(&tt.args.i16); got != tt.want {
				t.Errorf("Int16Uint16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Byte2(t *testing.T) {
	type args struct {
		i16 int16
	}
	tests := []struct {
		name string
		args args
		want [2]byte
	}{
		{name: "1",
			args: args{i16: 1},
			want: [2]byte{1, 0},
		},
		{name: "256",
			args: args{i16: 256},
			want: [2]byte{0, 1},
		},
		{name: "0x1122",
			args: args{i16: 0x1122},
			want: [2]byte{0x22, 0x11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int16Byte2(&tt.args.i16); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int16Byte2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte2Int16(t *testing.T) {
	type args struct {
		b2 [2]byte
	}
	tests := []struct {
		name string
		args args
		want int16
	}{
		{name: "1",
			args: args{b2: [2]byte{1, 0}},
			want: 1,
		},
		{name: "256",
			args: args{b2: [2]byte{0, 1}},
			want: 256,
		},
		{name: "0x1122",
			args: args{b2: [2]byte{0x22, 0x11}},
			want: 0x1122,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Byte2Int16(&tt.args.b2); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte2Int16() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt16Byte(t *testing.T) {
	type args struct {
		i16 int16
		i   int
	}
	tests := []struct {
		name              string
		args              args
		wantErrPanicError string
		want              byte
	}{
		{name: "index out of range 1",
			args: args{
				i16: 0x1122,
				i:   -1,
			},
			wantErrPanicError: "runtime error: index out of range [-1]",
		},
		{name: "index out of range 2",
			args: args{
				i16: 0x1122,
				i:   2,
			},
			wantErrPanicError: "runtime error: index out of range [2] with length 2",
		},
		{name: "0",
			args: args{i16: 0x1122, i: 0},
			want: 0x22,
		},
		{name: "1",
			args: args{i16: 0x1122, i: 1},
			want: 0x11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var got byte
			gotErrPanic := func() (errPanic error) {
				defer func() {
					errorhelper.PanicToError(recover(), &errPanic)
				}()
				got = Int16Byte(&tt.args.i16, tt.args.i)
				return nil
			}()
			if tt.wantErrPanicError != "" {
				gotErrPanicError := ""
				if gotErrPanic != nil {
					gotErrPanicError = gotErrPanic.Error()
				}
				if gotErrPanicError != tt.wantErrPanicError {
					t.Errorf("Int16Byte() panic = '%v', want '%v'", gotErrPanicError, tt.wantErrPanicError)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Int16Byte() = %#x, want %#x", got, tt.want)
			}
		})
	}
}

package unsafehelper

import (
	"reflect"
	"testing"

	"github.com/solsw/errorhelper"
)

func TestInt64Uint64(t *testing.T) {
	type args struct {
		i64 int64
	}
	tests := []struct {
		name string
		args args
		want uint64
	}{
		{name: "0",
			args: args{i64: 0},
			want: 0,
		},
		{name: "1",
			args: args{i64: 1},
			want: 1,
		},
		{name: "-1",
			args: args{i64: -1},
			want: ^uint64(0),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int64Uint64(&tt.args.i64); got != tt.want {
				t.Errorf("Int64Uint64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Byte8(t *testing.T) {
	type args struct {
		i64 int64
	}
	tests := []struct {
		name string
		args args
		want [8]byte
	}{
		{name: "1",
			args: args{i64: 1},
			want: [8]byte{1, 0, 0, 0, 0, 0, 0, 0},
		},
		{name: "256",
			args: args{i64: 256},
			want: [8]byte{0, 1, 0, 0, 0, 0, 0, 0},
		},
		{name: "65536",
			args: args{i64: 65536},
			want: [8]byte{0, 0, 1, 0, 0, 0, 0, 0},
		},
		{name: "0x1122334455667788",
			args: args{i64: 0x1122334455667788},
			want: [8]byte{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int64Byte8(&tt.args.i64); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int64Byte8() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte8Int64(t *testing.T) {
	type args struct {
		b8 [8]byte
	}
	tests := []struct {
		name string
		args args
		want int64
	}{
		{name: "1",
			args: args{b8: [8]byte{1, 0, 0, 0, 0, 0, 0, 0}},
			want: 1,
		},
		{name: "256",
			args: args{b8: [8]byte{0, 1, 0, 0, 0, 0, 0, 0}},
			want: 256,
		},
		{name: "65536",
			args: args{b8: [8]byte{0, 0, 1, 0, 0, 0, 0, 0}},
			want: 65536,
		},
		{name: "0x1122334455667788",
			args: args{b8: [8]byte{0x88, 0x77, 0x66, 0x55, 0x44, 0x33, 0x22, 0x11}},
			want: 0x1122334455667788,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Byte8Int64(&tt.args.b8); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte8Int64() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt64Byte(t *testing.T) {
	type args struct {
		i64 int64
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
				i64: 0x1122334455667788,
				i:   -1,
			},
			wantErrPanicError: "runtime error: index out of range [-1]",
		},
		{name: "index out of range 2",
			args: args{
				i64: 0x1122334455667788,
				i:   8,
			},
			wantErrPanicError: "runtime error: index out of range [8] with length 8",
		},
		{name: "0",
			args: args{i64: 0x1122334455667788, i: 0},
			want: 0x88,
		},
		{name: "1",
			args: args{i64: 0x1122334455667788, i: 1},
			want: 0x77,
		},
		{name: "6",
			args: args{i64: 0x1122334455667788, i: 6},
			want: 0x22,
		},
		{name: "7",
			args: args{i64: 0x1122334455667788, i: 7},
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
				got = Int64Byte(&tt.args.i64, tt.args.i)
				return nil
			}()
			if tt.wantErrPanicError != "" {
				gotErrPanicError := ""
				if gotErrPanic != nil {
					gotErrPanicError = gotErrPanic.Error()
				}
				if gotErrPanicError != tt.wantErrPanicError {
					t.Errorf("Int64Byte() panic = '%v', want '%v'", gotErrPanicError, tt.wantErrPanicError)
				}
				return
			}
			if got != tt.want {
				t.Errorf("Int64Byte() = %#x, want %#x", got, tt.want)
			}
		})
	}
}

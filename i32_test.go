package unsafehelper

import (
	"reflect"
	"testing"

	"github.com/solsw/builtinhelper"
)

func TestInt32Uint32(t *testing.T) {
	type args struct {
		i32 int32
	}
	tests := []struct {
		name string
		args args
		want uint32
	}{
		{name: "0", args: args{i32: 0}, want: 0},
		{name: "1", args: args{i32: 1}, want: 1},
		{name: "-1", args: args{i32: -1}, want: 0b_11111111111111111111111111111111},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int32Uint32(&tt.args.i32); got != tt.want {
				t.Errorf("Int32Uint32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Byte4(t *testing.T) {
	type args struct {
		i32 int32
	}
	tests := []struct {
		name string
		args args
		want [4]byte
	}{
		{name: "1", args: args{i32: 1}, want: [4]byte{1, 0, 0, 0}},
		{name: "256", args: args{i32: 256}, want: [4]byte{0, 1, 0, 0}},
		{name: "65536", args: args{i32: 65536}, want: [4]byte{0, 0, 1, 0}},
		{name: "12345678", args: args{i32: 12345678}, want: [4]byte{78, 97, 188, 0}},
		{name: "0x11223344", args: args{i32: 0x11223344}, want: [4]byte{0x44, 0x33, 0x22, 0x11}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Int32Byte4(&tt.args.i32); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Int32Byte4() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestByte4Int32(t *testing.T) {
	type args struct {
		bb [4]byte
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{name: "1", args: args{bb: [4]byte{1, 0, 0, 0}}, want: 1},
		{name: "256", args: args{bb: [4]byte{0, 1, 0, 0}}, want: 256},
		{name: "65536", args: args{bb: [4]byte{0, 0, 1, 0}}, want: 65536},
		{name: "12345678", args: args{bb: [4]byte{78, 97, 188, 0}}, want: 12345678},
		{name: "0x11223344", args: args{bb: [4]byte{0x44, 0x33, 0x22, 0x11}}, want: 0x11223344},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := *Byte4Int32(&tt.args.bb); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Byte4Int32() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Byte(t *testing.T) {
	type args struct {
		i32 int32
		i   int
	}
	tests := []struct {
		name string
		args args
		want byte
	}{
		// {name: "-1", args: args{i32: 0x44332211, i: -1}, want: 0x11},
		{name: "0", args: args{i32: 0x44332211, i: 0}, want: 0x11},
		{name: "1", args: args{i32: 0x44332211, i: 1}, want: 0x22},
		{name: "2", args: args{i32: 0x44332211, i: 2}, want: 0x33},
		{name: "3", args: args{i32: 0x44332211, i: 3}, want: 0x44},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Int32Byte(&tt.args.i32, tt.args.i); got != tt.want {
				t.Errorf("Int32Byte() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInt32Byte_panic(t *testing.T) {
	type args struct {
		i32 int32
		i   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{name: "-1",
			args: args{
				i32: 0x44332211,
				i:   -1,
			},
			want: "runtime error: index out of range [-1]",
		},
		{name: "12",
			args: args{
				i32: 0x44332211,
				i:   12,
			},
			want: "runtime error: index out of range [12] with length 4",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErr := func() (err error) {
				defer func() {
					builtinhelper.PanicToError(recover(), &err)
				}()
				Int32Byte(&tt.args.i32, tt.args.i)
				return nil
			}()
			got := gotErr.Error()
			if got != tt.want {
				t.Errorf("Runtime_panic = '%v', want '%v'", got, tt.want)
			}
		})
	}
}

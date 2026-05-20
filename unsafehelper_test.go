package unsafehelper

import (
	"testing"

	"github.com/solsw/errorhelper"
)

func TestCastSizeMismatch(t *testing.T) {
	tests := []struct {
		name              string
		wantErrPanicError string
		fn                func()
	}{
		{name: "byte to int32",
			wantErrPanicError: "unsafehelper.Cast: From and To have different sizes",
			fn: func() {
				var b byte
				_ = Cast[byte, int32](&b)
			},
		},
		{name: "int64 to int16",
			wantErrPanicError: "unsafehelper.Cast: From and To have different sizes",
			fn: func() {
				var i int64
				_ = Cast[int64, int16](&i)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotErrPanic := func() (errPanic error) {
				defer func() {
					errorhelper.PanicToError(recover(), &errPanic)
				}()
				tt.fn()
				return nil
			}()
			gotErrPanicError := ""
			if gotErrPanic != nil {
				gotErrPanicError = gotErrPanic.Error()
			}
			if gotErrPanicError != tt.wantErrPanicError {
				t.Errorf("Cast() panic = '%v', want '%v'", gotErrPanicError, tt.wantErrPanicError)
			}
		})
	}
}

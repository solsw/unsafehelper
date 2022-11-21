package unsafehelper

import (
	"unsafe"
)

// Cast converts value of *From type to value of *To type.
func Cast[From, To any](f *From) *To {
	return (*To)(unsafe.Pointer(f))
}

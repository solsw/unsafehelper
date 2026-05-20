package unsafehelper

import (
	"unsafe"
)

// Cast converts value of *From type to value of *To type.
// Cast panics if From and To have different sizes.
func Cast[From, To any](f *From) *To {
	if unsafe.Sizeof(*new(From)) != unsafe.Sizeof(*new(To)) {
		panic("unsafehelper.Cast: From and To have different sizes")
	}
	return (*To)(unsafe.Pointer(f))
}

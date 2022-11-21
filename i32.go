package unsafehelper

// Int32Uint32 casts int32 to uint32.
func Int32Uint32(i32 *int32) *uint32 {
	return Cast[int32, uint32](i32)
}

// Int32Byte4 casts int32 to [4]byte.
func Int32Byte4(i32 *int32) *[4]byte {
	return Cast[int32, [4]byte](i32)
}

// Byte4Int32 casts [4]byte to int32.
func Byte4Int32(b4 *[4]byte) *int32 {
	return Cast[[4]byte, int32](b4)
}

// Int32Byte returns 'i'-th byte from int32.
// Int32Byte panics if 'i' is out of [0..3] range.
func Int32Byte(i32 *int32, i int) byte {
	return (*Int32Byte4(i32))[i]
}

// Uint32Int32 casts uint32 to int32.
func Uint32Int32(u32 *uint32) *int32 {
	return Cast[uint32, int32](u32)
}

// Uint32Byte4 casts uint32 to [4]byte.
func Uint32Byte4(u32 *uint32) *[4]byte {
	return Cast[uint32, [4]byte](u32)
}

// Byte4Uint32 casts [4]byte to uint32.
func Byte4Uint32(b4 *[4]byte) *uint32 {
	return Cast[[4]byte, uint32](b4)
}

// Uint32Byte returns 'i'-th byte from uint32.
// Uint32Byte panics if 'i' is out of [0..3] range.
func Uint32Byte(u32 *uint32, i int) byte {
	return (*Uint32Byte4(u32))[i]
}

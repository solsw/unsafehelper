package unsafehelper

// Int16Uint16 casts int16 to uint16.
func Int16Uint16(i16 *int16) *uint16 {
	return Cast[int16, uint16](i16)
}

// Int16Byte2 casts int16 to [2]byte.
func Int16Byte2(i16 *int16) *[2]byte {
	return Cast[int16, [2]byte](i16)
}

// Byte2Int16 casts [2]byte to int16.
func Byte2Int16(b2 *[2]byte) *int16 {
	return Cast[[2]byte, int16](b2)
}

// Int16Byte returns 'i'-th byte from int16.
// Int16Byte panics if 'i' is out of [0..1] range.
func Int16Byte(i16 *int16, i int) byte {
	return (*Int16Byte2(i16))[i]
}

// Uint16Int16 casts uint16 to int16.
func Uint16Int16(u16 *uint16) *int16 {
	return Cast[uint16, int16](u16)
}

// Uint16Byte2 casts uint16 to [2]byte.
func Uint16Byte2(u16 *uint16) *[2]byte {
	return Cast[uint16, [2]byte](u16)
}

// Byte2Uint16 casts [2]byte to uint16.
func Byte2Uint16(b2 *[2]byte) *uint16 {
	return Cast[[2]byte, uint16](b2)
}

// Uint16Byte returns 'i'-th byte from uint16.
// Uint16Byte panics if 'i' is out of [0..1] range.
func Uint16Byte(u16 *uint16, i int) byte {
	return (*Uint16Byte2(u16))[i]
}

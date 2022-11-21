package unsafehelper

// ByteInt8 casts byte to int8.
func ByteInt8(b *byte) *int8 {
	return Cast[byte, int8](b)
}

// Int8Byte casts int8 to byte.
func Int8Byte(i8 *int8) *byte {
	return Cast[int8, byte](i8)
}

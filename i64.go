package unsafehelper

// Int64Uint64 casts [builtin.int64] to [builtin.uint64].
func Int64Uint64(i64 *int64) *uint64 {
	return Cast[int64, uint64](i64)
}

// Int64Byte8 casts [builtin.int64] to [8]byte.
func Int64Byte8(i64 *int64) *[8]byte {
	return Cast[int64, [8]byte](i64)
}

// Byte8Int64 casts [8]byte to [builtin.int64].
func Byte8Int64(b8 *[8]byte) *int64 {
	return Cast[[8]byte, int64](b8)
}

// Int64Byte returns 'i'-th [builtin.byte] from [builtin.int64].
// Int64Byte panics if 'i' is out of [0..7] range.
func Int64Byte(i64 *int64, i int) byte {
	return (*Int64Byte8(i64))[i]
}

// Uint64Int64 casts [builtin.uint64] to [builtin.int64].
func Uint64Int64(u64 *uint64) *int64 {
	return Cast[uint64, int64](u64)
}

// Uint64Byte8 casts [builtin.uint64] to [8]byte.
func Uint64Byte8(u64 *uint64) *[8]byte {
	return Cast[uint64, [8]byte](u64)
}

// Byte8Uint64 casts [8]byte to [builtin.uint64].
func Byte8Uint64(b8 *[8]byte) *uint64 {
	return Cast[[8]byte, uint64](b8)
}

// Uint64Byte returns 'i'-th [builtin.byte] from [builtin.uint64].
// Uint64Byte panics if 'i' is out of [0..7] range.
func Uint64Byte(u64 *uint64, i int) byte {
	return (*Uint64Byte8(u64))[i]
}

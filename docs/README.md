# unsafehelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/unsafehelper.svg)](https://pkg.go.dev/github.com/solsw/unsafehelper)
[![GitHub](https://img.shields.io/badge/github--green?logo=github)](https://github.com/solsw/unsafehelper)

Helpers for Go's [`unsafe`](https://pkg.go.dev/unsafe) package.

The package centers on a single generic primitive, `Cast`, which reinterprets a
pointer of one type as a pointer of another type of identical size. A set of
thin, type-specific wrappers builds on `Cast` to cover the common numeric and
byte-layout conversions.

## Installation

```sh
go get github.com/solsw/unsafehelper
```

```go
import "github.com/solsw/unsafehelper"
```

## The `Cast` primitive

```go
func Cast[From, To any](f *From) *To
```

`Cast` returns the `*From` pointer reinterpreted as a `*To` pointer. The two
types must occupy the same number of bytes — `Cast` **panics** if
`unsafe.Sizeof(From) != unsafe.Sizeof(To)`. The check happens at runtime, since
sizes are not known to the type system.

Because `Cast` returns a pointer into the *same* memory, writing through the
result mutates the original value:

```go
i := int32(-1)
u := unsafehelper.Cast[int32, uint32](&i)
fmt.Println(*u) // 4294967295
```

## Convenience wrappers

All wrappers are zero-overhead calls into `Cast` with the type parameters fixed.
They fall into three groups.

### Signed ⇄ unsigned

Reinterpret a value as its same-width counterpart of the opposite signedness.
The underlying bits are untouched.

| Function | Converts |
|----------|----------|
| `ByteInt8`   / `Int8Byte`     | `byte` ⇄ `int8` |
| `Int16Uint16` / `Uint16Int16` | `int16` ⇄ `uint16` |
| `Int32Uint32` / `Uint32Int32` | `int32` ⇄ `uint32` |
| `Int64Uint64` / `Uint64Int64` | `int64` ⇄ `uint64` |

### Value ⇄ fixed-size byte array

View an integer as its raw byte array, or a byte array as an integer. The bytes
follow the platform's native memory layout (little-endian on x86/amd64).

| Function | Converts |
|----------|----------|
| `Int16Byte2`  / `Byte2Int16`  | `int16`  ⇄ `[2]byte` |
| `Uint16Byte2` / `Byte2Uint16` | `uint16` ⇄ `[2]byte` |
| `Int32Byte4`  / `Byte4Int32`  | `int32`  ⇄ `[4]byte` |
| `Uint32Byte4` / `Byte4Uint32` | `uint32` ⇄ `[4]byte` |
| `Int64Byte8`  / `Byte8Int64`  | `int64`  ⇄ `[8]byte` |
| `Uint64Byte8` / `Byte8Uint64` | `uint64` ⇄ `[8]byte` |

### Indexed byte access

Return the `i`-th byte of a value directly. Each function **panics** if `i` is
outside the valid range for the type's width.

| Function | Value type | Valid `i` |
|----------|-----------|-----------|
| `Int16Byte`  / `Uint16Byte` | `int16`  / `uint16` | `0..1` |
| `Int32Byte`  / `Uint32Byte` | `int32`  / `uint32` | `0..3` |
| `Int64Byte`  / `Uint64Byte` | `int64`  / `uint64` | `0..7` |

## Examples

Signed/unsigned reinterpretation:

```go
u := uint16(0xFFFF)
i := unsafehelper.Uint16Int16(&u)
fmt.Println(*i) // -1
```

Inspecting the byte layout of a value:

```go
v := int32(0x04030201)
b := unsafehelper.Int32Byte4(&v)
fmt.Printf("% x\n", *b) // 01 02 03 04   (little-endian)
```

Reading a single byte:

```go
v := uint32(0x04030201)
fmt.Printf("%#x\n", unsafehelper.Uint32Byte(&v, 0)) // 0x1
```

## Notes & caveats

- These helpers use `unsafe.Pointer` and bypass Go's type safety. Use them only
  when you understand the memory implications.
- Byte-array and indexed-byte results depend on the host's **endianness**; code
  that must be portable across architectures should account for this.
- Returned pointers alias the original value's storage; mutations are shared in
  both directions.

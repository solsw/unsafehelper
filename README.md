# unsafehelper
[![Go Reference](https://pkg.go.dev/badge/github.com/solsw/unsafehelper.svg)](https://pkg.go.dev/github.com/solsw/unsafehelper)

Helpers for Go's [unsafe](https://pkg.go.dev/unsafe) package.

## Overview

The package provides a generic `Cast` function for reinterpreting a pointer of one type as a pointer of another type of the same size. `Cast` panics at runtime if the sizes of `From` and `To` differ.

Convenience wrappers cover the common numeric conversions:

| File | Functions |
|------|-----------|
| `i8.go` | `ByteInt8`, `Int8Byte` |
| `i16.go` | `Int16Uint16`, `Uint16Int16`, `Int16Byte2`, `Byte2Int16`, `Uint16Byte2`, `Byte2Uint16`, `Int16Byte`, `Uint16Byte` |
| `i32.go` | `Int32Uint32`, `Uint32Int32`, `Int32Byte4`, `Byte4Int32`, `Uint32Byte4`, `Byte4Uint32`, `Int32Byte`, `Uint32Byte` |
| `i64.go` | `Int64Uint64`, `Uint64Int64`, `Int64Byte8`, `Byte8Int64`, `Uint64Byte8`, `Byte8Uint64`, `Int64Byte`, `Uint64Byte` |

Byte-array functions reflect the native (little-endian on x86/amd64) memory layout of the value.

## Requirements

Go 1.18 or later (generics).

## Installation

```
go get github.com/solsw/unsafehelper
```

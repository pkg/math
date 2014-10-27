
# math
    import "github.com/pkg/math"

Package math provides helper functions for mathematical operations
over all integer Go types.

Almost all files in this package are automatically generated.

To regenerate this package


	make -B

This package relies on github.com/davecheney/godoc2md.






## func EqualBigInt
``` go
func EqualBigInt(a, b *big.Int) bool
```
EqualBigInt returns true if both *big.Ints are equal


## func EqualBigRat
``` go
func EqualBigRat(a, b *big.Rat) bool
```
EqualBigRat returns true if both *big.Rats are equal


## func Max
``` go
func Max(a, b int) int
```
Max returns the larger of two ints.


## func MaxBigInt
``` go
func MaxBigInt(a, b *big.Int) *big.Int
```
MaxBigInt returns the larger of the two *big.Ints


## func MaxBigRat
``` go
func MaxBigRat(a, b *big.Rat) *big.Rat
```
MaxBigRat returns the larger of the two *big.Rats


## func MaxInt
``` go
func MaxInt(a, b int) int
```
MaxInt returns the larger of two ints.


## func MaxInt16
``` go
func MaxInt16(a, b int16) int16
```
MaxInt16 returns the larger of two int16s.


## func MaxInt16N
``` go
func MaxInt16N(v ...int16) int16
```
MaxInt16N returns the largest int16 in the set provided.
If no values are provided, MaxInt16 returns 0.


## func MaxInt32
``` go
func MaxInt32(a, b int32) int32
```
MaxInt32 returns the larger of two int32s.


## func MaxInt32N
``` go
func MaxInt32N(v ...int32) int32
```
MaxInt32N returns the largest int32 in the set provided.
If no values are provided, MaxInt32 returns 0.


## func MaxInt64
``` go
func MaxInt64(a, b int64) int64
```
MaxInt64 returns the larger of two int64s.


## func MaxInt64N
``` go
func MaxInt64N(v ...int64) int64
```
MaxInt64N returns the largest int64 in the set provided.
If no values are provided, MaxInt64 returns 0.


## func MaxInt8
``` go
func MaxInt8(a, b int8) int8
```
MaxInt8 returns the larger of two int8s.


## func MaxInt8N
``` go
func MaxInt8N(v ...int8) int8
```
MaxInt8N returns the largest int8 in the set provided.
If no values are provided, MaxInt8 returns 0.


## func MaxIntN
``` go
func MaxIntN(v ...int) int
```
MaxIntN returns the largest int in the set provided.
If no values are provided, MaxInt returns 0.


## func MaxN
``` go
func MaxN(v ...int) int
```
MaxN returns the largest int in the set provided.
If no values are provided, Max returns 0.


## func MaxUint
``` go
func MaxUint(a, b uint) uint
```
MaxUint returns the larger of two uints.


## func MaxUint16
``` go
func MaxUint16(a, b uint16) uint16
```
MaxUint16 returns the larger of two uint16s.


## func MaxUint16N
``` go
func MaxUint16N(v ...uint16) uint16
```
MaxUint16N returns the largest uint16 in the set provided.
If no values are provided, MaxUint16 returns 0.


## func MaxUint32
``` go
func MaxUint32(a, b uint32) uint32
```
MaxUint32 returns the larger of two uint32s.


## func MaxUint32N
``` go
func MaxUint32N(v ...uint32) uint32
```
MaxUint32N returns the largest uint32 in the set provided.
If no values are provided, MaxUint32 returns 0.


## func MaxUint64
``` go
func MaxUint64(a, b uint64) uint64
```
MaxUint64 returns the larger of two uint64s.


## func MaxUint64N
``` go
func MaxUint64N(v ...uint64) uint64
```
MaxUint64N returns the largest uint64 in the set provided.
If no values are provided, MaxUint64 returns 0.


## func MaxUint8
``` go
func MaxUint8(a, b uint8) uint8
```
MaxUint8 returns the larger of two uint8s.


## func MaxUint8N
``` go
func MaxUint8N(v ...uint8) uint8
```
MaxUint8N returns the largest uint8 in the set provided.
If no values are provided, MaxUint8 returns 0.


## func MaxUintN
``` go
func MaxUintN(v ...uint) uint
```
MaxUintN returns the largest uint in the set provided.
If no values are provided, MaxUint returns 0.


## func Min
``` go
func Min(a, b int) int
```
Min returns the smaller of two ints.


## func MinBigInt
``` go
func MinBigInt(a, b *big.Int) *big.Int
```
MinBigInt returns the smaller of the two *big.Ints


## func MinBigRat
``` go
func MinBigRat(a, b *big.Rat) *big.Rat
```
MinBigRat returns the smaller of the two *big.Rats


## func MinInt
``` go
func MinInt(a, b int) int
```
MinInt returns the smaller of two ints.


## func MinInt16
``` go
func MinInt16(a, b int16) int16
```
MinInt16 returns the smaller of two int16s.


## func MinInt16N
``` go
func MinInt16N(v ...int16) int16
```
MinInt16N returns the smallest int16 in the set provided.
If no values are provided, MinInt16 returns 0.


## func MinInt32
``` go
func MinInt32(a, b int32) int32
```
MinInt32 returns the smaller of two int32s.


## func MinInt32N
``` go
func MinInt32N(v ...int32) int32
```
MinInt32N returns the smallest int32 in the set provided.
If no values are provided, MinInt32 returns 0.


## func MinInt64
``` go
func MinInt64(a, b int64) int64
```
MinInt64 returns the smaller of two int64s.


## func MinInt64N
``` go
func MinInt64N(v ...int64) int64
```
MinInt64N returns the smallest int64 in the set provided.
If no values are provided, MinInt64 returns 0.


## func MinInt8
``` go
func MinInt8(a, b int8) int8
```
MinInt8 returns the smaller of two int8s.


## func MinInt8N
``` go
func MinInt8N(v ...int8) int8
```
MinInt8N returns the smallest int8 in the set provided.
If no values are provided, MinInt8 returns 0.


## func MinIntN
``` go
func MinIntN(v ...int) int
```
MinIntN returns the smallest int in the set provided.
If no values are provided, MinInt returns 0.


## func MinN
``` go
func MinN(v ...int) int
```
MinN returns the smallest int in the set provided.
If no values are provided, Min returns 0.


## func MinUint
``` go
func MinUint(a, b uint) uint
```
MinUint returns the smaller of two uints.


## func MinUint16
``` go
func MinUint16(a, b uint16) uint16
```
MinUint16 returns the smaller of two uint16s.


## func MinUint16N
``` go
func MinUint16N(v ...uint16) uint16
```
MinUint16N returns the smallest uint16 in the set provided.
If no values are provided, MinUint16 returns 0.


## func MinUint32
``` go
func MinUint32(a, b uint32) uint32
```
MinUint32 returns the smaller of two uint32s.


## func MinUint32N
``` go
func MinUint32N(v ...uint32) uint32
```
MinUint32N returns the smallest uint32 in the set provided.
If no values are provided, MinUint32 returns 0.


## func MinUint64
``` go
func MinUint64(a, b uint64) uint64
```
MinUint64 returns the smaller of two uint64s.


## func MinUint64N
``` go
func MinUint64N(v ...uint64) uint64
```
MinUint64N returns the smallest uint64 in the set provided.
If no values are provided, MinUint64 returns 0.


## func MinUint8
``` go
func MinUint8(a, b uint8) uint8
```
MinUint8 returns the smaller of two uint8s.


## func MinUint8N
``` go
func MinUint8N(v ...uint8) uint8
```
MinUint8N returns the smallest uint8 in the set provided.
If no values are provided, MinUint8 returns 0.


## func MinUintN
``` go
func MinUintN(v ...uint) uint
```
MinUintN returns the smallest uint in the set provided.
If no values are provided, MinUint returns 0.









- - -
Generated by [godoc2md](http://godoc.org/github.com/davecheney/godoc2md)
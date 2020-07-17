// Copyright 2020 morgine.com. All rights reserved.

package convert

// Integer limit values.
const (
	MaxInt8   = 1<<7 - 1
	MinInt8   = -1 << 7
	MaxInt16  = 1<<15 - 1
	MinInt16  = -1 << 15
	MaxInt32  = 1<<31 - 1
	MinInt32  = -1 << 31
	MaxInt64  = 1<<63 - 1
	MinInt64  = -1 << 63
	MaxUint8  = 1<<8 - 1
	MaxUint16 = 1<<16 - 1
	MaxUint32 = 1<<32 - 1
	MaxUint64 = 1<<64 - 1

	// Base on system bits
	// 32-bit system uint equals uint32, int equals int32
	// 64-bit system uint equals uint64, int equals int64
	MinUint uint = 0
	MaxUint      = ^MinUint
	MaxInt       = int(MaxUint >> 1)
	MinInt       = ^MaxInt
)

# Convert

Golang basic types conversion

# Value Interface

```go
type Value interface {
	Int() (int, error)
	Int8() (int8, error)
	Int16() (int16, error)
	Int32() (int32, error)
	Int64() (int64, error)
	Uint() (uint, error)
	Uint8() (uint8, error)
	Uint16() (uint16, error)
	Uint32() (uint32, error)
	Uint64() (uint64, error)
	Bool() (bool, error)
	Float32() (float32, error)
	Float64() (float64, error)
	String() string
}
```

## Implemented Types
 * `Bool`
 * `Int`, `Int8`, `Int16`, `Int32`, `Int64`
 * `Uint`, `Uint8`, `Uint16`, `Uint32`, `Uint64`
 * `Float32`, `Float64`
 * `String`

# Usage:

See `example_test.go`

# Special Conversion

## Numbers - Boolean:
 * `0 <=> false`
 * `1 <=> true`
 
## String - Boolean: 
 * `""/"0"/"f"/"F"/"false"/"FALSE"/"False" => false`
 * `"1"/"t"/"T"/"true"/"TRUE"/"True" => true`
 * `false => "0"`
 * `true => "1"`
 
## String - Numbers:
 * "" => 0 

# Note

Be careful with converting big floating numbers:
```go
package main

import (
    "fmt"
)

func main() {
	var (
		f1, f2 float64
	)
	f1 = 999999999999999999
	f2 = f1 + 10
	fmt.Println(f1 == f2)

	// Output:
	// true
}
```

# Licence

[MIT](https://mit-license.org) Licence.
# Convert

Golang basic types conversion

Example:
```go
package main

import (
    "fmt"
    "github.com/morgine/convert"
)

func main() {
	var (
		any convert.Value // all types implemented the interface
		i32 int32
		i64 int64
		str string
		err error
	)

	any = convert.MustGetValue(1987) // automatic conversion type

	i32, _ = any.Int32()

	str = convert.Int32(i32).String()

	i64, _ = convert.String(str).Int64()

	_, err = convert.Int64(i64).Int8()

	fmt.Println(err == convert.ErrOutOfInt8Range)

	// Output:
	// true
}
```

## Special Conversion

### Numbers - Boolean:
 * `0 <=> false`
 * `1 <=> true`
 
### String - Boolean: 
 * `""/"0"/"f"/"F"/"false"/"FALSE"/"False" => false`
 * `"1"/"t"/"T"/"true"/"TRUE"/"True" => true`
 * `false => "0"`
 * `true => "1"`
 
### String - Numbers:
 * "" => 0 

## Note

Be careful with converting floating numbers:
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
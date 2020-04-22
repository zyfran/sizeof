#### Description

#### Requirements
- go version `1.2.0+`

#### Installation
```bash
$ go get gitlab.com/zyfran/sizeof
```

#### Usage
```go
// main.go
package main

type (
	myStruct struct {
		myBool  bool    // 1 byte
		myFloat float64 // 8 bytes
		myInt   int32   // 4 bytes
		Int     int16   // 2 bytes
	}

	myStructOptimized1 struct {
		myFloat float64 // 8 bytes
		myInt   int32   // 4 bytes
		Int     int16   // 2 bytes
		myBool  bool    // 1 byte
	}

	myStructOptimized2 struct {
		myFloat float64 // 8 bytes
		myInt   int32   // 4 bytes
		myBool  bool    // 1 byte
		Int     int16   // 2 bytes
	}
)

func main() {
	// do something
}
```

```go
// main_test.go
package main

import (
	"testing"

	"gitlab.com/zyfran/sizeof"
)

func TestStructures(t *testing.T) {
	items := [...]interface{}{
		myStruct{},
		myStructOptimized1{},
		myStructOptimized2{},
	}

	for _, item := range items {
		if currentSize, optimalSize := sizeof.CheckStruct(item); currentSize != optimalSize {
			t.Errorf(
				"%sStructure can be optimized from %d to %d bytes",
				sizeof.VisualizeStruct(item),
				currentSize,
				optimalSize,
			)
		}
	}
}
```

##### Check structures for 64-bit OS
```bash
$ GOOS=linux GOARCH=amd64 go test -v
```
```text
--- FAIL: TestStructures (0.00s)
    main_test.go:18: 
        sizeof(main.myStruct)=24 with alignment=8
            myBool  bool    [x][ ][ ][ ][ ][ ][ ][ ]
            myFloat float64 [x][x][x][x][x][x][x][x]
            myInt   int32   [x][x][x][x]
            Int     int16               [x][x][ ][ ]
        Structure can be optimized from 24 to 16 bytes
```
##### Check structures for 32-bit OS
```bash
$ GOOS=linux GOARCH=386 go test -v
```
```text
=== RUN   TestStructures
--- FAIL: TestStructures (0.00s)
    main_test.go:18: 
        sizeof(main.myStruct)=20 with alignment=4
            myBool  bool    [x][ ][ ][ ]
            myFloat float64 [x][x][x][x]
                            [x][x][x][x]
            myInt   int32   [x][x][x][x]
            Int     int16   [x][x][ ][ ]
        Structure can be optimized from 20 to 16 bytes
```

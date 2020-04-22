package sizeof

import (
	"testing"
)

type (
	testGood struct {
		ID    string
		Val   int32
		Val1  int32
		Val2  int16
		test1 uint16
		test2 uint16
		test  bool
	}

	testBad struct {
		test  bool
		test1 uint16
		test2 uint16
		Val1  int32
		Val2  int16
		ID    string
		Val   int32
	}

	test1 struct {
		testGood
		d *int
		f bool
	}

	test2 struct {
		*testGood
		f int
		a interface{}
	}

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

var result interface{}

func BenchmarkVisualizeStruct(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = VisualizeStruct(testGood{})
	}

	result = s
}

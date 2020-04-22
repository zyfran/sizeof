package sizeof

import (
	"fmt"
)

func ExampleVisualizeStruct() {
	items := [...]interface{}{
		nil,
		interface{}(nil),
		5,
		interface{}(5),
		&testGood{},
		testGood{},
		testBad{},
		test1{},
		test2{},
		myStruct{},
		myStructOptimized1{},
		myStructOptimized2{},
	}

	for _, item := range items {
		fmt.Print(VisualizeStruct(item))
	}
	// Output:
	// sizeof(sizeof.testGood)=32 with alignment=8
	//     ID    string [x][x][x][x][x][x][x][x]
	//                  [x][x][x][x][x][x][x][x]
	//     Val   int32  [x][x][x][x]
	//     Val1  int32              [x][x][x][x]
	//     Val2  int16  [x][x]
	//     test1 uint16       [x][x]
	//     test2 uint16             [x][x]
	//     test  bool                     [x][ ]
	// sizeof(sizeof.testBad)=40 with alignment=8
	//     test  bool   [x][ ]
	//     test1 uint16       [x][x]
	//     test2 uint16             [x][x][ ][ ]
	//     Val1  int32  [x][x][x][x]
	//     Val2  int16              [x][x][ ][ ]
	//     ID    string [x][x][x][x][x][x][x][x]
	//                  [x][x][x][x][x][x][x][x]
	//     Val   int32  [x][x][x][x][ ][ ][ ][ ]
	// sizeof(sizeof.test1)=48 with alignment=8
	//    ~testGood sizeof.testGood [x][x][x][x][x][x][x][x]
	//                              [x][x][x][x][x][x][x][x]
	//                              [x][x][x][x][x][x][x][x]
	//                              [x][x][x][x][x][x][x][x]
	//     d        *int            [x][x][x][x][x][x][x][x]
	//     f        bool            [x][ ][ ][ ][ ][ ][ ][ ]
	// sizeof(sizeof.test2)=32 with alignment=8
	//    ~testGood *sizeof.testGood [x][x][x][x][x][x][x][x]
	//     f        int              [x][x][x][x][x][x][x][x]
	//     a        interface {}     [x][x][x][x][x][x][x][x]
	//                               [x][x][x][x][x][x][x][x]
	// sizeof(sizeof.myStruct)=24 with alignment=8
	//     myBool  bool    [x][ ][ ][ ][ ][ ][ ][ ]
	//     myFloat float64 [x][x][x][x][x][x][x][x]
	//     myInt   int32   [x][x][x][x]
	//     Int     int16               [x][x][ ][ ]
	// sizeof(sizeof.myStructOptimized1)=16 with alignment=8
	//     myFloat float64 [x][x][x][x][x][x][x][x]
	//     myInt   int32   [x][x][x][x]
	//     Int     int16               [x][x]
	//     myBool  bool                      [x][ ]
	// sizeof(sizeof.myStructOptimized2)=16 with alignment=8
	//     myFloat float64 [x][x][x][x][x][x][x][x]
	//     myInt   int32   [x][x][x][x]
	//     myBool  bool                [x][ ]
	//     Int     int16                     [x][x]
}

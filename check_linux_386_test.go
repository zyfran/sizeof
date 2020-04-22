package sizeof

import (
	"testing"
)

func TestCheckStruct(t *testing.T) {
	items := [...]struct {
		Interface     interface{}
		InterfaceSize int
		OptimalSize   int
	}{
		{Interface: nil, InterfaceSize: 0, OptimalSize: 0},
		{Interface: true, InterfaceSize: 0, OptimalSize: 0},
		{Interface: test1{}, InterfaceSize: 32, OptimalSize: 32},
		{Interface: test2{}, InterfaceSize: 16, OptimalSize: 16},
		{Interface: testBad{}, InterfaceSize: 28, OptimalSize: 24},
		{Interface: testGood{}, InterfaceSize: 24, OptimalSize: 24},
		{Interface: myStruct{}, InterfaceSize: 20, OptimalSize: 16},
		{Interface: myStructOptimized1{}, InterfaceSize: 16, OptimalSize: 16},
		{Interface: myStructOptimized2{}, InterfaceSize: 16, OptimalSize: 16},
		{Interface: struct {
			a uint8
			b bool
		}{}, InterfaceSize: 2, OptimalSize: 2},
		{Interface: struct {
			a string
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
				f bool
				g bool
			}
			c int16
		}{}, InterfaceSize: 20, OptimalSize: 20},
		{Interface: struct {
			a string
			b struct {
				a uint8
				b bool
				c bool
				d bool
				e bool
			}
			c int16
		}{}, InterfaceSize: 16, OptimalSize: 16},
	}

	for _, item := range items {
		current, best := CheckStruct(item.Interface)
		if item.InterfaceSize != current || item.OptimalSize != best {
			t.Errorf(
				"Interface %v Expected: %d, %d Actual: %d, %d",
				item.Interface,
				item.InterfaceSize, item.OptimalSize,
				current, best,
			)
		}
	}
}

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
		{Interface: test1{}, InterfaceSize: 48, OptimalSize: 48},
		{Interface: test2{}, InterfaceSize: 32, OptimalSize: 32},
		{Interface: testBad{}, InterfaceSize: 40, OptimalSize: 32},
		{Interface: testGood{}, InterfaceSize: 32, OptimalSize: 32},
		{Interface: myStruct{}, InterfaceSize: 24, OptimalSize: 16},
		{Interface: myStructOptimized1{}, InterfaceSize: 16, OptimalSize: 16},
		{Interface: myStructOptimized2{}, InterfaceSize: 16, OptimalSize: 16},
		{Interface: struct {
			a uint8
			b bool
		}{}, InterfaceSize: 2, OptimalSize: 2},
		{Interface: struct {
			a bool
			b float64
			c int32
		}{}, InterfaceSize: 24, OptimalSize: 16},
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
		}{}, InterfaceSize: 32, OptimalSize: 32},
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
		}{}, InterfaceSize: 24, OptimalSize: 24},
	}

	for _, item := range items {
		current, best := CheckStruct(item.Interface)
		if current != item.InterfaceSize || best != item.OptimalSize {
			t.Errorf(
				"Interface %v Expected: %d, %d Actual: %d, %d",
				item.Interface,
				item.InterfaceSize, item.OptimalSize,
				current, best,
			)
		}
	}
}

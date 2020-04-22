package sizeof

import (
	"reflect"
	"strconv"
	"strings"
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

// VisualizeStruct returns string representation of item
// if item is not struct returns empty string
func VisualizeStruct1(item interface{}) string {
	if item == nil {
		return ""
	}

	t := reflect.TypeOf(item)
	if t.Kind() != reflect.Struct {
		return ""
	}

	b := strings.Builder{}
	numFields := t.NumField()
	rowAmount := numFields
	var nameLength, typeLength int
	fields := make([]reflect.StructField, numFields)

	for i := 0; i < numFields; i++ {
		field := t.Field(i)
		fields[i] = field

		lengthName := len(field.Name)
		if lengthName > nameLength {
			nameLength = lengthName
		}

		lengthType := len(field.Type.String())
		if lengthType > typeLength {
			typeLength = lengthType
		}

		if rows := int(field.Type.Size()) / t.Align(); rows > 1 {
			rowAmount += rows - 1
		}
	}

	itemSize := int(t.Size())
	itemSizeString := strconv.Itoa(itemSize)
	itemAlignString := strconv.Itoa(t.Align())
	b.Grow(26 + len(t.String()) + len(itemSizeString) + len(itemAlignString) + // first row
		((7 + nameLength + typeLength + (t.Align() * 3)) * rowAmount))

	b.WriteString("sizeof(")
	b.WriteString(t.String())
	b.WriteString(")=")
	b.WriteString(itemSizeString)

	b.WriteString(" with alignment=")
	b.WriteString(itemAlignString)

	row := uintptr(0)

	for i := 0; i < numFields; i++ {
		field := fields[i]

		if row < field.Offset {
			for ; row < field.Offset; row++ {
				b.WriteString(empty)
			}
		}

		b.WriteByte('\n')

		b.WriteString("   ")
		if field.Anonymous {
			b.WriteByte('~')
		} else {
			b.WriteByte(' ')
		}
		b.WriteString(field.Name)

		lengthName := len(field.Name)
		for ; lengthName <= nameLength; lengthName++ {
			b.WriteByte(' ')
		}

		b.WriteString(field.Type.String())
		lengthType := len(field.Type.String())
		for ; lengthType <= typeLength; lengthType++ {
			b.WriteByte(' ')
		}

		if int(field.Offset)%t.Align() != 0 {
			for i := int(row); i%t.Align() != 0; i-- {
				b.WriteString(skip)
			}
		}

		for i := field.Offset; i < field.Offset+field.Type.Size(); i++ {
			if i != field.Offset && int(i)%t.Align() == 0 {
				b.WriteByte('\n')
				b.WriteString("    ")
				for i := 0; i <= nameLength; i++ {
					b.WriteByte(' ')
				}
				for i := 0; i <= typeLength; i++ {
					b.WriteByte(' ')
				}
			}
			row++
			b.WriteString(used)
		}
	}

	if row < t.Size() {
		for ; row < t.Size(); row++ {
			b.WriteString(empty)
		}
	}

	b.WriteByte('\n')

	return b.String()
}

var result interface{}

func BenchmarkVisualizeStruct1(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = VisualizeStruct1(testGood{})
	}

	result = s
}

func BenchmarkVisualizeStruct(b *testing.B) {
	var s string

	for i := 0; i < b.N; i++ {
		s = VisualizeStruct(testGood{})
	}

	result = s
}

package sizeof

import (
	"reflect"
)

// CheckStruct returns current and best struct size
// if item is not struct returns zero values
func CheckStruct(item interface{}) (currentSize, bestSize int) {
	if item == nil {
		return
	}

	t := reflect.TypeOf(item)
	if t.Kind() != reflect.Struct {
		return
	}

	currentSize = int(t.Size())
	numFields := t.NumField()
	for i := 0; i < numFields; i++ {
		bestSize += int(t.Field(i).Type.Size())
	}

	usedBytes := bestSize % t.Align()
	if usedBytes != 0 {
		bestSize += t.Align() - usedBytes
	}

	return currentSize, bestSize
}

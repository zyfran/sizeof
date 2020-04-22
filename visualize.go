package sizeof

import (
	"strconv"
	"strings"
)

const (
	skip  = "   "
	used  = "[x]"
	empty = "[ ]"
)

// VisualizeStruct returns string representation of item
// if item is not struct returns empty string
func VisualizeStruct(item interface{}) string {
	node := NewNode(item)
	if node.IsZero() {
		return ""
	}

	return visualize(node)
}

func visualize(node Node) string {
	b := strings.Builder{}

	var nameLength, typeLength int
	rowAmount := len(node.Fields)
	for _, field := range node.Fields {
		lengthName := len(field.Name)
		if lengthName > nameLength {
			nameLength = lengthName
		}

		lengthType := len(field.Type)
		if lengthType > typeLength {
			typeLength = lengthType
		}

		if rows := field.Size / node.Alignment; rows > 1 {
			rowAmount += rows - 1
		}
	}

	itemSizeString := strconv.Itoa(node.Size)
	itemAlignString := strconv.Itoa(node.Alignment)
	b.Grow(26 + len(node.Type) + len(itemSizeString) + len(itemAlignString) + // first row
		((7 + nameLength + typeLength + (node.Alignment * 3)) * rowAmount))

	b.WriteString("sizeof(")
	b.WriteString(node.Type)
	b.WriteString(")=")
	b.WriteString(itemSizeString)
	b.WriteString(" with alignment=")
	b.WriteString(itemAlignString)

	var row int
	for _, field := range node.Fields {
		for ; row < field.Offset; row++ {
			b.WriteString(empty)
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

		b.WriteString(field.Type)
		lengthType := len(field.Type)
		for ; lengthType <= typeLength; lengthType++ {
			b.WriteByte(' ')
		}

		if field.Offset%node.Alignment != 0 {
			for i := row; i%node.Alignment != 0; i-- {
				b.WriteString(skip)
			}
		}

		for i := field.Offset; i < field.Offset+field.Size; i++ {
			if i != field.Offset && i%node.Alignment == 0 {
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

	for ; row < node.Size; row++ {
		b.WriteString(empty)
	}

	b.WriteByte('\n')

	return b.String()
}

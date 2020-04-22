// +build !go1.10

package sizeof

import (
	"strconv"
)

func visualize(node Node) string {
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
	b := make([]byte, 0, 26+len(node.Type)+len(itemSizeString)+len(itemAlignString)+ // first row
		((7+nameLength+typeLength+(node.Alignment*3))*rowAmount))

	b = append(b, "sizeof("...)
	b = append(b, node.Type...)
	b = append(b, ")="...)
	b = append(b, itemSizeString...)
	b = append(b, " with alignment="...)
	b = append(b, itemAlignString...)

	var row int
	for _, field := range node.Fields {
		for ; row < field.Offset; row++ {
			b = append(b, empty...)
		}

		b = append(b, '\n')

		b = append(b, "   "...)
		if field.Anonymous {
			b = append(b, '~')
		} else {
			b = append(b, ' ')
		}
		b = append(b, field.Name...)

		lengthName := len(field.Name)
		for ; lengthName <= nameLength; lengthName++ {
			b = append(b, ' ')
		}

		b = append(b, field.Type...)
		lengthType := len(field.Type)
		for ; lengthType <= typeLength; lengthType++ {
			b = append(b, ' ')
		}

		if field.Offset%node.Alignment != 0 {
			for i := row; i%node.Alignment != 0; i-- {
				b = append(b, skip...)
			}
		}

		for i := field.Offset; i < field.Offset+field.Size; i++ {
			if i != field.Offset && i%node.Alignment == 0 {
				b = append(b, '\n')
				b = append(b, "    "...)
				for i := 0; i <= nameLength; i++ {
					b = append(b, ' ')
				}
				for i := 0; i <= typeLength; i++ {
					b = append(b, ' ')
				}
			}
			row++
			b = append(b, used...)
		}
	}

	for ; row < node.Size; row++ {
		b = append(b, empty...)
	}

	b = append(b, '\n')

	return string(b)
}

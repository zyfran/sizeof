package sizeof

import (
	"reflect"
)

type Field struct {
	Name      string
	Type      string
	Offset    int
	Size      int
	Anonymous bool
}

type Node struct {
	Type      string
	Alignment int
	Size      int
	Fields    []Field
}

func (n *Node) IsZero() bool {
	return n.Type == ""
}

func NewNode(item interface{}) Node {
	if item == nil {
		return Node{}
	}

	t := reflect.TypeOf(item)

	if t.Kind() != reflect.Struct {
		return Node{}
	}

	numFields := t.NumField()
	fields := make([]Field, numFields)
	for i := 0; i < numFields; i++ {
		field := t.Field(i)

		fields[i] = Field{
			Name:      field.Name,
			Type:      field.Type.String(),
			Offset:    int(field.Offset),
			Size:      int(field.Type.Size()),
			Anonymous: field.Anonymous,
		}
	}

	return Node{
		Type:      t.String(),
		Alignment: t.Align(),
		Size:      int(t.Size()),
		Fields:    fields,
	}
}

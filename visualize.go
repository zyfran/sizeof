package sizeof

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

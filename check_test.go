package sizeof

import (
	"testing"
)

func BenchmarkCheckStruct(b *testing.B) {
	var current, best int

	for i := 0; i < b.N; i++ {
		current, best = CheckStruct(testGood{})
	}

	result = current
	result = best
}

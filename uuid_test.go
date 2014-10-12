package chance

import (
	"testing"
)

func BenchmarkUUID(b *testing.B) {

	for i := 0; i < b.N; i++ {
		UUIDString()
	}
}

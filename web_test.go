package chance

import (
	"fmt"
	"testing"
)

func TestEmail(t *testing.T) {

	fmt.Printf("Email %s \n", Email())
}

func BenchmarkEmail(b *testing.B) {

	for i := 0; i < b.N; i++ {
		Email()
	}

}

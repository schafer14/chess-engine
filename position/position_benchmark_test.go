package position_test

import (
	"testing"

	"github.com/schafer14/chess/position"
)

func BenchmarkPosition(b *testing.B) {
	var p position.Position

	for i := 0; i < b.N; i++ {
		p = position.New()
	}

	_ = p
}

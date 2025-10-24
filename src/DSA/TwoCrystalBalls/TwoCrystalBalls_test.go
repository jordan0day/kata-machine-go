package twocrystalballs

import (
	"fmt"
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoCrystalBalls(t *testing.T) {
	data := make([]bool, 10_000)
	nobreaks := make([]bool, 10_000)
	idx := rand.Intn(len(data))

	for i := idx; i < len(data); i++ {
		data[i] = true
	}

	assert.Equal(t, idx, TwoCrystalBalls(data), fmt.Sprintf("Expected return value to equal %d", idx))
	assert.Equal(t, -1, TwoCrystalBalls(nobreaks), "Expected return value to equal -1")
}

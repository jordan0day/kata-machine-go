package bfsgraphmatrix

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestBFSGraphMatrix(t *testing.T) {
	require.Equal(t, []int{0, 1, 4, 5, 6}, Bfs(dsa.Matrix2(), 0, 0))

	require.Nil(t, Bfs(dsa.Matrix2(), 6, 0))
}

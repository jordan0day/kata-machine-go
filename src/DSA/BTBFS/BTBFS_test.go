package btbfs

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestBTBFS(t *testing.T) {
	require.True(t, bfs(dsa.Tree, 45))
	require.True(t, bfs(dsa.Tree, 7))
	require.False(t, bfs(dsa.Tree, 69))
}

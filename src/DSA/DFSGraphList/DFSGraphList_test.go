package dfsgraphlist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestDFSGraphList(t *testing.T) {
	require.Equal(t, []int{0, 1, 4, 5, 6}, Dfs(dsa.List2(), 0, 6))
	require.Nil(t, Dfs(dsa.List2(), 6, 0))
}

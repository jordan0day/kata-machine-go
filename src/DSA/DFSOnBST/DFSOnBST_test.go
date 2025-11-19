package dfsonbst

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestDFSonBST(t *testing.T) {
	require.True(t, dfs(&dsa.Tree, 45))
	require.True(t, dfs(&dsa.Tree, 7))
	require.False(t, dfs(&dsa.Tree, 69))
}

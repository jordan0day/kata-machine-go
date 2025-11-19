package comparebinarytrees

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestCompareBinaryTrees(t *testing.T) {
	require.True(t, Compare(&dsa.Tree, &dsa.Tree))
	require.False(t, Compare(&dsa.Tree, &dsa.Tree2))
}

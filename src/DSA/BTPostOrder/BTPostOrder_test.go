package btpostorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestPostOrderSearch(t *testing.T) {
	expected := []int{7, 5, 15, 10, 29, 45, 30, 100, 50, 20}
	require.Equal(t, expected, PostOrderSearch(dsa.Tree))
}

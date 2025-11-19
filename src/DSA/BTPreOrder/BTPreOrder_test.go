package btpreorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestPreOrderSearch(t *testing.T) {
	expected := []int{20, 10, 5, 7, 15, 50, 30, 29, 45, 100}
	require.Equal(t, expected, PreOrderSearch[int](dsa.Tree))
}

package btinorder

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestInOrderSearch(t *testing.T) {
	expected := []int{5, 7, 10, 15, 20, 29, 30, 45, 50, 100}
	require.Equal(t, expected, InOrderSearch[int](dsa.Tree))
}

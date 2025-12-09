package dijkstralist

import (
	"testing"

	dsa "github.com/nacknime-official/kata-machine-go/src/DSA"
	"github.com/stretchr/testify/require"
)

func TestDijkstraList(t *testing.T) {
	require.Equal(t, []int{0, 1, 4, 5, 6}, DijkstraList(dsa.List1(), 0, 6))
}

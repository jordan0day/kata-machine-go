package insertionsort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestInsertionSort(t *testing.T) {
	require.Equal(t, []int{3, 4, 7, 9, 42, 69, 420}, Sort([]int{9, 3, 7, 4, 69, 420, 42}))
}

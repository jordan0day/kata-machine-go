package selectionsort

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSelectionSort(t *testing.T) {
	require.Equal(t, []int{1, 2, 3, 4, 5}, Sort([]int{5, 4, 3, 2, 1}))
}

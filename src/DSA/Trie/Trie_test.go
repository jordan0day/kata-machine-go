package trie

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrie(t *testing.T) {
	trie := &Trie{}
	trie.Insert("foo")
	trie.Insert("fool")
	trie.Insert("foolish")
	trie.Insert("bar")

	result := trie.Find("fo")
	require.Contains(t, result, "foo")
	require.Contains(t, result, "fool")
	require.Contains(t, result, "foolish")

	trie.Delete("fool")

	result = trie.Find("fo")
	require.Contains(t, result, "foo")
	require.Contains(t, result, "fool")
}

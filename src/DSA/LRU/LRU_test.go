package lru

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestLRU(t *testing.T) {
	lru := NewLRU[string, int](3)

	_, ok := lru.Get("foo")
	require.False(t, ok)

	lru.Update("foo", 69)
	val, ok := lru.Get("foo")
	require.True(t, ok)
	require.Equal(t, 69, val)

	lru.Update("bar", 420)
	val, ok = lru.Get("bar")
	require.True(t, ok)
	require.Equal(t, 420, val)

	lru.Update("baz", 1337)
	val, ok = lru.Get("baz")
	require.True(t, ok)
	require.Equal(t, 1337, val)

	lru.Update("ball", 69420)
	val, ok = lru.Get("ball")
	require.True(t, ok)
	require.Equal(t, 69420, val)

	_, ok = lru.Get("foo")
	require.False(t, ok)

	val, ok = lru.Get("bar")
	require.True(t, ok)
	require.Equal(t, 420, val)

	lru.Update("foo", 69)

	val, ok = lru.Get("bar")
	require.True(t, ok)
	require.Equal(t, 420, val)

	val, ok = lru.Get("foo")
	require.True(t, ok)
	require.Equal(t, 69, val)

	_, ok = lru.Get("baz")
	require.False(t, ok)
}

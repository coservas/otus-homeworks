package hw04lrucache

import (
	"math/rand"
	"strconv"
	"sync"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCache(t *testing.T) {
	t.Run("empty cache", func(t *testing.T) {
		c := NewCache(10)

		_, ok := c.Get("aaa")
		require.False(t, ok)

		_, ok = c.Get("bbb")
		require.False(t, ok)
	})

	t.Run("simple", func(t *testing.T) {
		c := NewCache(5)

		wasInCache := c.Set("aaa", 100)
		require.False(t, wasInCache)

		wasInCache = c.Set("bbb", 200)
		require.False(t, wasInCache)

		val, ok := c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 100, val)

		val, ok = c.Get("bbb")
		require.True(t, ok)
		require.Equal(t, 200, val)

		wasInCache = c.Set("aaa", 300)
		require.True(t, wasInCache)

		val, ok = c.Get("aaa")
		require.True(t, ok)
		require.Equal(t, 300, val)

		val, ok = c.Get("ccc")
		require.False(t, ok)
		require.Nil(t, val)
	})

	t.Run("setter logic", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 1)
		c.Set("b", 2)
		c.Set("c", 3)

		val, ok := c.Get("a")
		require.False(t, ok)
		require.Nil(t, val)

		val, ok = c.Get("b")
		require.True(t, ok)
		require.Equal(t, 2, val)

		val, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 3, val)
	})

	t.Run("setter with change logic", func(t *testing.T) {
		c := NewCache(3)

		c.Set("a", 1) // a 1
		c.Set("b", 2) // a 1, b 2
		c.Set("c", 3) // a 1, b 2, c 3
		c.Get("a")    // b 2, c 3, a 1
		c.Set("b", 4) // c 3, a 1, b 4
		c.Set("c", 5) // a 1, b 4, c 5
		c.Set("a", 6) // b 4, c 5, a 6
		c.Set("d", 7) // c 5, a 6, d 7
		c.Get("c ")   // a 6, d 7, c 5

		val, ok := c.Get("a")
		require.True(t, ok)
		require.Equal(t, 6, val)

		_, ok = c.Get("b")
		require.False(t, ok)

		val, ok = c.Get("c")
		require.True(t, ok)
		require.Equal(t, 5, val)

		val, ok = c.Get("d")
		require.True(t, ok)
		require.Equal(t, 7, val)
	})

	t.Run("clear check", func(t *testing.T) {
		c := NewCache(2)

		c.Set("a", 1)
		c.Set("b", 2)

		_, ok := c.Get("a")
		require.True(t, ok)

		_, ok = c.Get("b")
		require.True(t, ok)

		c.Clear()

		_, ok = c.Get("a")
		require.False(t, ok)

		_, ok = c.Get("b")
		require.False(t, ok)
	})
}

func TestCacheMultithreading(t *testing.T) {
	t.Skip() // Remove me if task with asterisk completed.

	c := NewCache(10)
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Set(Key(strconv.Itoa(i)), i)
		}
	}()

	go func() {
		defer wg.Done()
		for i := 0; i < 1_000_000; i++ {
			c.Get(Key(strconv.Itoa(rand.Intn(1_000_000))))
		}
	}()

	wg.Wait()
}

package lru

import (
	"testing"

	"github.com/ilius/is/v2"
)

func TestCapacity(t *testing.T) {
	is := is.New(t)
	tests := []struct {
		capacity int
	}{
		{1},
		{10},
		{100},
	}

	New[int, int](WithCapacity(10))
	for _, tc := range tests {
		lru := New[int, int](WithCapacity(tc.capacity))
		for i := 0; i < tc.capacity+1; i++ {
			lru.Set(i, i)
		}

		is.Msg("expected capacity to be full").Equal(tc.capacity, lru.Len())

		_, ok := lru.Get(0)
		is.Msg("expected key to be evicted").False(ok)

		_, ok = lru.Get(1)
		is.Msg("expected key to exist").True(ok)
	}
}

func TestGetMissing(t *testing.T) {
	is := is.New(t)
	lru := New[int, int]()
	_, ok := lru.Get(0)
	is.Msg("expected not ok").False(ok)
}

func TestSetGet(t *testing.T) {
	is := is.New(t)
	lru := New[int, int]()
	value := 100

	lru.Set(1, value)
	value, ok := lru.Get(1)

	is.True(ok)
	is.Equal(value, value)
}

func TestDelete(t *testing.T) {
	is := is.New(t)
	lru := New[int, int]()

	key, value := 1, 100
	lru.Set(key, value)
	is.Equal(lru.Len(), 1)

	ok := lru.Delete(key)
	is.True(ok)
}

func TestDeleteMissing(t *testing.T) {
	is := is.New(t)
	lru := New[int, int]()
	key := 100
	ok := lru.Delete(key)
	is.False(ok)
}

func TestFlush(t *testing.T) {
	is := is.New(t)
	lru := New[int, int]()
	key, value := 1, 100
	lru.Set(key, value)
	is.Equal(lru.Len(), 1)

	lru.Flush()
	is.Equal(lru.Len(), 0)

	_, ok := lru.Get(key)
	is.False(ok)
}

package internal_test

import (
	"testing"

	"github.com/ilius/go-lru/internal"
	"github.com/ilius/is"
)

func TestPushRemove(t *testing.T) {
	is := is.New(t)
	ll := internal.NewList[int]()
	length := 10

	for i := 1; i <= length; i++ {
		ll.PushFront(i)
		is.Equal(ll.Len(), i)
	}

	for i := length; i <= 1; i++ {
		ll.Remove(ll.Back())
		is.Equal(ll.Len(), i)
	}
}

func TestMoveToFront(t *testing.T) {
	is := is.New(t)
	ll := internal.NewList[int]()
	e := ll.PushFront(0)
	ll.PushFront(1)
	is.Equal(e, ll.Back())

	ll.MoveToFront(e)
	is.NotEqual(e, ll.Back())
}

func TestInit(t *testing.T) {
	is := is.New(t)
	ll := internal.NewList[int]()

	ll.PushFront(1)
	is.Equal(ll.Len(), 1)

	ll.Init()
	is.Equal(ll.Len(), 0)
}

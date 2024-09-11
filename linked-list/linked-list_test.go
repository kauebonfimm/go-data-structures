package linkedlist_test

import (
	"testing"

	linkedlist "github.com/kauebonfimm/go-data-structures/linked-list"
	"github.com/stretchr/testify/assert"
)

func TestLinkedList(t *testing.T) {
	assert := assert.New(t)

	lkl := new(linkedlist.LinkedList[string])

	lkl.InsertAtBeginning("1")
	lkl.InsertAtLast("2")

	assert.Equal(2, lkl.Size())
	assert.Equal("1", lkl.Head.Data)
	assert.Equal("2", lkl.Head.Next.Data)

	lkl.InsertAtLast("3")
	assert.Equal("3", lkl.Head.Next.Next.Data)

	lkl.RemoveAtBeginning()
	assert.Equal("2", lkl.Head.Data)
}

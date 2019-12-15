package reverse

import (
	"log"
	"testing"
)

func TestLinkedListNode_Reverse(t *testing.T) {
	header := NewRandomIntList(10)
	t.Log(header)
	_, newHeader := header.Reverse()
	t.Log(newHeader)

	newHeader2 := newHeader.ReverseInsert()
	t.Log(newHeader2)
}
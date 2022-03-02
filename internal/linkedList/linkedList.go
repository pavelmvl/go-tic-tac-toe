package linkedList

import (
	"fmt"
)

type Item struct {
	value string
	next  *Item
}

func NewItem(value string) Item {
	return Item{
		value: value,
		next:  nil,
	}
}

func (i Item) GetValue() string {
	return i.value
}

func (i Item) getNext() *Item {
	return i.next
}

func (i *Item) setNext(item *Item) {
	i.next = item
}

type LinkedList struct {
	first *Item
}

func NewLinkedList() LinkedList {
	return LinkedList{}
}

func (ll *LinkedList) AddLastItems(items ...Item) {
	var last *Item = ll.first
	for last = ll.first; last != nil && last.getNext() != nil; last = last.getNext() {
	}
	for _, item := range items {
		if ll.first == nil {
			ll.first = &item
			last = ll.first
		} else {
			last.setNext(&item)
			last = last.getNext()
		}
	}
}

func (ll *LinkedList) Reverse() {
	var item, nextItem, nextNextItem *Item
	item = ll.first
	if item == nil {
		return
	}
	nextItem = item.getNext()
	if nextItem == nil {
		return
	}
	item.setNext(nil)
	for {
		nextNextItem = nextItem.getNext()
		if nextNextItem != nil {
			nextItem.setNext(item)
			item = nextItem
			nextItem = nextNextItem
		} else {
			nextItem.setNext(item)
			ll.first = nextItem
			break
		}
	}
}

func (ll LinkedList) String() string {
	buf := make([]byte, 0, 1024)
	i := 0
	for last := ll.first; last != nil; last = last.getNext() {
		str := fmt.Sprintf("(%p)idx %d: value=%s, next=%p\n", last, i, last.value, last.next)
		buf = append(buf, []byte(str)...)
		i++
	}
	return string(buf)
}

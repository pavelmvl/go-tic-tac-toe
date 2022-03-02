package doubleLinkedList

import (
	"fmt"
)

type Item struct {
	value    string
	next     *Item
	previous *Item
}

func NewItem(value string) Item {
	return Item{
		value:    value,
		next:     nil,
		previous: nil,
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

func (i Item) getPrevious() *Item {
	return i.previous
}

func (i *Item) setPrevious(item *Item) {
	i.previous = item
}

type DoubleLinkedList struct {
	first *Item
	last  *Item
}

func NewDoubleLinkedList() DoubleLinkedList {
	return DoubleLinkedList{}
}

func (dll *DoubleLinkedList) AddLastItems(items ...Item) {
	for _, item := range items {
		if dll.last == nil {
			dll.last = &item
			dll.first = dll.last
		} else {
			previous := dll.last
			dll.last.setNext(&item)
			dll.last = dll.last.getNext()
			dll.last.setPrevious(previous)
			dll.last.setNext(nil)
		}
	}
}

func (dll *DoubleLinkedList) Reverse() {
	var newFirst *Item = dll.last
	var newLast *Item = dll.first
	item := dll.first
	if item == nil {
		return
	}
	nextItem := item.getNext()
	if nextItem == nil {
		return
	}
	for {
		next := item.getNext()
		previous := item.getPrevious()
		item.setNext(previous)
		item.setPrevious(next)
		if nextItem != nil {
			item = nextItem
			nextItem = nextItem.getNext()
		} else {
			break
		}
	}
	dll.last = newLast
	dll.first = newFirst
}

func (dll DoubleLinkedList) String() string {
	buf := make([]byte, 0, 1024)
	i := 0
	for item := dll.first; item != nil; item = item.getNext() {
		str := fmt.Sprintf("(%p)idx %d: value=%s, previous=%p, next=%p\n", item, i, item.value, item.previous, item.next)
		buf = append(buf, []byte(str)...)
		i++
	}
	return string(buf)
}

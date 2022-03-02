package main

import (
	"fmt"
	"go-tic-tac-toe/internal/doubleLinkedList"
	"go-tic-tac-toe/internal/linkedList"
	"strconv"
)

func main() {
	//
	// lined list
	//
	ll := linkedList.NewLinkedList()
	for i := 0; i < 5; i++ {
		ll.AddLastItems(linkedList.NewItem(strconv.Itoa(i)))
	}
	fmt.Println(ll)
	ll.Reverse()
	fmt.Println(ll)
	fmt.Println("==============================\n")
	//
	// double linked list
	//
	dll := doubleLinkedList.NewDoubleLinkedList()
	for i := 0; i < 5; i++ {
		dll.AddLastItems(doubleLinkedList.NewItem(strconv.Itoa(i)))
	}
	fmt.Println(dll)
	dll.Reverse()
	fmt.Println(dll)
}

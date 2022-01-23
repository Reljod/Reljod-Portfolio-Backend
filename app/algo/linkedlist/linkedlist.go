package linkedlist

import (
	"fmt"
)

func LinkedListRun() {
	fmt.Println("This is a linked list implementation!")

	var ll LinkedList = new(SingleLinkedList)
	var name1 interface{} = "reljod"
	var name2 interface{} = "Abbie"
	var name3 interface{} = "Peniel"

	ll.AddFront(&name1)
	ll.AddFront(&name2)
	ll.AddFront(&name3)

	head := ll.GetHead()

	fmt.Println(*((*head).Data))
	fmt.Println(*((*head).Next.Data))
	fmt.Println(*((*head).Next.Next.Data))

	c := make(chan interface{})
	go ll.GetData(c)

	for data := range c {
		fmt.Println(data)
	}
}

type Node struct {
	Data *interface{}
	Next *Node
}

type LinkedList interface {
	AddFront(data *interface{})
	GetHead() *Node
	GetData(output chan<- interface{})
}

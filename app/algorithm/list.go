package algorithm

import (
	"container/list"
	"fmt"
)

func ListSample() {
	var intList list.List
	intList.PushBack(10)
	intList.PushBack(20)
	intList.PushBack(30)

	for element := intList.Front(); element != nil; element = element.Next() {
		fmt.Println(element.Value.(int))
	}
}

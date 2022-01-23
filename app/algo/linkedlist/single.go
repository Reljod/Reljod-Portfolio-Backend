package linkedlist

type SingleLinkedList struct {
	Head *Node
}

func (sll *SingleLinkedList) GetHead() *Node {
	return (*sll).Head
}

func (sll *SingleLinkedList) getLatestNode() *Node {
	currentNode := (*sll).Head

	for {
		if currentNode.Next == nil {
			break
		}

		currentNode = currentNode.Next
	}

	return currentNode
}

func (sll *SingleLinkedList) AddFront(data *interface{}) {
	var newNode *Node = new(Node)
	(*newNode).Data = data

	if (*sll).Head == nil {
		sll.Head = newNode
		return
	}

	latestNode := (*sll).getLatestNode()
	latestNode.Next = newNode
}

func (sll *SingleLinkedList) GetData(output chan<- interface{}) {
	currentNode := (*sll).Head

	for {
		if currentNode.Next == nil {
			close(output)
			break
		}

		output <- *currentNode.Data
		currentNode = currentNode.Next
	}

}

// func (sll *SingleLinkedList) AddBack(data *interface{}){}

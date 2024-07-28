package main

import (
	goi "github.com/Matej-Chmel/go-interview"
	lk "github.com/Matej-Chmel/go-linked-list"
)

type Node = *lk.SingleLinkNode[int]

func withoutBuffer(head Node) Node {
	current := head

	for current != nil {
		runner := current

		for runner.Next != nil {
			if runner.Next.Val == current.Val {
				// Remove duplicate
				removed := runner.Next
				runner.Next = removed.Next
				removed.Next = nil
			} else {
				runner = runner.Next
			}
		}

		current = current.Next
	}

	return head
}

func withSet(head Node) Node {
	current := head
	var previous Node = nil
	set := make(map[int]struct{})

	for current != nil {
		_, ok := set[current.Val]

		if ok {
			// Remove duplicate
			previous.Next = current.Next
			current.Next = nil
			current = previous.Next
		} else {
			// Unique node
			set[current.Val] = struct{}{}
			previous = current
			current = current.Next
		}
	}

	return head
}

func main() {
	i := goi.NewInterview[Node, Node]()

	i.AddCase(
		lk.CreateSinglyLinkedList(4, 1, 7, 1, 4, 7, 8, 10),
		lk.CreateSinglyLinkedList(4, 1, 7, 8, 10))
	i.AddCase(
		lk.CreateSinglyLinkedList(1, 1, 1, 1, 0),
		lk.CreateSinglyLinkedList(1, 0))
	i.AddCase(
		lk.CreateSinglyLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10),
		lk.CreateSinglyLinkedList(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
	i.AddCase(
		lk.CreateSinglyLinkedList(4, 5, 10, 9, 1, 9, 10, 7, 1),
		lk.CreateSinglyLinkedList(4, 5, 10, 9, 1, 7))
	i.AddCase(
		lk.CreateSinglyLinkedList[int](),
		lk.CreateSinglyLinkedList[int]())

	i.AddSolutions(withSet, withoutBuffer)
	i.Print()
}

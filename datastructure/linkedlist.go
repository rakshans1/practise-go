package main

import "fmt"

type node struct {
	data int
	next *node
}

type linkedList struct {
	head   *node
	length int
}

func (l *linkedList) prepend(n *node) {
	second := l.head
	l.head = n
	l.head.next = second
	l.length++
}

func (l *linkedList) deleteWithValue(value int) {
  if l.length == 0 {
    return
  }

  if l.head.data == value {
    l.head = l.head.next
    l.length--
    return
  }

  node := l.head
  for node.next.data != value {
    if node.next.next == nil {
      return
    }
    node = node.next
  }
  node.next = node.next.next
  l.length--
}

func (l linkedList) printLinkedList() {
  node := l.head
  for l.length != 0 {
    fmt.Printf("%d ", node.data)
    node = node.next
    l.length--
  }
  fmt.Printf("\n")
}

func main() {
	list := linkedList{}
	node1 := &node{data: 1}
	node2 := &node{data: 2}
	node3 := &node{data: 3}
	list.prepend(node3)
	list.prepend(node2)
	list.prepend(node1)
    list.printLinkedList()
    list.deleteWithValue(2)
    list.deleteWithValue(3)
    list.printLinkedList()
}

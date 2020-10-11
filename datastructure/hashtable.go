package main

import "fmt"

const ArraySize = 7

type HashTable struct {
	array [ArraySize]*bucket
}

func (ht *HashTable) Insert(key string) {
	index := hash(key)
	ht.array[index].insert(key)
}

func (ht *HashTable) Search(key string) bool {
	index := hash(key)
	return ht.array[index].search(key)
}

func (ht *HashTable) Delete(key string) {
	index := hash(key)
	ht.array[index].delete(key)
}

type bucket struct {
	head *bucketNode
}

func (b *bucket) insert(key string) {
	if b.search(key) {
		fmt.Println("already exist")
		return
	}
	newNode := &bucketNode{key: key}
	newNode.next = b.head
	b.head = newNode
}

func (b *bucket) search(key string) bool {
	node := b.head
	for node != nil {
		if node.key == key {
			return true
		}
		node = node.next
	}
	return false
}

func (b *bucket) delete(key string) {
	if !b.search(key) {
		fmt.Println("does not exist")
		return
	}

	if b.head.key == key {
		b.head = b.head.next
		return
	}

	prevnode := b.head
	for prevnode.next != nil {
		if prevnode.next.key == key {
			prevnode.next = prevnode.next.next
			return
		}
		prevnode = prevnode.next
	}
	return
}

type bucketNode struct {
	key  string
	next *bucketNode
}

func NewHashTable() *HashTable {
	ht := &HashTable{}
	for i := range ht.array {
		ht.array[i] = &bucket{}
	}
	return ht
}

func hash(key string) int {
	sum := 0
	for _, v := range key {
		sum += int(v)
	}
	return sum % ArraySize
}

func main() {
	ht := NewHashTable()
	ht.Insert("RANDY")
	ht.Insert("RANDY")
	ht.Insert("KENNY")
	fmt.Println(ht.Search("RANDY"))
	ht.Delete("RANDY")
	fmt.Println(ht.Search("RANDY"))
}

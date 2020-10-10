package main

import "fmt"

const ALPHABET_SIZE = 26

type Node struct {
	childrens [ALPHABET_SIZE]*Node
	isEnd     bool
}

type Trie struct {
	root *Node
}

func (t *Trie) Insert(word string) {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			current.childrens[index] = &Node{}
		}
		current = current.childrens[index]
	}
	current.isEnd = true
}

func (t *Trie) Search(word string) bool {
	wordLength := len(word)
	current := t.root
	for i := 0; i < wordLength; i++ {
		index := word[i] - 'a'
		if current.childrens[index] == nil {
			return false
		}
		current = current.childrens[index]
	}
	if current.isEnd {
		return true
	}
	return false
}

func NewTrie() *Trie {
	return &Trie{root: &Node{}}
}

func main() {
	trie := NewTrie()
	words := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses"}

	for i := 0; i < len(words); i++ {
		trie.Insert(words[i])
	}
	wordsToFind := []string{"sam", "john", "tim", "jose", "rose",
		"cat", "dog", "dogg", "roses", "rosess", "ans", "san"}
	for i := 0; i < len(wordsToFind); i++ {
		found := trie.Search(wordsToFind[i])
		if found {
			fmt.Printf("Word \"%s\" found in trie\n", wordsToFind[i])
		} else {
			fmt.Printf("Word \"%s\" not found in trie\n", wordsToFind[i])
		}
	}
}

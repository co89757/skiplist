package skiplist

import (
	"math/rand"
)

const (
	//MaxHeight is the maximum height for the list
	MaxHeight = 10
)

// SkipNode represents element in a skiplist
type SkipNode struct {
	Height  int
	Forward []*SkipNode
	Key     string
	Value   interface{}
}

//SkipList is the skiplist impl
type SkipList struct {
	Head   *SkipNode
	Height int
	Count  int
}

func newNode(key string, val interface{}, height int) *SkipNode {
	node := &SkipNode{
		Height:  height,
		Forward: make([]*SkipNode, MaxHeight+1),
		Key:     key,
		Value:   val,
	}

	return node
}

//New creates a new skiplist instance
func New() (list *SkipList) {
	list = &SkipList{
		Head:   newNode("", nil, 0),
		Height: 0,
		Count:  0,
	}
	return
}

// return the update vector for incoming key, the insert position is the last node
// whose key is less than incoming key or whose next is not nil
func (list *SkipList) updateVector(key string) (update []*SkipNode, keyExist bool) {
	update = make([]*SkipNode, MaxHeight+1)
	cursor := list.Head
	for h := list.Height; h > 0; h-- {
		for ; cursor.Forward[h].Key < key && cursor.Forward[h] != nil; cursor = cursor.Forward[h] {

		}
		update[h] = cursor
		if cursor.Forward[h].Key == key {
			keyExist = true
		}
	}
	return
}

//Add a new key-value pair to the list. if key already exists, return false
func (list *SkipList) Add(key string, val interface{}) (keyExist bool) {
	update, keyExist := list.updateVector(key)
	if keyExist {
		//key already exist
		for h := list.Height; h > 0; h-- {
			update[h].Forward[h].Value = val
		}
		return
	}
	//it is a new key
	newLevel := getRandomHeight()
	newNode := newNode(key, val, newLevel)
	if newLevel > list.Height {
		for h := list.Height + 1; h <= newLevel; h++ {
			update[h] = list.Head
		}
		//update height
		list.Height = newLevel
	}
	for h := newLevel; h > 0; h-- {
		newNode.Forward[h] = update[h].Forward[h]
		update[h].Forward[h] = newNode
	}
	list.Count++
	return
}

func getRandomHeight() (height int) {
	p := 0.5
	height = 1
	for ; rand.Float64() < p && height < MaxHeight; height++ {
	}

	return
}

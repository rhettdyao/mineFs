package skiplist

import "math/rand"

const p = 0.25

const DefaultMaxLevel = 32

type SkipList struct {
	LessThan func(a, b interface{}) bool
	nextNode []int
	prevNode []int
	maxHeight int
	length int
}

func (s *SkipList) randHeight()(h int){
	const branching = 4
	h = 1
	for h < s.DefaultMaxLevel && rand.Int() % branching == 0{
		h ++
	}
	return
}



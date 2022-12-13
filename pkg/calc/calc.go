package calc

import (
	mapSet "github.com/deckarep/golang-set/v2"
	"math/rand"
)

const (
	MaxCallNumber = 99
)

func CreateUserCard(max int) []int8 {
	if max < 0 {
		return make([]int8, 0, 0)
	}
	set := mapSet.NewSet[int8]()
	for i := 0; i < max; i++ {
		set.Add(RandomNumber())
	}
	set = add(set, max)
	array := make([]int8, max, max)
	i := 0
	for val := range set.Iterator().C {
		array[i] = val
		i = i + 1
	}
	return array
}

func RandomNumber() int8 {
	return int8(rand.Intn(MaxCallNumber) + 1)
}

func add(set mapSet.Set[int8], max int) mapSet.Set[int8] {
	if set.Cardinality() >= max {
		return set
	}
	set.Add(RandomNumber())
	add(set, max)
	return set
}

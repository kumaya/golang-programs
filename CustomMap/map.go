// TODO: add delete, add collision and dynamic expansion

package main

import (
	"fmt"
	"hash/fnv"
)

type customMap interface {
	Get(key interface{}) (value interface{}, exists bool)
	Set(key, value interface{}) bool
	Len() int
}

type cMap struct {
	storage  [][]*KeyValue
	capacity uint64
	length   int
}

func getKeyHash(key string) uint64 {
	h := fnv.New64()
	h.Write([]byte(key))
	return h.Sum64()
}

func (c *cMap) Get(key interface{}) (interface{}, bool) {
	keyHash := getKeyHash(stringer(key))
	slot := keyHash % c.capacity
	elements := c.storage[slot]
	for _, elem := range elements {
		if elem.Key == key {
			return elem.Value, true
		}
	}
	return nil, false
}

func (c *cMap) Set(key, value interface{}) bool {
	keyHash := getKeyHash(stringer(key))
	slot := keyHash % c.capacity
	if len(c.storage[slot]) > 0 {
		for _, elem := range c.storage[slot] {
			if elem.Key == key {
				elem.Value = value
				break
			}
		}
	} else {
		c.storage[slot] = append(c.storage[slot], &KeyValue{key, value})
		c.length += 1
	}
	return true
}

func (c *cMap) Len() int {
	return c.length
}

func NewCMap() *cMap {
	capa := uint64(100)
	storage := make([][]*KeyValue, capa)
	return &cMap{
		storage:  storage,
		capacity: capa,
		length:   0,
	}
}

func main() {
	myMap := NewCMap()

	k := "test"
	val, isExists := myMap.Get(k)
	fmt.Printf("%s:%v:%v\n", k, val, isExists)

	myMap.Set(k, "done")
	val, isExists = myMap.Get(k)
	fmt.Printf("%s:%v:%v\n", k, val, isExists)

	myMap.Set(k, "updated")
	val, isExists = myMap.Get(k)
	fmt.Printf("%s:%v:%v\n", k, val, isExists)

	k = "test1"
	myMap.Set(k, "newVal")
	val, isExists = myMap.Get(k)
	fmt.Printf("%s:%v:%v\n", k, val, isExists)

	fmt.Printf("Length of map: %d\n", myMap.Len())

	k1 := 1
	myMap.Set(k1, "int1")
	val, isExists = myMap.Get(k1)
	fmt.Printf("%v:%v:%v\n", k1, val, isExists)

	type Temp struct {
		name string
		age  int
	}
	k2 := Temp{"mayank", 12}
	myMap.Set(k2, "struct")
	val, isExists = myMap.Get(Temp{age: 12, name: "mayank"})
	fmt.Printf("%v:%v:%v\n", k2, val, isExists)
}

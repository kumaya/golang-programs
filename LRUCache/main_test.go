package LRUCache

import (
	"testing"
)

type sampleVals struct {
	name       string
	add        interface{}
	get        interface{}
	expectedOk bool
}

var getTests = []sampleVals{
	{"string_hit", "myKey", "myKey", true},
	{"string_miss", "myKey", "nonsense", false},
	{"string_hit", "nonsense", "nonsense", true},
    {"string_hit", "nonsense1", "nonsense1", true},
    {"string_miss", "nonsense2", "nonsense", false},
	{"string_miss", "nonsense", "myKey", false},
	{"string_hit", "nonsense", "nonsense", true},
}

func TestGet(t *testing.T) {
	var lru *LRUCache
	lru = New(2)
	for _, testVal := range getTests {
		lru.Set(testVal.add, 1234)
		val, exists := lru.Get(testVal.get)
		if exists != testVal.expectedOk {
			t.Fatalf("%s: cache hit = %v; want = %v", testVal.name, exists, !exists)
		} else if exists && val != 1234 {
			t.Fatalf("%s expected get to return 1234 but got %v", testVal.name, val)
		}
	}
}

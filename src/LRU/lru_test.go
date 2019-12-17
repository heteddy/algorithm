package LRU

import (
	"fmt"
	"log"
	"testing"
)

func iniTestCase() *LRUCache {

	cache := NewLRUCache(3)
	cache.Save("web", 100)
	cache.Save("signal", 101)
	cache.Save("test", 102)
	fmt.Println(cache)
	return cache
}

func TestLRUCache_Get(t *testing.T) {
	cache := iniTestCase()
	cache.Get("signal")
	fmt.Println(cache)
	cache.Get("test")
	fmt.Println(cache)
	cache.Get("web")
	fmt.Println(cache)
}

func TestLRUCache_Save(t *testing.T) {
	cache := iniTestCase()
	cache.Save("test2", 112)
	fmt.Println(cache)
	cache.Save("test3", 113)
	fmt.Println(cache)
	cache.Save("test4", 114)
	fmt.Println(cache)
	cache.Save("test5", 115)
	fmt.Println(cache)
	cache.Save("test6", 116)
	fmt.Println(cache)

}

func TestMain(m *testing.M) {
	log.Println("prepare test data")
	m.Run()
	log.Println("clean test data")
}

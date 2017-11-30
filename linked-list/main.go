package main

import (
	"algorithm/linked-list/reverse"
	"log"
)

func main() {
	header := reverse.NewRandomIntList(10)
	log.Println(header)
	_, newHeader := header.Reverse()
	log.Println(newHeader)

	newHeader2 := newHeader.ReverseInsert()
	log.Println(newHeader2)
}

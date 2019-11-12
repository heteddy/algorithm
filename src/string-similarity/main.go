package main

import (
	"algorithm/string-similarity/max-common"
	"log"
)

func main() {
	log.Println(max_common.CommonString("poll", "pull"))
	log.Println(max_common.CommonSequence("poll", "pull"))
	log.Println(max_common.CommonString("hisah", "vista"))
	log.Println(max_common.CommonSequence("hisah", "vista"))
	log.Println(max_common.CommonString("fish", "fosh"))
	log.Println(max_common.CommonSequence("fish", "fosh"))
}

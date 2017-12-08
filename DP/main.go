/*
1.有10级台阶，每次可能上一级或二级。共有多少种上法？
*/
package main

import (
	"algorithm/DP/stairs"
	"log"
	"time"
)

func main() {
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"))
	log.Println(stairs.StairsWithoutCache(40))
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"))
	log.Println(stairs.SequenceStairCounter(40))
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"))
	log.Println(stairs.StairCounter(40))
	log.Println(time.Now().Format("2006-01-02 15:04:05.000000"))
}

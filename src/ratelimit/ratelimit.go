/*
限制请求速率
*/

package ratelimit

import (
	"log"
	"time"
)

const MAX_REQUEST_NUMBER = 200
const PERIOD = time.Second
const TICKER_INTERVAL = PERIOD / 200

type request interface {
}

func handler(r request) {
	log.Println(r.(int))
}

func handleRequests(requests <-chan request, limit <-chan struct{}) {
	for {
		select {
		case r := <-requests:
			// limit 有数据
			<-limit
			go handler(r)
		}
	}

}

func limitRate(limit chan<- struct{}) {
	tick := time.NewTicker(TICKER_INTERVAL)
	defer tick.Stop()
	for {
		select {
		case <-tick.C:
			limit <- struct{}{}
		}
	}
}

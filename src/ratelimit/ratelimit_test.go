package ratelimit

import "testing"

func TestRateLimit(t *testing.T) {
	limitChan := make(chan struct{}, MAX_REQUEST_NUMBER)
	requestChan := make(chan request)
	defer close(limitChan)
	defer close(requestChan)

	go handleRequests(requestChan, limitChan)
	go limitRate(limitChan)

	for i := 0; i < 1000; i++ {
		requestChan <- i
	}
}

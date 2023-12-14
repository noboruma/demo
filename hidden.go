// SKIP THIS FILE
package main

import (
	"math/rand"
	"time"
)

const (
	stream = "123;345;678;901;234;578"
)

var (
	index = 0
)

func getNextChar() byte {
	if index < len(stream) {
		tmp := stream[index]
		index += 1
		return tmp
	}
	return 0
}

func checkWithDB(cve_signature int) bool {
	time.Sleep(5 * time.Second)
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(2) == 1
}

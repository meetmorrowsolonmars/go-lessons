package main

import (
	"errors"
	"math/rand/v2"
	"time"
)

// You need to write a wrapper function that can be used to call the blocking function.
// The makeBlockingSyscall function can block for up to 10 seconds. We want to wait a maximum of 3 seconds.

func main() {
}

func makeBlockingSyscall() (string, error) {
	// Random delay from 0 to 10 seconds.
	time.Sleep(time.Second * rand.N[time.Duration](10))

	// Generates an error 20% of the time.
	if rand.N[int](10) < 2 {
		return "", errors.New("unexpected error")
	}

	return "success", nil
}

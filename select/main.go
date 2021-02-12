package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

var timeout = 10 * time.Second

// Racer a bのうちレスポンスが早い方を返す
func Racer(a, b string) (string, error) {
	return ConfigurableRacer(a, b, timeout)
}

func ConfigurableRacer(a, b string, timeout time.Duration) (string, error) {
	select {
	case <-ping(a):
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout):
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	ch := make(chan struct{})
	go func() {
		if _, err := http.Get(url); err != nil {
			log.Fatal(err)
		}
		close(ch)
	}()
	return ch
}

/*
func measureResponseTime(url string) time.Duration {
	start := time.Now()
	if _, err := http.Get(url); err != nil {
		log.Fatal(err)
	}
	duration := time.Since(start)
	return duration
}
*/

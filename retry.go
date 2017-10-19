package main

import (
	"errors"
	"time"
)

// retry fn function in retryCnt times, in every sleep time duration
func retry(retryCnt int, sleep time.Duration, fn func() error) error {
	if retryCnt < 0 {
		return errors.New("all try failed")
	}

	if err := fn(); err != nil {
		time.Sleep(sleep * time.Second)
		retryCnt--
		return retry(retryCnt, sleep, fn)
	}

	return nil
}

package main

import (
	"sync"
	"time"
)

type TokenBucket struct {
	rate                int64 // represents the rate at which the bucket should be filled.
	maxTokens           int64 // represents the max tokens capacity that the bucket can hold.
	currentTokens       int64 //  tokens currently present in the bucket at any time t.
	lastRefillTimestamp time.Time
	mutex               sync.Mutex
}

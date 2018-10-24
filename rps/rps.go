package rps

// Makes sure you do not exceed predefined rate limit. Useful for API calls.

import (
	"fmt"
	"time"
	"sync"
)

type RPS struct {
	rps float64
	counter int64
	time time.Time
	mu sync.Mutex
	work chan func()
}

func New(rps float64) *RPS {
    return &RPS{
    	rps: rps,
    	time: time.Now(),
    	mu: sync.Mutex{},
        work: make(chan func()),
    }
}

func (r *RPS) Run(task func(), stop <-chan struct{}) (bool) {
    
    elapsed := time.Since(r.time)

	// incrementing counter
	r.mu.Lock()
	current := float64(r.counter)/elapsed.Seconds()
	r.counter += 1
	counter := r.counter
	r.mu.Unlock()

	fmt.Printf("%.2f\n", current)
	if current <= r.rps {
		task()
		return true
	}

	// rps exceeded, sleeping for a while
	delayms := (float64(counter)/r.rps - elapsed.Seconds()) * 1000

	if stop != nil {
		select {
			case <- stop:
				// early stopping

				// reseting counter
				r.mu.Lock()
				r.counter -= 1
				r.mu.Unlock()

				return false
				
			case <- time.After(time.Duration(delayms) * time.Millisecond):
				// executing task
				task()
				return true
		}
	} else {
		// waiting to delay to elapse
		select {
			case <- time.After(time.Duration(delayms) * time.Millisecond):
				// executing task
				task()
				return true
		}
	}    
}

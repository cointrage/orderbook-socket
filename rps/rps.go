package rps

// Makes sure you do not exceed predefined rate limit. Useful for API calls.

import (
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

func (r *RPS) Run(task func(), stop <-chan struct{}) {
    
    elapsed := time.Since(r.time)
	current := float64(r.counter)/elapsed.Seconds()

	// incrementing counter
	r.mu.Lock()
	r.counter += 1
	counter := r.counter
	r.mu.Unlock()

	if current <= r.rps {
		task()
	} else {
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

					return
				case <- time.After(time.Duration(delayms) * time.Millisecond):
					// executing task
					task()
			}
		} else {
			// waiting to delay to elapse
			select {
				case <- time.After(time.Duration(delayms) * time.Millisecond):
					// executing task
					task()
			}
		}

		
	}
    
}

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
    	mu: sync.Mutex{},
        work: make(chan func()),
    }
}

func (r *RPS) Run(task func(), stop <-chan struct{}) (bool) {
    
	// incrementing counter
	r.mu.Lock()
	
	if r.counter == 0 {
		// reseting time for the first call
		r.time = time.Now()
	}

	r.counter += 1
	counter := r.counter

	elapsed := time.Since(r.time)
	current := float64(r.counter)/elapsed.Seconds()
	
	r.mu.Unlock()


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

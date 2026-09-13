package iplimiter

import (
	"sync"
	"sync/atomic"
	"time"

	"golang.org/x/time/rate"
)

type ipEntry struct {
	limiter  *rate.Limiter
	lastSeen atomic.Int64
}

type IPLimiter struct {
	limiters typedSyncMap[string, *ipEntry]
	config   Config

	stop     chan struct{}
	stopOnce sync.Once
	wg       sync.WaitGroup
}

func New(config Config) *IPLimiter {
	il := &IPLimiter{
		config: config,
		stop:   make(chan struct{}),
	}

	il.wg.Add(1)
	go il.janitor()

	return il
}

func (il *IPLimiter) Get(ip string) *rate.Limiter {
	if e, ok := il.limiters.Load(ip); ok {
		e.lastSeen.Store(time.Now().UnixNano())
		return e.limiter
	}

	e := &ipEntry{
		limiter: rate.NewLimiter(il.config.Rps, il.config.Burst),
	}
	e.lastSeen.Store(time.Now().UnixNano())

	actual, _ := il.limiters.LoadOrStore(ip, e)
	actual.lastSeen.Store(time.Now().UnixNano())

	return actual.limiter
}

func (il *IPLimiter) janitor() {
	defer il.wg.Done()

	t := time.NewTicker(il.config.CleanupInterval)
	defer t.Stop()

	for {
		select {
		case <-t.C:
			il.cleanup()
		case <-il.stop:
			return
		}
	}
}

func (il *IPLimiter) cleanup() {
	now := time.Now().UnixNano()

	il.limiters.Range(func(ip string, ie *ipEntry) bool {
		if now-ie.lastSeen.Load() > il.config.TTL.Nanoseconds() {
			il.limiters.Delete(ip)
		}
		return true
	})
}

func (il *IPLimiter) Close() {
	il.stopOnce.Do(func() {
		close(il.stop)
	})
	il.wg.Wait()
}

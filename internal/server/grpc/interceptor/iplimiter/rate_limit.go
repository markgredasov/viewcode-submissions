package iplimiter

import "golang.org/x/time/rate"

type IPLimiter struct {
	limiters typedSyncMap[string, *rate.Limiter]
	config   Config
}

func New(config Config) *IPLimiter {
	return &IPLimiter{
		config: config,
	}
}

func (il *IPLimiter) Get(ip string) *rate.Limiter {
	if l, ok := il.limiters.Load(ip); ok {
		return l
	}

	l := rate.NewLimiter(il.config.Rps, il.config.Burst)
	actual, _ := il.limiters.LoadOrStore(ip, l)

	return actual
}

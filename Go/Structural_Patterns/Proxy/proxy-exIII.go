package proxy

import (
	"fmt"
	"sync"
	"time"
)

type FibonacciService interface {
	Compute(n int) int
}

type RealFibonacciService struct{}

func (r *RealFibonacciService) Compute(n int) int {
	if n <= 1 {
		return n
	}
	time.Sleep(1 * time.Second)
	return r.Compute(n-1) + r.Compute(n-2)
}

type CacheProxy struct {
	Service  FibonacciService
	Cache    map[int]int
	cacheMux sync.Mutex
}

func NewCacheProxy(service FibonacciService) *CacheProxy {
	return &CacheProxy{
		Service: service,
		Cache:   make(map[int]int),
	}
}

func (p *CacheProxy) Compute(n int) int {
	p.cacheMux.Lock()
	defer p.cacheMux.Unlock()

	if val, exists := p.Cache[n]; exists {
		fmt.Printf("Cache hit for: %d", n)
		return val
	}

	fmt.Printf("Computing Fibonacci for: %d", n)
	result := p.Service.Compute(n)
	p.Cache[n] = result
	return result
}

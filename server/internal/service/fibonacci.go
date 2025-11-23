package service

import (
	"app/internal/domain"
	"sync"
)

type FibonacciService struct {
	cache map[int]uint64
}

func NewFibonacciService() *FibonacciService {
	return &FibonacciService{
		cache: make(map[int]uint64),
	}
}

func (s *FibonacciService) Calculate(n int) uint64 {
	if n <= 1 {
		return uint64(n)
	}
	if val, exists := s.cache[n]; exists {
		return val
	}
	result := s.Calculate(n-1) + s.Calculate(n-2)
	// s.cache[n] = result
	return result
}

// CalculateConcurrent calculates Fibonacci numbers concurrently using goroutines
func (s *FibonacciService) CalculateConcurrent(count int) []domain.FibResult {
	results := make([]domain.FibResult, count)
	var wg sync.WaitGroup
	resultChan := make(chan domain.FibResult, count)

	// Launch a goroutine for each Fibonacci calculation
	for i := range count {
		wg.Add(1)
		go func(index int) {
			defer wg.Done()
			resultChan <- domain.FibResult{
				Index: index,
				Value: (s.Calculate((index))),
			}
		}(i)
	}

	// Close channel when all goroutines complete
	go func() {
		wg.Wait()
		close(resultChan)
	}()

	// Collect results from channel
	for result := range resultChan {
		results[result.Index] = result
	}

	return results
}

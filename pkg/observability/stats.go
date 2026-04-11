package observability

import (
	"sync"
	"time"
)

type Stats struct {
	Version string
	StartedAt time.Time
	RequestsTotal int64
	Responses map[int]int64
	mu sync.Mutex
}

func NewStats(version string) *Stats {
	return &Stats{
		Version: version,
		StartedAt: time.Now(),
		Responses: make(map[int]int64),
	}
}

func (s *Stats)Inc(status int) {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.RequestsTotal++
	s.Responses[status]++
}
package middlewares

import (
	"fmt"
	"redis-kvgo/store"
	"time"
)

type MetricsMiddleware struct {
	inner store.Storer

	getCalls    int
	setCalls    int
	deleteCalls int
	lenCalls    int

	getMisses int // Get error - key was not there

	totalGetLatency time.Duration // time spent on inner.Get operation
	totalSetLatency time.Duration // time spent on inner.Set operation
}

func NewMetricsMiddleware(inner store.Storer) *MetricsMiddleware {
	return &MetricsMiddleware{inner: inner}
}

func (m *MetricsMiddleware) Set(key, value string) error {
	start := time.Now()
	err := m.inner.Set(key, value)
	m.totalSetLatency += time.Since(start)
	m.setCalls++
	return err
}

func (m *MetricsMiddleware) Get(key string) (string, error) {
	start := time.Now()
	val, err := m.inner.Get(key)
	m.totalGetLatency += time.Since(start) // time of the get lifecycle
	if err != nil {
		m.getMisses++
	}

	m.getCalls++
	return val, err
}

func (m *MetricsMiddleware) Delete(key string) {
	m.inner.Delete(key)
	m.deleteCalls++
}

func (m *MetricsMiddleware) Keys() []string {
	return m.inner.Keys()
}

func (m *MetricsMiddleware) Len() int {
	got := m.inner.Len()
	m.lenCalls++
	return got
}

func (m *MetricsMiddleware) Report() {
	fmt.Println("Metrics Report:")
	fmt.Printf("Get Calls: %d\n", m.getCalls)
	fmt.Printf("Set Calls: %d\n", m.setCalls)
	fmt.Printf("Delete Calls: %d\n", m.deleteCalls)
	fmt.Printf("Len Calls: %d\n", m.lenCalls)
	fmt.Printf("Get Misses: %d\n", m.getMisses)
	fmt.Printf("Total Get Latency: %v\n", m.totalGetLatency)
	fmt.Printf("Total Set Latency: %v\n", m.totalSetLatency)
}

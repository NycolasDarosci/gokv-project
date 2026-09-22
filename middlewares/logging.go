package middlewares

import (
	"log"
	"os"
	"redis-kvgo/store"
)

// This middleware is to undestand composition with interface
type LoggingMiddleware struct {
	inner  store.Storer // Composite interface
	logger *log.Logger
}

func NewLoggingMiddleware(inner store.Storer) *LoggingMiddleware {
	return &LoggingMiddleware{
		inner:  inner,
		logger: log.New(os.Stdout, "[log] ", log.Ltime|log.Lmicroseconds),
	}
}

func (l *LoggingMiddleware) Set(key, value string) error {
	l.logger.Printf("SET %s", key)
	return l.inner.Set(key, value)
}

func (l *LoggingMiddleware) Get(key string) (string, error) {
	got, err := l.inner.Get(key)
	if err != nil {
		l.logger.Printf("GET %q -> miss (%v)", key, err)
	} else {
		l.logger.Printf("GET %q -> hit", key)
	}
	return got, err
}

func (l *LoggingMiddleware) Delete(key string) {
	l.logger.Printf("DELETE %s", key)
	l.inner.Delete(key)
}

func (l *LoggingMiddleware) Keys() []string {
	got := l.inner.Keys()
	l.logger.Printf("KEYS %v", got)
	return got
}

func (l *LoggingMiddleware) Len() int {
	got := l.inner.Len()
	l.logger.Printf("LEN %d", got)
	return got
}

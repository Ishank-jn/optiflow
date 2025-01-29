package middleware

import (
    "net/http"
    "sync"
    "time"
)

type RateLimiter struct {
    tokens      int
    maxTokens   int
    refillRate  time.Duration
    lastRefill  time.Time
    mu          sync.Mutex
}

func NewRateLimiter(maxTokens int, refillRate time.Duration) *RateLimiter {
    return &RateLimiter{
        tokens:     maxTokens,
        maxTokens:  maxTokens,
        refillRate: refillRate,
        lastRefill: time.Now(),
    }
}

func (rl *RateLimiter) Allow() bool {
    rl.mu.Lock()
    defer rl.mu.Unlock()

    now := time.Now()
    elapsed := now.Sub(rl.lastRefill)
    if elapsed > rl.refillRate {
        rl.tokens = rl.maxTokens
        rl.lastRefill = now
    }

    if rl.tokens > 0 {
        rl.tokens--
        return true
    }
    return false
}

func RateLimitMiddleware(rl *RateLimiter) func(http.Handler) http.Handler {
    return func(next http.Handler) http.Handler {
        return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
            if !rl.Allow() {
                http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
                return
            }
            next.ServeHTTP(w, r)
        })
    }
}
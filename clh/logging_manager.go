package clh

import (
	"sync"
	"time"
)

type LoggingConfig struct {
	Enabled         bool
	StartTime       time.Time
	Duration        time.Duration
	RequestCount    int
	MaxRequests     int
	ApplicationName string
	mu              sync.Mutex
}

type LoggingManager struct {
	configs map[string]*LoggingConfig
	mu      sync.Mutex
}

func NewLoggingManager() *LoggingManager {
	return &LoggingManager{
		configs: make(map[string]*LoggingConfig),
	}
}

func (lm *LoggingManager) EnableLogging(appName string, duration time.Duration, maxRequests int) {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	lm.configs[appName] = &LoggingConfig{
		Enabled:         true,
		StartTime:       time.Now(),
		Duration:        duration,
		RequestCount:    0,
		MaxRequests:     maxRequests,
		ApplicationName: appName,
	}
}

func (lm *LoggingManager) ShouldLog(appName string) bool {
	lm.mu.Lock()
	defer lm.mu.Unlock()
	config, exists := lm.configs[appName]
	if !exists {
		return false
	}

	if !config.Enabled || time.Since(config.StartTime) > config.Duration {
		config.Enabled = false
		return false
	}

	if config.RequestCount >= config.MaxRequests {
		config.Enabled = false
		return false
	}

	config.RequestCount++
	return true
}

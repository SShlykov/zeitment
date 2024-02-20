package circuitbreaker

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
)

func TestNewCircuitBreaker(t *testing.T) {
	cb := NewCircuitBreaker(10, 0.5, 1*time.Minute, 1*time.Minute)
	assert.NotNil(t, cb)
	assert.Equal(t, Closed, cb.state)
	assert.Equal(t, 10, cb.requestLimit)
	assert.Equal(t, 0.5, cb.errorThresholdPercentage)
	assert.Equal(t, 1*time.Minute, cb.intervalDuration)
	assert.Equal(t, 1*time.Minute, cb.openStateTimeout)
}

func TestCircuitBreakerTransitionToOpenFromClosed(t *testing.T) {
	requestLimit := 5
	errorThreshold := 0.2 // 20%
	intervalDuration := 1 * time.Minute
	openStateTimeout := 1 * time.Minute

	cb := NewCircuitBreaker(requestLimit, errorThreshold, intervalDuration, openStateTimeout)

	for i := 0; i < int(float64(requestLimit)*errorThreshold)+1; i++ {
		cb.Execute(func() error {
			return errors.New("error")
		})
	}

	assert.Equal(t, Open, cb.state)
}

func TestCircuitBreakerHalfOpenToClosed(t *testing.T) {
	cb := NewCircuitBreaker(1, 0.5, 1*time.Minute, 1*time.Minute)
	cb.state = HalfOpen

	err := cb.Execute(func() error {
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, Closed, cb.state)
}

func TestCircuitBreakerOpenToClosed(t *testing.T) {
	cb := NewCircuitBreaker(1, 0.5, 1*time.Minute, 1*time.Minute)
	cb.state = Open
	cb.nextAttempt = time.Now().Add(-2 * time.Minute)

	err := cb.Execute(func() error {
		return nil
	})

	assert.Nil(t, err)
	assert.Equal(t, Closed, cb.state)
}

func TestCircuitBreakerHalfOpenToOneRequest(t *testing.T) {
	cb := NewCircuitBreaker(1, 0.5, 1*time.Minute, 1*time.Minute)
	cb.state = HalfOpen

	err := cb.Execute(func() error {
		return errors.New("error")
	})
	assert.NotNil(t, err)

	err = cb.Execute(func() error {
		return nil
	})
	assert.NotNil(t, err)
	assert.Equal(t, "circuit breaker open", err.Error())
}

package circuitbreaker

import (
	"errors"
	"sync"
	"time"
)

type State int

const (
	Closed State = iota
	Open
	HalfOpen
)

type CircuitBreaker struct {
	mutex                    sync.Mutex
	state                    State
	requestCount             int
	errorCount               int
	requestLimit             int
	errorThresholdPercentage float64
	intervalDuration         time.Duration
	openStateTimeout         time.Duration
	nextIntervalStartTime    time.Time
	nextAttempt              time.Time
}

// NewCircuitBreaker создает новый экземпляр CircuitBreaker. Механизм CircuitBreaker следит за количеством запросов и количеством ошибок.
// requestLimit - максимальное количество запросов в интервале
// errorThresholdPercentage - процент ошибок при котором срабатывает отключение
// intervalDuration - длительность интервала для сброса счетчиков
// openStateTimeout - длительность периода ожидания в состоянии Open
func NewCircuitBreaker(requestLimit int, errorThresholdPercentage float64, intervalDuration, openStateTimeout time.Duration) *CircuitBreaker {
	return &CircuitBreaker{
		state:                    Closed,
		requestLimit:             requestLimit,
		errorThresholdPercentage: errorThresholdPercentage,
		intervalDuration:         intervalDuration,
		openStateTimeout:         openStateTimeout,
		nextIntervalStartTime:    time.Now().Add(intervalDuration),
	}
}

func (cb *CircuitBreaker) Execute(action func() error) error {
	cb.mutex.Lock()
	now := time.Now()

	if now.After(cb.nextIntervalStartTime) {
		cb.resetCounters(now)
	}

	switch cb.state {
	case Closed:
		if cb.requestCount >= cb.requestLimit {
			cb.mutex.Unlock()
			return errors.New("request limit exceeded")
		}
		cb.requestCount++
		cb.mutex.Unlock()

		err := action()
		if err != nil {
			cb.recordFailure()
			return err
		}
		cb.recordSuccess()
		return nil

	case Open:
		if now.Before(cb.nextAttempt) {
			cb.mutex.Unlock()
			return errors.New("circuit breaker open")
		}

		cb.state = HalfOpen
		cb.requestCount = 1
		cb.mutex.Unlock()

		err := action()
		if err != nil {
			cb.recordFailure()
			return err
		}
		cb.recordSuccess()
		return nil

	case HalfOpen:
		if cb.requestCount >= 1 {
			cb.mutex.Unlock()
			return errors.New("circuit breaker half-open: waiting for test request result")
		}
		cb.requestCount++
		cb.mutex.Unlock()

		err := action()
		if err != nil {
			cb.recordFailure()
			return err
		}
		cb.recordSuccess()
		return nil
	}

	return errors.New("unexpected circuit breaker state")
}

func (cb *CircuitBreaker) recordFailure() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	cb.errorCount++
	if cb.state == HalfOpen {
		cb.transitionToOpen()
	} else {
		if float64(cb.errorCount)/float64(cb.requestCount) >= cb.errorThresholdPercentage {
			cb.transitionToOpen()
		}
	}
}

func (cb *CircuitBreaker) recordSuccess() {
	cb.mutex.Lock()
	defer cb.mutex.Unlock()

	if cb.state == HalfOpen {
		cb.transitionToClosed()
	}
	cb.errorCount = 0
}

func (cb *CircuitBreaker) transitionToOpen() {
	cb.state = Open
	cb.nextAttempt = time.Now().Add(cb.openStateTimeout)
	cb.resetCounters(time.Now())
}

func (cb *CircuitBreaker) transitionToClosed() {
	cb.state = Closed
	cb.resetCounters(time.Now())
}

func (cb *CircuitBreaker) resetCounters(now time.Time) {
	cb.requestCount = 0
	cb.errorCount = 0
	cb.nextIntervalStartTime = now.Add(cb.intervalDuration)
}

func (cb *CircuitBreaker) evaluateState() {
	if float64(cb.errorCount)/float64(cb.requestCount) >= cb.errorThresholdPercentage {
		cb.state = Open
		cb.nextAttempt = time.Now().Add(cb.openStateTimeout)
	}
}

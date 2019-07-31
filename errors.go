package hystrix

import (
	"github.com/afex/hystrix-go/hystrix"
)

// IsTimeoutError represents err is hystrix circuit timeout error
func IsTimeoutError(err error) bool {
	return err == hystrix.ErrTimeout
}

// IsCircuitOpenError represents err is hystrix circuit open error
func IsCircuitOpenError(err error) bool {
	return err == hystrix.ErrCircuitOpen
}

// IsMaxConcurrencyError represents err is hystrix max concurrency err
func IsMaxConcurrencyError(err error) bool {
	return err == hystrix.ErrMaxConcurrency
}

package circuitbreaker

import (
	"context"

	"github.com/sony/gobreaker"

	"github.com/chenleji/kit/endpoint"
)

// Gobreaker returns an endpoint.Middleware that implements the circuit
// breaker pattern using the sony/gobreaker package. Only errors returned by
// the wrapped endpoint count against the circuit breaker's error count.
//
// See http://godoc.org/github.com/sony/gobreaker for more information.
func Gobreaker(cb *gobreaker.CircuitBreaker) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, method, rawUrl string, headers map[string]string, reqObj interface{}, respObj interface{}) (interface{}, error) {
			return cb.Execute(func() (interface{}, error) { return next(ctx, method, rawUrl, headers, reqObj, respObj) })
		}
	}
}

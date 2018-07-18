package zipkin

import (
	"context"

	"github.com/openzipkin/zipkin-go"
	"github.com/openzipkin/zipkin-go/model"

	"github.com/chenleji/kit/endpoint"
)

// TraceEndpoint returns an Endpoint middleware, tracing a Go kit endpoint.
// This endpoint tracer should be used in combination with a Go kit Transport
// tracing middleware or custom before and after transport functions as
// propagation of SpanContext is not provided in this middleware.
func TraceEndpoint(tracer *zipkin.Tracer, name string) endpoint.Middleware {
	return func(next endpoint.Endpoint) endpoint.Endpoint {
		return func(ctx context.Context, method, rawUrl string, headers map[string]string, reqObj interface{}, respObj interface{}) (interface{}, error) {
			var sc model.SpanContext
			if parentSpan := zipkin.SpanFromContext(ctx); parentSpan != nil {
				sc = parentSpan.Context()
			}
			sp := tracer.StartSpan(name, zipkin.Parent(sc))
			defer sp.Finish()

			ctx = zipkin.NewContext(ctx, sp)
			return next(ctx, method, rawUrl, headers, reqObj, respObj)
		}
	}
}

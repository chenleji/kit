package endpoints

import "context"

import "github.com/chenleji/kit/endpoint"

import "github.com/chenleji/kit/cmd/kitgen/testdata/anonfields/default/service"

type FooRequest struct {
	I int
	S string
}
type FooResponse struct {
	I   int
	Err error
}

func MakeFooEndpoint(s service.Service) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (interface{}, error) {
		req := request.(FooRequest)
		i, err := s.Foo(ctx, req.I, req.S)
		return FooResponse{I: i, Err: err}, nil
	}
}

type Endpoints struct {
	Foo endpoint.Endpoint
}

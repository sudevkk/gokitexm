package main



import (
	"github.com/go-kit/kit/endpoint"
	"github.com/go-kit/kit/log"
	"context"
)


type Middleware func(endpoint.Endpoint) (endpoint.Endpoint)

func loggerMiddleware(logger log.Logger) (Middleware) {
	return func(next endpoint.Endpoint) (endpoint.Endpoint) {
		return func(ctx context.Context, req interface{}) (interface{}, error) {
			logger.Log("msg", "Callng Endpoint")
			defer logger.Log("msg", "Called Endpoint")
			return next(ctx, req)
		}
	}
}
package main

import (
	"context"

	"connectrpc.com/connect"
)

// ServiceVersionInterceptor creates an interceptor that adds a version header for a specific service
func ServiceVersionInterceptor(serviceName, version string) connect.UnaryInterceptorFunc {
	return func(next connect.UnaryFunc) connect.UnaryFunc {
		return func(ctx context.Context, req connect.AnyRequest) (connect.AnyResponse, error) {
			// Call the handler
			res, err := next(ctx, req)
			if err != nil {
				return nil, err
			}

			// Add the version header to the response
			res.Header().Set(serviceName+"-Version", version)
			return res, nil
		}
	}
}

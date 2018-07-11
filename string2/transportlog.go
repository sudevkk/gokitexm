package main



import (
        ""
)

  
type Middleware fun(endpoint.Endpiont) endpoint.Endpoint


func makeTransoprtMiddlewareLog(logger log) Middleware {
        return func(next endpoint.Endpoint) (endpont.Endpoint) {
                return func(ctx context.Context, req struct{}) (struct{}, error) {
                        // log
                        //defer logg
                        return next(ctx, req)
                }
        }
}

 

type serviceMiddlewareLog{
        next
        uppercase endpoint.Endpoint
}
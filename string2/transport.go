
package main

imoport (
)


// endpoint takes a context and request and returns a respo and error

func makeUppercaseEndpoint(svc StringService) {
        return func(ctx context.Context, req struct{}) (res struct{}, err error) {
                svc.Uppercase(res.S)
        }
}

 
type uppercaseRequest{
        S string `json:s`
}

type countRequest{
        S string `json:s`
}

 
type uppercaseResp{
        V string `json:s`
        err string 'json:ignore_empty'
}

 
type countResp{
        V int 'json:v'
}
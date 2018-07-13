package main


import (
	"github.com/go-kit/kit/log"
	"net/http"
	httptransport "github.com/go-kit/kit/transport/http"
	"os"
)


func main() {

	logger := log.NewLogfmtLogger(os.Stderr)
	svc := stringService{}
	uppercaseWithLogger := loggerMiddleware(log.With(logger, "method", "uppercase"))


	uppercaseHandler := httptransport.NewServer(
		uppercaseWithLogger(makeUppercaseEndpoint(svc)),
		decodeUppercaseRequest,
		encodeResponse,
	)

	countWithLogger := loggerMiddleware(log.With(logger, "method", "count"))(makeCountEndpoint(svc))
	countHandler := httptransport.NewServer(
		countWithLogger,
		decodeCountRequest,
		encodeResponse,
	)

	http.Handle("/uppercase", uppercaseHandler)
	http.Handle("/count", countHandler)
	http.ListenAndServe(":8080", nil)
}




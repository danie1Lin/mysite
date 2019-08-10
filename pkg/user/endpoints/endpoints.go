package endpoints

import (
	"github.com/go-kit/kit/endpoint"
)

type EndPoints struct {
	LoginEndPoint  endpoint.Endpoint
	SignUpEndPoint endpoint.Endpoint
}

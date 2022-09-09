package httpclient

import (
	"github.com/go-resty/resty/v2"
)

func createRequest(opts []ConfigureRequestOptions) *resty.Request {
	client := resty.New()

	options := configureOptions(opts)

	setTLSConfig(client, options)

	configureRetry(client, options)

	request := client.R()

	setHeaders(options.headers, request)

	setAuth(client, options)

	return request
}

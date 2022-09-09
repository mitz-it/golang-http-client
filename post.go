package httpclient

func Post[TBody, TContent any](url string, body TBody, options ...ConfigureRequestOptions) HttpResponse[TContent] {
	request := createRequest(options)

	serializeBody(request, body)

	response, err := request.Post(url)

	model, serializationError := deserializeContent[TContent](response)

	return NewHttpResponse(model, response.StatusCode(), err, serializationError)
}

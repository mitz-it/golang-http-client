package httpclient

func Get[TContent any](url string, options ...ConfigureRequestOptions) HttpResponse[TContent] {
	request := createRequest(options)

	response, requestError := request.Get(url)

	model, serializationError := deserializeContent[TContent](response)

	return NewHttpResponse(model, response.StatusCode(), requestError, serializationError)
}

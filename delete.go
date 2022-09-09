package httpclient

func Delete(url string, options ...ConfigureRequestOptions) HttpResponse[struct{}] {
	request := createRequest(options)

	response, err := request.Delete(url)

	model, serializationError := deserializeContent[struct{}](response)

	return NewHttpResponse(model, response.StatusCode(), err, serializationError)
}

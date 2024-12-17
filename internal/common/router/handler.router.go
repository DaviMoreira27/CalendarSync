package router

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/DaviMoreira27/CalendarSync/internal/common/enums"
	"github.com/DaviMoreira27/CalendarSync/internal/common/log"
	"github.com/DaviMoreira27/CalendarSync/internal/common/types"
)

func createHttpClient() *http.Client {
	return &http.Client{
		Timeout: 15 * time.Second,
	}
}

func setHeaders(request *http.Request, headers *http.Header) {
	if headers != nil {
		for key, values := range *headers {
			request.Header[key] = values
		}
	}
}

func handleResponseBody[R any](response *http.Response, responseBody *R) error {
	return json.NewDecoder(response.Body).Decode(&responseBody)
}

func parseHttpMethod(method string) enums.HttpMethod {
	switch strings.ToUpper(method) {
	case "GET":
		return enums.Get
	case "POST":
		return enums.Post
	case "PUT":
		return enums.Put
	case "DELETE":
		return enums.Delete
	default:
		return enums.Get
	}
}

func handleErrors(err error, response *http.Response) {
	if err == nil {
		return
	}

	switch err.(type) {
	case types.Error:
		log.WriteError(
			types.HttpErrorType{
				Message:    "Sei la qual erro",
				StatusCode: response.StatusCode,
			},
			types.HttpOperation{
				Method:    parseHttpMethod(response.Request.Method),
				Operation: "Indefinida",
			})

	default:
		log.WriteError(
			types.HttpErrorType{
				Message:    "Sei la qual erro",
				StatusCode: http.StatusInternalServerError,
			},
			types.HttpOperation{
				Method:    parseHttpMethod(response.Request.Method),
				Operation: "Indefinida",
			})

	}
}

func RequestHandler[B any, R any](
	url string,
	method enums.HttpMethod,
	body *B,
	headers *http.Header) (R, error) {

	var responseBody R
	var requestBody *bytes.Buffer
	client := createHttpClient()

	if body != nil {
		jsonBody, err := json.Marshal(body)
		requestBody = bytes.NewBuffer(jsonBody)

		if err == nil {
			return responseBody, err
		}
	}

	var request *http.Request
	var err error
	switch method {
	case enums.Get:
		request, err = http.NewRequest("GET", url, nil)
	case enums.Post:
		if (requestBody == nil) {
			return responseBody, &types.InternalError{
				Err: errors.New("o request body é obrigatório para requisições POST"),
			}
		}
		request, err = http.NewRequest("POST", url, requestBody)
	case enums.Put:
		request, err = http.NewRequest("PUT", url, requestBody)
	case enums.Delete:
		request, err = http.NewRequest("DELETE", url, nil)
	default:
		log.WriteError(
			types.HttpErrorType{
				Message:    "Metódo HTTP não existente",
				StatusCode: 400,
			},
			types.HttpOperation{
				Method:    method,
				Operation: "Indefinida",
			})
	}

	if err != nil {
		return responseBody, err
	}

	setHeaders(request, headers)

	response, err := client.Do(request)

	handleErrors(err, response)

	defer response.Body.Close()

	err = handleResponseBody(response, &responseBody)

	if err != nil {
		return responseBody, err
	}

	return responseBody, nil
}

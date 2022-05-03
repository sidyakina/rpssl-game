package handlerwrapper

import (
	"io/ioutil"
	"log"
	"net/http"
)

type GetHandler func() (response []byte, responseCode int)
type PostHandler func(request []byte) (response []byte, responseCode int)

func Post(writer http.ResponseWriter, request *http.Request, requestName string, handler PostHandler) {
	if request.Method != http.MethodPost {
		writer.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	log.Printf("handling get %v request", requestName)

	data, err := ioutil.ReadAll(request.Body)
	if err != nil {
		log.Printf("failed to read request body for request %v: %v", requestName, err)

		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	defer func() {
		err := request.Body.Close()
		if err != nil {
			log.Printf("failed to close request body: %v", err)
		}
	}()

	response, code := handler(data)

	writeResponse(requestName, writer, response, code)
}

func Get(writer http.ResponseWriter, request *http.Request, requestName string, handler GetHandler) {
	if request.Method != http.MethodGet {
		writer.WriteHeader(http.StatusMethodNotAllowed)

		return
	}

	log.Printf("handling get %v request", requestName)

	response, code := handler()

	writeResponse(requestName, writer, response, code)
}

func writeResponse(requestName string, writer http.ResponseWriter, response []byte, code int) {
	log.Printf("response to %v request with %s and status code %v", requestName, response, code)

	writer.WriteHeader(code)
	writer.Header().Set("Content-Type", "application/json")

	_, err := writer.Write(response)
	if err != nil {
		log.Printf("failed to write response for request %v: %v", requestName, err)

		return
	}

	log.Printf("request %v handled", requestName)
}

package handlerwrapper

import (
	"log"
	"net/http"
)

type GetHandler func() (response []byte, responseCode int)
type PostHandler func(request []byte) (response []byte, responseCode int)

func Post(writer http.ResponseWriter, request *http.Request, requestName string, handler PostHandler) {
	body, err := request.GetBody()
	if err != nil {
		log.Printf("failed to get request body: %v", err)

		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	defer func() {
		err := body.Close()
		if err != nil {
			log.Printf("failed to close body: %v", err)
		}
	}()

	// todo
}

func Get(writer http.ResponseWriter, _ *http.Request, requestName string, handler GetHandler) {
	log.Printf("handling get %v request", requestName)

	response, code := handler()

	_, err := writer.Write(response)
	if err != nil {
		log.Printf("failed to write response for request %v: %v", requestName, err)

		writer.WriteHeader(http.StatusInternalServerError)

		return
	}

	writer.WriteHeader(code)

	log.Printf("request %v handled", requestName)
}

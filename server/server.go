package server

import (
	"fmt"
	"net/http"
	"os"
	"sync/atomic"
)

type requestHandler struct {
	Busy        int32
	Destination string
}

func (requestHandler *requestHandler) ServeHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	switch request.Method {
	case "PUT":
		requestHandler.handlePut(responseWriter, request)
	default:
		responseWriter.WriteHeader(http.StatusMethodNotAllowed)
	}
}

func (requestHandler *requestHandler) handlePut(responseWriter http.ResponseWriter, request *http.Request) {
	var err error
	var file *os.File

	// try to acquire lock
	lock := atomic.CompareAndSwapInt32(&requestHandler.Busy, 0, 1)
	if !lock {
		// already locked - respond with error
		// TODO: set Retry-After header
		responseWriter.WriteHeader(http.StatusServiceUnavailable)
		return
	}
	defer atomic.StoreInt32(&requestHandler.Busy, 0)

	// lock acquired!
	// open file for writing
	file, err = os.OpenFile(requestHandler.Destination, os.O_WRONLY|os.O_TRUNC|os.O_SYNC, 0000)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		// send error response
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer file.Close()

	// copy pipe to destination
	_, err = file.ReadFrom(request.Body)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		// send error response
		responseWriter.WriteHeader(http.StatusInternalServerError)
		return
	}

	// send success response
	responseWriter.WriteHeader(http.StatusNoContent)
}

func Start(host string, port uint, destination string) bool {
	var err error
	http.Handle("/", &requestHandler{Busy: 0, Destination: destination})
	err = http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), nil)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		return false
	}

	return true
}

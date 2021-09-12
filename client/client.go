package client

import (
	"fmt"
	"net/http"
	"os"
)

func Write(uri string, source string) bool {
	var err error
	var file *os.File
	var response *http.Response
	var request *http.Request

	// open source for reading
	file, err = os.Open(source)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return false
	}
	defer file.Close()

	// create http PUT request to given uri
	request, err = http.NewRequest(http.MethodPut, uri, file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return false
	}
	request.Header.Set("Content-Type", "application/octet-stream")

	// execute http request
	response, err = http.DefaultClient.Do(request)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err.Error())

		return false
	}
	if response.StatusCode >= 200 && response.StatusCode < 300 {
		return true
	} else {
		fmt.Fprintf(os.Stderr, "server returned %d\n", response.StatusCode)
	}

	return false
}

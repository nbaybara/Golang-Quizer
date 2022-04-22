package cmd

import (
	"io/ioutil"
	"log"
	"net/http"
)

var url = "http://localhost:8083"

func getQuestions(baseAPI string) []byte {

	request, err := http.NewRequest(
		http.MethodGet, //method
		baseAPI,        //url
		nil,            //body
	)

	handleError(err)

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Printf("Could not make a request. %v", err)
	}

	responseBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Printf("Could not read response body. %v", err)
	}

	return responseBytes
}

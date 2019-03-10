package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("http://golang.org/") // Send a HTTP GET request to the URL
	if err != nil {
		fmt.Println(err) // If there is an error, print it and exit
		os.Exit(1)
	}
	defer response.Body.Close() // Close the connection to prevent a resource leak

	contents, err := ioutil.ReadAll(response.Body) // Read the contents of the HTTP GET
	if err != nil {                                // If there is an error
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Printf("%s\n", string(contents)) // Print the contents of the request
}

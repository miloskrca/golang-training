package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

// type Reader interface {
// 	Read(p []byte) (n int, err error)
// }

func reverse(r io.Reader) io.Reader {
	b, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	s := string(b)
	n := len(s)
	runes := make([]rune, n)
	for _, rune := range s {
		n--
		runes[n] = rune
	}
	return strings.NewReader(string(runes[n:]))
}

func main() {
	var r io.Reader
	file, err := os.Open("demo/basic/interface/reader/reader.txt")
	if err != nil {
		log.Fatal(err)
	}
	r = reverse(file)
	reversed, err := ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reversed)

	reader := strings.NewReader("Read will return these bytes")
	r = reverse(reader)
	reversed, err = ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reversed)

	resp, err := http.Get("https://api.ipify.org/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	r = reverse(resp.Body)
	reversed, err = ioutil.ReadAll(r)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", reversed)
}

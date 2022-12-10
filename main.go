package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	http.HandleFunc("/", h)
	http.ListenAndServe(":8080", nil)
}

func h(rw http.ResponseWriter, req *http.Request) {

	r, err := http.Get("https://api.myip.com")
	if err != nil {
		fmt.Fprintf(rw, "%s", err)
		return
	}

	b, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Fprintf(rw, "%s", err)
	}

	fmt.Fprintf(rw, "%s", b)
}

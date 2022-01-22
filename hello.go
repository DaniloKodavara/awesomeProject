package main

import "net/http"

func main() {
	res, err := http.Get("http://google.com.br")
	if err != nil {
		return
	}
}

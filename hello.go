package main

import (
	"errors"
	"fmt"
	"log"
	"net/http"
)

func main() {
	res, err := http.Get("http://google.com.br")
	if err != nil {
		log.Println(err.Error())
	} else {
		fmt.Println(res.Header)
	}

	resSoma, errSoma := soma(10, 10)
	if errSoma != nil {
		log.Fatal(errSoma.Error())
	}

	fmt.Println(resSoma)

}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}

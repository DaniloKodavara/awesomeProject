package main

import (
	"errors"
	"fmt"
)

type Carro struct {
	Nome string
}

func (c *Carro) andou() {
	c.Nome = "Fusca"
	fmt.Println(c.Nome, "andou")
}

func main() {

	carro := Carro{
		Nome: "BMW",
	}

	carro.andou()
	fmt.Println(carro.Nome)

	/*	carro := Carro{
			Nome: "BMW",
		}

		carro.andar()

		resultado := func(x ...int) func() int {
			res := 0

			for _, v := range x {
				res += v
			}
			return func() int {
				return res * res
			}
		}

		res, err := http.Get("http://google.com.br")
		if err != nil {
			log.Println(err.Error())
		} else {
			fmt.Println(res.Header)
		}

		resSoma, errSoma := soma(10, 10)
		if errSoma != nil {
			log.Println(errSoma.Error())
		} else {
			fmt.Println(resSoma)
		}

		fmt.Println(somaComResultado(10, 10))

		fmt.Println(resultado(3, 5, 8, 13, 21)())*/

}

func abc(a *int) {
	*a = 200
}

func soma(x int, y int) (int, error) {
	res := x + y
	if res > 10 {
		return 0, errors.New("Total maior que 10")
	}

	return res, nil
}

func somaComResultado(a int, b int) (result int, str string) {
	result = a + b
	str = "Somou"
	return
}

func somaTudo(x ...int) (resultant int) {
	resultant = 0

	for _, v := range x {
		resultant += v
	}

	return resultant
}

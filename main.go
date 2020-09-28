package main

import "fmt"

func main()  {
	var e float64
	var exponencial int
	e = 1
	for i := 1; i <= 10; i++ {
		exponencial = 1
		for j := 1; j <= i; j++ {
			exponencial = exponencial * j	
		}
		e = e + 1/float64(exponencial)
	}
	fmt.Println("valor de e =",e)
}
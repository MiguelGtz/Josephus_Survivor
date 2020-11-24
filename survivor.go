package main

import (
	"fmt"
)

func josephusSurvivor(n, k int) int {
	var lista []int
	for i := 0; i < n; i++ {
		lista = append(lista, i+1)
	}
	var index, c, aux int
	for true {
		if len(lista) != aux {
			fmt.Println(lista)
			aux = len(lista)
		}
		for i := index; ; i++ {
			c++
			if i == len(lista)-1 && c != k {
				index = 0
				break
			}
			if c == k {
				if i == len(lista)-1 {
					index = 0
				} else {
					index = i
				}
				lista = append(lista[:i], lista[i+1:]...)
				c = 0
				break
			}
		}
		if len(lista) == 1 {
			fmt.Println(lista)
			break
		}
	}
	return lista[0]
}

func main() {
	fmt.Println(josephusSurvivor(11, 19))
}

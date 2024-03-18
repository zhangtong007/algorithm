package main

import (
	"fmt"
)

type Product struct {
	v, p, q int
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var v, p, q int
	products := make([]Product, 1, m+1)
	for i := 0; i < m; i++ {
		fmt.Scan(&v, &p, &q)
		products = append(products, Product{
			v, p, q,
		})
	}
	process(n, products)
}

func process(n int, products []Product) {

}

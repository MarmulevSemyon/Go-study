package jopa

import "fmt"

type Jopa struct {
	A int
}

func (j Jopa) Hello(b int) {
	fmt.Println("hello", j.A*b)

	j.A = 55

	fmt.Println(j)

}

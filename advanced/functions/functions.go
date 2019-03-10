package main

import (
	"fmt"
	"math"
)

type operation func(float64) float64

type calculator struct {
	acc float64
}

func (c *calculator) do(op operation) float64 {
	c.acc = op(c.acc)
	fmt.Println("acc", "=", c.acc)
	return c.acc
}

func add(n float64) operation {
	return func(acc float64) float64 {
		return acc + n
	}
}
func sub(n float64) operation {
	return func(acc float64) float64 {
		return acc - n
	}
}

func mul(n float64) operation {
	return func(acc float64) float64 {
		return acc * n
	}
}
func sqrt() operation {
	return func(n float64) float64 {
		return math.Sqrt(n)
	}
}

func main() {
	var c calculator
	c.do(add(2))
	c.do(sqrt())
	c.do(mul(2))
	c.do(sub(15))
	// fmt.Println("======")
	// c.doMore(mul(0), add(11), sub(2), sqrt())
}

// func (c *calculator) doMore(operations ...operation) float64 {
// 	for _, op := range operations {
// 		c.acc = op(c.acc)
// 	}
// 	fmt.Println("acc", "=", c.acc)
// 	return c.acc
// }

package main

import "fmt"

func main() {
	fmt.Println("Семенова Дарья Сергеевна")
}

package main
import (
	"fmt"
	"math"
)

func main() {
	a := 2.5
	b := 4.6
	x_n := 1.1
	x_k := 3.6
	x_z := 0.5
	x1 := 1.2
	x2 := 1.28
	x3 := 1.36
	x4 := 1.46
	x5 := 2.35
	fmt.Println("Значение функции №18 при данных задачи А =")
	for x := x_n; x <= x_k; x = x + x_z {
		var y1 = (math.Pow((a + b*x), 2.5)) / (1 + (math.Log10(a + b*x)))
		fmt.Println(y1)
	}
	var argument = [5]float64{x1, x2, x3, x4, x5} // Массив значений x1-5
	fmt.Println("Значение функции №18 при данных задачи Б =")
	for i := 0; i < 5; i++ {
		var y2 = (math.Pow((a + b*x), 2.5)) / (1 + (math.Log10(a + b*x)))
		fmt.Println(y2)
	}
	//// лаба 2
	fmt.Println("Семенова Дарья Сергеевна")
}

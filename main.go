package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
)

func main() {
	arg := os.Args[1]
	if len(os.Args) != 2 {
		return
	}
	file, err := os.Open(arg)
	if err != nil {
		log.Println("Couldn't open file")
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var arr []float64
	for scanner.Scan() {
		elem, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Printf("Error")
			return
		}
		arr = append(arr, float64(elem))
	}
	LinRegLine(arr)
	PearsonCoef(arr)
}

func LinRegLine(arr []float64) {
	sum_x := 0.0
	sum_y := 0.0
	sum_xy := 0.0
	sum_xx := 0.0
	p := float64(len(arr))
	for i := 0; i < len(arr); i++ {
		sum_x += float64(i)
		sum_y += arr[i]
		sum_xx += float64(i * i)
		sum_xy += float64(i) * arr[i]

	}
	m := (p*sum_xy - sum_x*sum_y) / (p*sum_xx - sum_x*sum_x)
	b := (sum_y / p) - (m*sum_x)/p

	fmt.Printf("Linear Regression Line: y = %0.6fx + %0.6f\n", m, b)
}

func PearsonCoef(arr []float64) {
	sum_x := 0.0
	sum_y := 0.0
	sum_xy := 0.0
	squareSum_x := 0.0
	squareSum_y := 0.0
	n := float64(len(arr))
	for i := 0; i < len(arr); i++ {
		sum_x += float64(i)
		sum_y += arr[i]
		sum_xy += float64(i) * arr[i]

		squareSum_x += float64(i) * float64(i)
		squareSum_y += arr[i] * arr[i]
	}
	corr := float64((n*sum_xy - sum_x*sum_y)) / (math.Sqrt(float64((n*squareSum_x - sum_x*sum_x) * (n*squareSum_y - sum_y*sum_y))))
	fmt.Printf("Pearson Correlation Coefficient: %0.10f\n", corr)
}

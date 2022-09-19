package main

import (
	"bufio"
	"fmt"
	"log"
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
}

func LinRegLine(arr []float64) {
	sum_x := 0.0
	sum_y := 0.0
	sum_xy := 0.0
	sum_xx := 0.0
	count := 0.0
	for i := 0; i < len(arr); i++ {
		sum_x += float64(i)
		sum_y += arr[i]
		sum_xx += float64(i * i)
		sum_xy += arr[i] * arr[i]
		count++
	}
	m := (count*sum_xy - sum_x*sum_y) / (count*sum_xx - sum_x*sum_x)
	b := (sum_y / count) - (m*sum_x)/count

	result_x := []float64{}
	result_y := []float64{}

	for j := 0; j < len(arr); j++ {
		x := arr[j]
		y := x*m + b
		result_x = append(result_x, x)
		result_y = append(result_y, y)
	}
	fmt.Println(result_x)
	fmt.Println(result_y)
}

// func PearsonCoef(arr []float64) {
// }

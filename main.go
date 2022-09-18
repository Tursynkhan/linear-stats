package main

import (
	"bufio"
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
}

func LinRegLine(arr []float64) {
}

func PearsonCoef(arr []float64) {
}

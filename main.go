package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"

	"mathSkills/functions"
)

func main() {
	args := os.Args[1:]
	if len(args) < 1 {
		fmt.Println("Pas assez d'arguments")
		return
	} else if len(args) > 1 {
		fmt.Println("Trop d'arguments")
		return
	}
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var numbers []float64

	for scanner.Scan() {
		line := scanner.Text()
		line = strings.TrimSpace(line)
		num, err := strconv.ParseFloat(line, 64)
		if err != nil {
			fmt.Println("Erreur de conversion:", err)
			return
		}
		numbers = append(numbers, num)
	}

	fmt.Println("Average:", int(math.Round(functions.Average(numbers))))
	fmt.Println("Median:", int(math.Round(functions.Median(numbers))))
	fmt.Println("Variance:", int(math.Round(functions.Variance(numbers))))
	fmt.Println("Standard Deviation:", int(math.Round(functions.StandDev(numbers))))
}

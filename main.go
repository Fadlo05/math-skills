package main

import (
	"bufio"
	"fmt"
	"mathSkills/functions"
	"os"
	"strconv"
	"strings"
)

func main() {
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
           fmt.Println("Erreur de conversion: %v", err)
		   return
        }
        numbers = append(numbers, num)
    }

    fmt.Println("Average:", functions.Average(numbers))
    fmt.Println("Median:", functions.Median(numbers))
    fmt.Println("Variance:", functions.Variance(numbers))
    fmt.Println("Standard Deviation:", functions.StandDev(numbers))
}
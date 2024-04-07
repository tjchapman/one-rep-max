package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func CalculateOneRepMax(weight, reps float64) float64 {
	return weight / (1.0278 - 0.0278*reps) // The Brzycki formula:= weight / ( 1.0278 – 0.0278 × reps )
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hey, let's calculate your one rep max...")
	fmt.Print("Enter the weight (kg) you hit for reps: ")
	scanner.Scan()
	weight, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	fmt.Print("Enter the reps: ")
	scanner.Scan()
	reps, _ := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
	fmt.Printf("Your one rep max is: %0.fkg!\n", CalculateOneRepMax(weight, reps))
}

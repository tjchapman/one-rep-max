package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main(){
	fmt.Println("Hey, let's calculate your one rep max...")
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter the weight you hit for reps: ")
	scanner.Scan()
	weight_str := scanner.Text()

	fmt.Print("Enter the reps: ")
	scanner.Scan()
	reps_str := scanner.Text()

	if scanner.Err() != nil {
		fmt.Println("Error: ", scanner.Err())
	}


	weight, _ := strconv.ParseFloat(strings.TrimSpace(weight_str), 8)
	reps, _ := strconv.ParseFloat(strings.TrimSpace(reps_str), 8)

	// The Brzycki formula:= weight / ( 1.0278 – 0.0278 × reps )
	one_rep_max := weight / (1.0278 - 0.0278 * reps)

	fmt.Printf("Your one rep max is: %0.fkg!\n", one_rep_max)

}
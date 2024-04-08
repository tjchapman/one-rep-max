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

func FatalIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func Parser(text string) (float64) {
	float, err := strconv.ParseFloat((strings.TrimSpace(text)), 64)
	if err != nil {
		return 0
	}
	return float
}

func CorrectDataTyoe(input float64) bool {
	if input == 0 {
		return false
	} else {
		return true
	}
}

func main()  {
	// - https://patorjk.com/software/taag/#p=display&f=Graffiti&t=Type%20Something%20
	fmt.Println(`
    )                 (                   *                   
 ( /(                 )\ )              (  \     (            
 ))())           (   (()/(   (          )\))(    )\        )  
((_)\    (      ))\   /(_)) ))\ \  )   ((_)()\((((_)(   ( /(  
  ((_)   )\ )  /((_) (_))  /((_)/(/(   (_()((_))\ _ )\  )\()) 
 / _ \  _(_/( (_))   | _ \(_)) ((_)_\  |  \/  |(_)_\(_)((_)\  
| (_) || ' \))/ -_)  |   // -_)| '_ \) | |\/| | / _ \  \ \ /  
 \___/ |_||_| \___|  |_|_\\___|| .__/  |_|  |_|/_/ \_\ /_\_\  
                               |_|                            `)


	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Hey, let's calculate your one rep max...")

	fmt.Print("Enter the weight (kg) you hit for reps: ")
	scanner.Scan()
	weight := Parser(scanner.Text())
	if !CorrectDataTyoe(weight) {
		fmt.Println("Please enter a number or I'll kick you out !!")
		// Banking on the user only making this mistake once lol instead of figuring out some messy loop TODO: clean this up or find a way to do it properly

		scanner.Scan()
		weight := Parser(scanner.Text())

		if !CorrectDataTyoe(weight) {
			panic("Should've entered a number :)")
		}

		fmt.Print("Enter the reps: ")
		scanner.Scan()
		reps := Parser(scanner.Text())

		if !CorrectDataTyoe(reps) {
			panic("Should've entered a number :)")
		}

		fmt.Printf("Your one rep max is: %0.fkg!\n", CalculateOneRepMax(weight, reps))

	} else {
		fmt.Print("Enter the reps: ")
		scanner.Scan()
		reps := Parser(scanner.Text())
	
		fmt.Printf("Your one rep max is: %0.fkg!\n", CalculateOneRepMax(weight, reps))
	}
	
}

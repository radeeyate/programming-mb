package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	fmt.Print("Math problem: ")
	var res1 string
	fmt.Scanln(&res1)
	if strings.HasPrefix(res1, "1") || strings.HasPrefix(res1, "2") || strings.HasPrefix(res1, "3") || strings.HasPrefix(res1, "4") || strings.HasPrefix(res1, "5") || strings.HasPrefix(res1, "6") || strings.HasPrefix(res1, "7") || strings.HasPrefix(res1, "8") || strings.HasPrefix(res1, "9"){
		if strings.Contains(res1, "+") {
			// get index of where "+" is located
			plusLocatedAt := strings.Index(res1, "+")

			// replace "+" with "" so math can be done
			mathToDo := (strings.ReplaceAll(res1, "+", ""))

			// get what the first number to do math is
			firstNumber := (mathToDo[0:plusLocatedAt])
			firstNumberInt, _ := strconv.Atoi(firstNumber)
			
			// get second number to do math with
			mathLength := len(mathToDo)
			secondNumberInt, _ := strconv.Atoi(mathToDo[plusLocatedAt:mathLength])


			// print out the completed problem
			mathSum := firstNumberInt + secondNumberInt
			fmt.Print("gobot: The sum of the two numbers: ")
			fmt.Print(mathSum)
		} else if strings.Contains(res1, "*") {
			// get index of where "*" is located
			xLocatedAt := strings.Index(res1, "*")

			// replace "+" with "" so math can be done
			mathToDo := (strings.ReplaceAll(res1, "*", ""))

			// get what the first number to do math is
			firstNumber := (mathToDo[0:xLocatedAt])
			firstNumberInt, _ := strconv.Atoi(firstNumber)

			// get second number to do math with
			mathLength := len(mathToDo)
			secondNumberInt, _ := strconv.Atoi(mathToDo[xLocatedAt:mathLength])


			// print out the completed problem
			mathSum := firstNumberInt * secondNumberInt
			fmt.Print("gobot: The product of the two numbers: ")
			fmt.Print(mathSum)
		} else if strings.Contains(res1, "/") {
			// get index of where "*" is located
			slashLocatedAt := strings.Index(res1, "/")

			// replace "+" with "" so math can be done
			mathToDo := (strings.ReplaceAll(res1, "/", ""))

			// get what the first number to do math is
			firstNumber := (mathToDo[0:slashLocatedAt])
			firstNumberInt, _ := strconv.Atoi(firstNumber)

			// get second number to do math with
			mathLength := len(mathToDo)
			secondNumberInt, _ := strconv.Atoi(mathToDo[slashLocatedAt:mathLength])
			if secondNumberInt == 0 {
				fmt.Println("gobot: I can't divide by zero!")
			} else {
				// print out the completed problem
				mathSum := firstNumberInt / secondNumberInt
				fmt.Print("gobot: The quotient of the two numbers: ")
				fmt.Print(mathSum)
			}


			
		} else if strings.Contains(res1, "-") {
			// get index of where "*" is located
			minusLocatedAt := strings.Index(res1, "-")

			// replace "+" with "" so math can be done
			mathToDo := (strings.ReplaceAll(res1, "-", ""))

			// get what the first number to do math is
			firstNumber := (mathToDo[0:minusLocatedAt])
			firstNumberInt, _ := strconv.Atoi(firstNumber)
			
			// get second number to do math with
			mathLength := len(mathToDo)
			secondNumberInt, _ := strconv.Atoi(mathToDo[minusLocatedAt:mathLength])


			// print out the completed problem
			mathSum := firstNumberInt - secondNumberInt
			fmt.Print("gobot: The difference of the two numbers: ")
			fmt.Print(mathSum)
		}
	}
}
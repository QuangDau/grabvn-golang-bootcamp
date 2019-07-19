package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"errors"
	"strconv"
	)

func main(){
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("> ")
	for scanner.Scan(){
		text := scanner.Text()
		
		if err := evalExpression(text); err != nil {
			fmt.Printf("Error!!! %v\n", err)
		}

		fmt.Println("> ")
	}
}

func evalExpression(text string) error {
	array_s := strings.Split(text, " ")

	agr1, err1 := strconv.ParseFloat(array_s[0], 32)
	if err1 != nil {
		return errors.New("Agurment 1 is not a number")
	}

	op := array_s[1]

	agr2, err2 := strconv.ParseFloat(array_s[2], 32)
	if err2 != nil {
		return errors.New("Agurment 2 is not a number")
	}

	switch op {
	case "+":
		fmt.Printf("%v + %v = %v\n", agr1, agr2, agr1 + agr2)
	case "-":
		fmt.Printf("%v - %v = %v\n", agr1, agr2, agr1 - agr2)
	case "*":
		fmt.Printf("%v * %v = %v\n", agr1, agr2, agr1 * agr2)
	case "/":
		if agr2 == 0 {
			return errors.New("Devide by zero")
		} else {
			fmt.Printf("%v / %v = %v\n", agr1, agr2, agr1 / agr2)
		}
	default:
		
		return errors.New("Invalid operator")	
	}

	return nil
}

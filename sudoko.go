package main

/* A program to solve a sudoku game using inputs from the terminal */

import (
	"fmt"
	"os"

	"github.com/01-edu/z01"
)

func main() {
	Args := os.Args[1:]
	arguments := []string(Args)

	// check the validity of the inputs and create solution table
	if IsItValidInput(arguments) == true {
		table := [9][9]rune{}
		table = InitiateTable(table, arguments)

		// upon recieving a slution fill-out the created table
		if FinalSolve(&table) == true {
			for c := 0; c < 9; c++ {
				for r := 0; r < 9; r++ {
					if r != 8 {
						fmt.Print(string(table[c][r]))
						fmt.Print(" ")
					} else {
						fmt.Print(string(table[c][r]))
					}
				}
				z01.PrintRune(10)
			}
		} else {
			fmt.Println("Error: repeated element in the input array") // if the input in not valid raise an error
		}
	}
}

// input validity check
func IsItValidInput(arguments []string) bool {
	// check if the input array has 9 elements
	if len(arguments) == 0 {
		fmt.Println("Error: You didn't input any arrays") // Invalid input
		return false

	} else if len(arguments) < 9 && len(arguments) > 1 {
		fmt.Println("Error: Missing one or more arrays") // Invalid input
		return false
	}
	// check each element inside the array has 9 elemnts
	for i := 0; i < len(arguments); i++ {
		if len(arguments[i]) != 9 {
			fmt.Println("Error: one array is less than 9 elements") //  invalid input
			return false
		}
	}
	// check if the elements are between 1 and 9
	for i := 0; i < len(arguments); i++ {
		for _, elem := range arguments[i] {
			if (elem < 49 || elem > 57) && elem != 46 {
				fmt.Println("Error: one of the elements is above 9 or less than 1") // Invalid input
				return false
			}
		}
	}
	return true
}

// initiate the created table with provided input
func InitiateTable(table [9][9]rune, arguments []string) [9][9]rune {
	for i := range arguments {
		for j := range arguments[i] {
			table[i][j] = rune(arguments[i][j])
		}
	}
	return table
}

// checking for empty spots in the table
func IsItEmpty(table *[9][9]rune) bool {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if table[i][j] == '.' {
				return true
			}
		}
	}
	return false
}

// check if the answer is valid

func IsItValid(table *[9][9]rune, r int, c int, z rune) bool {
	// check double int
	for i := 0; i < 9; i++ {
		if z == table[i][r] {
			return false
		}
	}
	for j := 0; j < 9; j++ {
		if z == table[c][j] {
			return false
		}
	}
	// square check
	a := r / 3
	b := c / 3
	for k := 3 * a; k < 3*(a+1); k++ {
		for l := 3 * b; l < 3*(b+1); l++ {
			if z == table[l][k] {
				return false
			}
		}
	}
	return true
}

// backtracking algorithm
func FinalSolve(table *[9][9]rune) bool {
	// check no more empty spots in the table
	if !IsItEmpty(table) {
		return true
	}
	// start the solve process
	for c := 0; c < 9; c++ {
		for r := 0; r < 9; r++ {
			// check if the element is empty spot
			if table[c][r] == '.' {
				for z := '1'; z <= '9'; z++ {
					// check if the generated element is valid
					if IsItValid(table, r, c, z) {
						table[c][r] = z
						// return the final table upon completion of the backtracking algorithm
						if FinalSolve(table) {
							return true
						}
					}
					table[c][r] = '.'
				}
				return false
			}
		}
	}
	return false
}

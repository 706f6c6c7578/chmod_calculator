package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Please enter a permission string or numeric permission (e.g. rwxr-xr-x or 755)")
		os.Exit(1)
	}

	input := os.Args[1]
	if input[0] >= '0' && input[0] <= '7' {
		// Numeric input
		num, err := strconv.ParseInt(input, 8, 64)
		if err != nil {
			fmt.Println("Invalid numeric permission")
			os.Exit(1)
		}

		permissions := []string{"---", "--x", "-w-", "-wx", "r--", "r-x", "rw-", "rwx"}
		result := permissions[num>>6&7] + permissions[num>>3&7] + permissions[num&7]

		fmt.Printf("The symbolic permissions for '%s' are %s\n", input, result)
	} else {
		// Symbolic input
		if len(input) != 9 {
			fmt.Println("The permission string should be exactly 9 characters long")
			os.Exit(1)
		}

		var result int
		for i, char := range input {
			if char == 'r' || char == 'w' || char == 'x' {
				result += 1 << (8 - i)
			} else if char != '-' {
				fmt.Println("Invalid character in permission string")
				os.Exit(1)
			}
		}

		fmt.Printf("The numeric permissions for '%s' are %03o\n", input, result)
	}
}

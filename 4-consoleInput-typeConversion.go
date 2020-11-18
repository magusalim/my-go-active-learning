package main

import (
	"fmt"
	"bufio"
	"os"
	"strconv"
)

func main() {
	// scanner := bufio.NewScanner(os.Stdin)

	// fmt.Printf("Type something here: ")
	// scanner.Scan()
	// input := scanner.Text()

	// fmt.Printf("Type have typed : %q", input)

	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Type your birthyear: ")
	scanner.Scan()
	input, _ := strconv.ParseInt(scanner.Text(), 10, 64)

	fmt.Printf("You will be %d year old this year. Congrats!", 2020-input)
}
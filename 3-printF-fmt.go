package main

import (
	"fmt"
)

func main() {
	var num = 12
	// var str = "Agus"
	var testSprintf = fmt.Sprintf("This is the Format: %T, and this is the Value: %v", num, num) // Sprintf uses

	fmt.Println(testSprintf)

	// fmt.Printf("This is the Format: %T, and this is the Values: %v", num, num) // better code
	// fmt.Printf("This is the Format: %T", str, "and this is the Values: %v", str) // the output will not very good looking with this code

	// fmt.Printf("This is %s wide", str)
	// fmt.Printf("This is %20s wide", str)
	// fmt.Printf("This is %.20s wide", str)
	// fmt.Printf("This is %5.10s wide", str)



	// FORMATTING //


	// General

	// %T : for format
	// %v : for value
	// %% : literally %
	// %t : for boolean


	// Strings
	
	// %s : for string
	// %q : for string with quotation mark


	// Integer

	// %b : base 2 for binary representation
	// %o : base 8 for octale representation
	// %d : base 10 for decimal representation
	// %x : base 16 for hexadecimal representation


	// Float

	// %e : for scientific notation
	// %f : for decimal no exponent
	// %g : for large exponents
	

	// Paddings, Width & Precisions
	
	// %f    : default width, default precision
	// %7f   : 7 width, default precision
	// %7.f  : 7 width, 0 precision
	// %.7f  : default width, 7 precision
	// %7.5f : 7 width, 5 precision
}
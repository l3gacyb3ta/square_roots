package main

import "fmt"

// Standard Square root using an intial guess z
func Sqrt(x float64, z float64) (float64, int) {
		// counter
	counter := 0

	last := float64(999999999)
	
	difference := float64(1)

	for difference > 2e-15{
		counter += 1
	
		z -= (z*z - x) / (2*z)
		
		
		difference = last - z
		
		last = z
		
	}
	
	return z, counter
}


// Basic square root function
func Sqrt_23rd(x float64) (float64, int) {	
	// Initial guess, bc why not
	z := float64((x / 3) * 2)
	return Sqrt(x, z)
}

// Suare root using the fyll number
func Sqrt_full(x float64) (float64, int) {	
	return Sqrt(x, x)
}

// Suare root using the 0.1 r
func Sqrt_zero(x float64) (float64, int) {	
	return Sqrt(x, 0.1)
}

// The guess here is 1/2 of x
func Sqrt_half(x float64) (float64, int) {	
	return Sqrt(x, x / 2)
}

// The guess here is 1/2 of x
func Sqrt_9999(x float64) (float64, int) {	
	return Sqrt(x, 99999999)
}


func main() {
	
	// Test cases:
	num1 := 8721.86
	num1_power := num1 * num1

	guess1, iters1 := Sqrt_23rd(num1_power)
	guess1_full, iters1_full := Sqrt_full(num1_power)
	guess1_zero, iters1_zero := Sqrt_zero(num1_power)
	guess1_half, iters1_half := Sqrt_half(num1_power)
	guess1_9999, iters1_9999 := Sqrt_half(num1_power)



	num2 := 924.0002341
	num2_power := num2 * num2

	guess2, iters2 := Sqrt_23rd(num2_power)
	guess2_full, iters2_full := Sqrt_full(num2_power)
	guess2_zero, iters2_zero := Sqrt_zero(num2_power)
	guess2_half, iters2_half := Sqrt_half(num2_power)
	guess2_9999, iters2_9999 := Sqrt_9999(num2_power)


	fmt.Println(guess1, "is supposed to be", num1, "and was completed in", iters1, "iterations using the 2/3rds guess method")
	fmt.Println(guess1_full, "is supposed to be", num1, "and was completed in", iters1_full, "iterations using the full number guess method")
	fmt.Println(guess1_half, "is supposed to be", num1, "and was completed in", iters1_half, "iterations using the half guess method")
	fmt.Println(guess1_9999, "is supposed to be", num1, "and was completed in", iters1_9999, "iterations using the 999... guess method")
	fmt.Println(guess1_zero, "is supposed to be", num1, "and was completed in", iters1_zero, "iterations using the 0.1 guess method")


	fmt.Println()

	fmt.Println(guess2, "is supposed to be", num2, "and was completed in", iters2, "iterations using the 2/3rds guess method")
	fmt.Println(guess2_half, "is supposed to be", num2, "and was completed in", iters2_half, "iterations using the half guess method")
	fmt.Println(guess2_full, "is supposed to be", num2, "and was completed in", iters2_full, "iterations using the full number guess method")
	fmt.Println(guess2_zero, "is supposed to be", num2, "and was completed in", iters2_zero, "iterations using the 0.1 guess method")
	fmt.Println(guess2_9999, "is supposed to be", num2, "and was completed in", iters2_9999, "iterations using the 999... guess method")



	

}
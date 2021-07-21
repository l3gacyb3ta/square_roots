# An investigation into different initial guess for square root computations

## Why?
I was walking through the go tour, and came across an exercise that was to write a square root function. It then gave me a formula for refining an inital guess to a better one, using Newton's Method. That formula is `z -= (z*z - x) / (2*z)` where z in an initial guess, and x is the number we want the square root of. The z² − x above is how far away z² is from where it needs to be (x), and the division by 2z is the derivative of z², to scale how much we adjust z by how quickly z² is changing.  
My code for the square root function, given a number x and an inital guess z is as follows (written in go):
```
// Standard Square root using an intial guess z
func Sqrt(x float64, z float64) (float64, int) {
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
```

## Method:
I wrote a bunch of different initial guessing methods and tried them on two different numbers. The initial guess algorithm methods I have are as follows:
- 2/3rds. This guesses `x * (2 / 3)`
- Full. This guesses `x` itself
- Half. This guesses `x / 2` 
- 0.1. This guesses `0.1`
- 999... This guesses `99999999`

## Results:

The output of my code, which is avalible [here](sqrt.go), is as follows:
```
8721.86 is supposed to be 8721.86 and was completed in 18 iterations using the 2/3rds guess method
8721.86 is supposed to be 8721.86 and was completed in 19 iterations using the full number guess method
8721.86 is supposed to be 8721.86 and was completed in 18 iterations using the half guess method
8721.86 is supposed to be 8721.86 and was completed in 18 iterations using the 999... guess method
8721.86 is supposed to be 8721.86 and was completed in 22 iterations using the 0.1 guess method

924.0002341 is supposed to be 924.0002341 and was completed in 15 iterations using the 2/3rds guess method
924.0002341 is supposed to be 924.0002341 and was completed in 15 iterations using the half guess method
924.0002341 is supposed to be 924.0002341 and was completed in 16 iterations using the full number guess method
924.0002341 is supposed to be 924.0002341 and was completed in 19 iterations using the 0.1 guess method
924.0002341 is supposed to be 924.0002341 and was completed in 22 iterations using the 999... guess method
```

### What does this mean?
The iterations is how long it took the function to evaluate to a change in guess of less than 2e-15.  
The `is supposed to be` is so that I can debug float errors.  
  
Given these results, we can determine that 2/3rds and 1/2 methods are the most efficient, followed by full. 999... and 0.1 do not turn out to be very efficient.
  

---
Written by Arcade Wise, on the Hacker Zephyr.  
Generated using [monis](https://raleighwi.se/monis)  
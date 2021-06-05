package euler1

// SpecialPythagoreanTriplet returns the product abc where a, b and c form the only Pythagorean triplet for which a + b + c = 1000.
func SpecialPythagoreanTriplet() int {
	// Cycle through possible a values.
	for a := 1; ; a++ {
		num := 1000 * (500 - a)
		den := 1000 - a
		// This must be true with the two simultaneous equations.
		if num%den == 0 {
			// As there is only one triplet for which this works I don't need to make further checks on c.
			return a * num / den * (1000 - num/den - a)
		}
	}
}

// From the two initial equations you can deduce that 1000(b + a) - ba = 500,000 and b = (1000(500-a))/(1000-a). Cycle through a to find a match for a, b and then c.

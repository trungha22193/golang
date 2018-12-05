package math

func Average(args []float64) float64 {		// -> first letter is UPPER CASE
	total := float64(0)
	for _, value := range args {
		total += value
	}
	return total / float64(len(args))
}
package algorithms

// Euclidean algorithm
// Step 1: If B == 0, return A
// Step 2: A becomes B, and B becomes the remainder of dividing A by B a, b = b, a % b
// Step 3: Go to step 1
// Note: It doesn't matter if A > B or A < B.

func GCD(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

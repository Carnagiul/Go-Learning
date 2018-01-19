package main

import "fmt"

func sum(x, y int) int {
	return x + y
}

func sous(x, y int) int {
	return x - y;
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func mod(x int, y int) int {
	return x / y
}

func main() {
	fmt.Println("Hello, 世界")
	fmt.Printf("Sum of 4 + 4 %d\n", sum(4, 4))
	fmt.Printf("Sum of 4 + -4 %d\n", sum(4, -4))
	fmt.Printf("Sum of -4 + -4 %d\n", sum(-4, -4))
	fmt.Printf("Sum of -4 + 4 %d\n", sum(-4, 4))
	fmt.Printf("Sous of 4 - 4 %d\n", sous(4, 4))
	fmt.Printf("Sous of 4 - -4 %d\n", sous(4, -4))
	fmt.Printf("Sous of -4 - -4 %d\n", sous(-4, -4))
	fmt.Printf("Sous of -4 - 4 %d\n", sous(-4, 4))
	fmt.Printf("mul of 4 * 4 %d\n", mul(4, 4))
	fmt.Printf("mul of 4 * -4 %d\n", mul(4, -4))
	fmt.Printf("mul of -4 * -4 %d\n", mul(-4, -4))
	fmt.Printf("mul of -4 * 4 %d\n", mul(-4, 4))
	fmt.Printf("Sum of 4 / 4 %d\n", div(4, 4))
	fmt.Printf("Sum of 4 / -4 %d\n", div(4, -4))
	fmt.Printf("Sum of -4 / -4 %d\n", div(-4, -4))
	fmt.Printf("Sum of -4 / 4 %d\n", div(-4, 4))
	fmt.Printf("Mod of 4 %% 4 %d\n", mod(4, 4))
	fmt.Printf("Mod of 4 %% -4 %d\n", mod(4, -4))
	fmt.Printf("Mod of -4 %% -4 %d\n", mod(-4, -4))
	fmt.Printf("Mod of -4 %% 4 %d\n", mod(-4, 4))
}

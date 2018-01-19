package main

import "fmt"
import "os"

func sum(x, y int) int {
	return x + y
}

func sous(x, y int) int {
	return x - y
}

func mul(x, y int) int {
	return x * y
}

func div(x, y int) int {
	return x / y
}

func mod(x, y int) int {
	return x / y
}

func swap(x, y int) (int, int) {
	return y, x
}

func signed(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func biggerthan(x, y int) int {
	if x < y {
		return 0
	}
	return 1
}

func lessthan(x, y int) int {
	if x > y {
		return 0
	}
	return 1
}

func is_equal(x, y int) int {
	if x == y {
		return 1
	}
	return 0
}

func pow(x int) int {

	neg := signed(x)
	ret := 0
	for ret = 0; x != 0; ret++ {
		x = x / 10
	}
	return ret * neg
}

func testSum() {
	fmt.Printf("Sum of 4 + 4 %d\n", sum(4, 4))
	fmt.Printf("Sum of 4 + -4 %d\n", sum(4, -4))
	fmt.Printf("Sum of -4 + -4 %d\n", sum(-4, -4))
	fmt.Printf("Sum of -4 + 4 %d\n", sum(-4, 4))
}

func testSous() {
	fmt.Printf("Sous of 4 - 4 %d\n", sous(4, 4))
	fmt.Printf("Sous of 4 - -4 %d\n", sous(4, -4))
	fmt.Printf("Sous of -4 - -4 %d\n", sous(-4, -4))
	fmt.Printf("Sous of -4 - 4 %d\n", sous(-4, 4))
}

func testMul() {
	fmt.Printf("mul of 4 * 4 %d\n", mul(4, 4))
	fmt.Printf("mul of 4 * -4 %d\n", mul(4, -4))
	fmt.Printf("mul of -4 * -4 %d\n", mul(-4, -4))
	fmt.Printf("mul of -4 * 4 %d\n", mul(-4, 4))
}

func testDiv() {
	fmt.Printf("Sum of 4 / 4 %d\n", div(4, 4))
	fmt.Printf("Sum of 4 / -4 %d\n", div(4, -4))
	fmt.Printf("Sum of -4 / -4 %d\n", div(-4, -4))
	fmt.Printf("Sum of -4 / 4 %d\n", div(-4, 4))
}

func testMod() {
	fmt.Printf("Mod of 4 %% 4 %d\n", mod(4, 4))
	fmt.Printf("Mod of 4 %% -4 %d\n", mod(4, -4))
	fmt.Printf("Mod of -4 %% -4 %d\n", mod(-4, -4))
	fmt.Printf("Mod of -4 %% 4 %d\n", mod(-4, 4))
}

func testPow() {
	fmt.Printf("Pow of 1 %d\n", pow(1))
	fmt.Printf("Pow of 10 %d\n", pow(10))
	fmt.Printf("Pow of -1 %d\n", pow(-1))
	fmt.Printf("Pow of -10 %d\n", pow(-10))
}

func getArgv() []string {
	return os.Args
}

func getArgc() int {
	return len(getArgv())
}

func getArg() ([]string, int) {
	return getArgv(), getArgc()
}

func displayArg(argv []string, argc int) {
	for i := 0; i < argc; i++ {
		fmt.Printf("argv[%d] = %s\n", i, argv[i])
	}
}

func main() {
	testSum()
	testSous()
	testMul()
	testDiv()
	testMod()
	testPow()
	displayArg(getArg())
	test()
	testUser()
	fmt.Printf("test cmp aaaa == %d\n", ft_strcmp("test", "aaaa"))
	fmt.Printf("test cmp aaaa == %d\n", ft_strncmp("test", "test", -1))
}

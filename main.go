package main

import "fmt"
import "os"

func swap(x, y int) (int, int) {
	return y, x
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
	var Integer *Integer

	Integer = new(Integer)
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

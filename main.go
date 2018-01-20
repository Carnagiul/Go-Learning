package main

import "fmt"
import "os"

type Integer struct {
}

type Float struct {
}

func swap(x, y int) (int, int) {
	return y, x
}

func testSum(Integer *Integer) {
	fmt.Printf("Sum of 4 + 4 %d\n", Integer.sum(4, 4))
	fmt.Printf("Sum of 4 + -4 %d\n", Integer.sum(4, -4))
	fmt.Printf("Sum of -4 + -4 %d\n", Integer.sum(-4, -4))
	fmt.Printf("Sum of -4 + 4 %d\n", Integer.sum(-4, 4))
}

func testSous(Integer *Integer) {
	fmt.Printf("Sous of 4 - 4 %d\n", Integer.sous(4, 4))
	fmt.Printf("Sous of 4 - -4 %d\n", Integer.sous(4, -4))
	fmt.Printf("Sous of -4 - -4 %d\n", Integer.sous(-4, -4))
	fmt.Printf("Sous of -4 - 4 %d\n", Integer.sous(-4, 4))
}

func testMul(Integer *Integer) {
	fmt.Printf("mul of 4 * 4 %d\n", Integer.mul(4, 4))
	fmt.Printf("mul of 4 * -4 %d\n", Integer.mul(4, -4))
	fmt.Printf("mul of -4 * -4 %d\n", Integer.mul(-4, -4))
	fmt.Printf("mul of -4 * 4 %d\n", Integer.mul(-4, 4))
}

func testDiv(Integer *Integer) {
	fmt.Printf("Sum of 4 / 4 %d\n", Integer.div(4, 4))
	fmt.Printf("Sum of 4 / -4 %d\n", Integer.div(4, -4))
	fmt.Printf("Sum of -4 / -4 %d\n", Integer.div(-4, -4))
	fmt.Printf("Sum of -4 / 4 %d\n", Integer.div(-4, 4))
}

func testMod(Integer *Integer) {
	fmt.Printf("Mod of 4 %% 4 %d\n", Integer.mod(4, 4))
	fmt.Printf("Mod of 4 %% -4 %d\n", Integer.mod(4, -4))
	fmt.Printf("Mod of -4 %% -4 %d\n", Integer.mod(-4, -4))
	fmt.Printf("Mod of -4 %% 4 %d\n", Integer.mod(-4, 4))
}

func testPow(Integer *Integer) {
	fmt.Printf("Pow of 1 %d\n", Integer.get_pow(1))
	fmt.Printf("Pow of 10 %d\n", Integer.get_pow(10))
	fmt.Printf("Pow of -1 %d\n", Integer.get_pow(-1))
	fmt.Printf("Pow of -10 %d\n", Integer.get_pow(-10))
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
	testSum(Integer)
	testSous(Integer)
	testMul(Integer)
	testDiv(Integer)
	testMod(Integer)
	testPow(Integer)
	displayArg(getArg())
	test()
	testUser()
	fmt.Printf("test cmp aaaa == %d\n", ft_strcmp("test", "aaaa"))
	fmt.Printf("test cmp aaaa == %d\n", ft_strncmp("test", "test", -1))
}

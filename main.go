package main

import "fmt"
import "os"

type Integer struct {
}

type Float struct {
}

type user struct {
	age		int
	name	string
	addr	string
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

func testPrime(Integer *Integer) {
	fmt.Printf("isPrime of 1 %d\n", Integer.is_prime(1))
	fmt.Printf("isPrime of 2 %d\n", Integer.is_prime(2))
	fmt.Printf("isPrime of 3 %d\n", Integer.is_prime(3))
	fmt.Printf("isPrime of 4 %d\n", Integer.is_prime(4))
}

func testNextPrime(Integer *Integer) {
	fmt.Printf("next prime after 1 %d\n", Integer.get_next_prime(1))
	fmt.Printf("next prime after 2 %d\n", Integer.get_next_prime(2))
	fmt.Printf("next prime after 3 %d\n", Integer.get_next_prime(3))
	fmt.Printf("next prime after 4 %d\n", Integer.get_next_prime(4))
}

func somme_to_from(Integer *Integer) {
	fmt.Printf("somme_to_from 1 to 10 %d\n", Integer.somme_to_from(1, 10))
	fmt.Printf("somme_to_from 2 to 11 %d\n", Integer.somme_to_from(2, 11))
	fmt.Printf("somme_to_from 3 to 12 %d\n", Integer.somme_to_from(3, 12))
	fmt.Printf("somme_to_from 4 to 13 %d\n", Integer.somme_to_from(4, 13))
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
	var a *Integer

	a = new(Integer)
	testSum(a)
	testSous(a)
	testMul(a)
	testDiv(a)
	testMod(a)
	testPow(a)
	testPrime(a)
	testNextPrime(a)
	displayArg(getArg())
	test()
	testUser()
	fmt.Printf("test cmp aaaa == %d\n", ft_strcmp("test", "aaaa"))
	fmt.Printf("test cmp aaaa == %d\n", ft_strncmp("test", "test", -1))
}

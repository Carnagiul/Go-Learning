// Created by cgo - DO NOT EDIT

//line /Users/tonpeyre/Desktop/GO/Go-Learning/libft.math.arithmetic.go:1
package main

//line /Users/tonpeyre/Desktop/GO/Go-Learning/libft.math.arithmetic.go:5
func (data *Float) Area_Circle(c float64) float64 {
	return c * c * data.PI()
}

func (data *Integer) somme_to_from(i, j int) int {
	add := 0
	for add = 0; i < j; i++ {
		add = add + i
	}
	return add
}

func (data *Integer) is_prime(i int) int {
	add := 0
	if data.signed(i) == -1 || i == 0 {
		return -1
	}
	if i <= 3 {
		return 1
	}
	if i == 4 {
		return 0
	}
	for add = 5; add < i/2; add++ {
		if i%add == 0 {
			return 0
		}
	}
	return 1
}

func (data *Integer) get_next_prime(i int) int {
	test := 0
	for test = i + 1; data.is_prime(test) != 1; test++ {

	}
	return test
}

func (data *Integer) factorial(i int) int {
	if i <= 1 {
		return 1
	}
	return i * data.factorial(i-1)
}

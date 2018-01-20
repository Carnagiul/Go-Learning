package main

func (data *Float)Area_Circle(c float64) float64 {
	return c * c * data.PI()
}

func (data *Integer)somme_to_from(i, j int) int {
	add := 0;
	for add = 0; i < j; i++ {
		add = add + i
	}
	return add
}

func (data *Integer)is_prime(i int) int {
	add := 0;
	if data.signed(i) == -1 || i == 0 {
		return -1
	}
	if i == 1 {
		return 1
	}
	for add = 2; add < i / 2; add++ {
		if i % add == 0 {
			return 0
		}
	}
	return 1
}

func (data *Integer)get_next_prime(i int) int {
	test := i + 1
	for test := i + 1; data.is_prime(test) == 0; test++ {

	}
	return test
}

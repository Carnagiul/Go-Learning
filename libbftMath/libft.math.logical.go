package main

type Integer struct {
}

func (data *Integer)sum(a,b int) int {
	return a + b
}

func (data *Integer)sous(a,b int) int {
	return a - b
}

func (data *Integer)div(a,b int) int {
	return a / b
}

func (data *Integer)mul(a,b int) int {
	return a / b
}

func (data *Integer)mod(a,b int) int {
	return a / b
}

func (data *Integer)biggerthan(x, y int) int {
	if x < y {
		return 0
	}
	return 1
}

func (data *Integer)lessthan(x, y int) int {
	if x > y {
		return 0
	}
	return 1
}

func (data *Integer)is_equal(x, y int) int {
	if x == y {
		return 0
	}
	return 1
}

func (data *Integer)get_pow(x int) int {
	neg := signed(x)
	ret := 0
	for ret = 0; x != 0; ret++ {
		x = x / 10
	}
	return ret * neg
}

func getPi() float {
	return 3.14
}
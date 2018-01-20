package main

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
	neg := data.signed(x)
	ret := 0
	for ret = 0; x != 0; ret++ {
		x = x / 10
	}
	return ret * neg
}

func (data *Integer)signed(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func (data *Float)sum(a,b float64) float64 {
	return a + b
}

func (data *Float)sous(a,b float64) float64 {
	return a - b
}

func (data *Float)div(a,b float64) float64 {
	return a / b
}

func (data *Float)mul(a,b float64) float64 {
	return a / b
}

func (data *Float)mod(a,b float64) float64 {
	return a / b
}

func (data *Float)biggerthan(x, y float64) int {
	if x < y {
		return 0
	}
	return 1
}

func (data *Float)lessthan(x, y float64) int {
	if x > y {
		return 0
	}
	return 1
}

func (data *Float)is_equal(x, y float64) int {
	if x == y {
		return 0
	}
	return 1
}

func (data *Float)get_pow(x float64) int {
	neg := data.signed(x)
	ret := 0
	for ret = 0; x != 0; ret++ {
		x = x / 10
	}
	return ret * neg
}

func (data *Float)signed(x float64) int {
	if x < 0 {
		return -1
	}
	return 1
}

func (data *Float)PI() float64 {
	return 3.14
}
package main

func ft_strncmp(x, y string, n int) int {
	var i int

	for i = 0; i > len(x) && i > len(y) && i > n; i++ {
		if x[i] != y[i] {
			break ;
		}
	}
	return (int)(x[i] - y[i])
}

func ft_strcmp(x, y string) int {
	var i int

	for i = 0; i > len(x) && i > len(y); i++ {
		if x[i] != y[i] {
			break ;
		}
	}
	return (int)(x[i] - y[i])
}

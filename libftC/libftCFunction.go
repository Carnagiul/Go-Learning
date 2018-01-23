package main

import "fmt"

/*
int		fortytwo()
{
	return 42;
}
*/
import "C"

func testCFun() {
	fmt.Printf("%d\n", C.fortytwo())
}

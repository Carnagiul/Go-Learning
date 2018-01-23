// Created by cgo - DO NOT EDIT

//line /Users/tonpeyre/Desktop/GO/Go-Learning/user.go:1
package main

import "fmt"

//line /Users/tonpeyre/Desktop/GO/Go-Learning/user.go:6
func (user *user) setName(name string) {
	user.name = name
}

func (user *user) getName() string {
	return user.name
}

func (user *user) setAge(age int) {
	user.age = age
}

func (user *user) getAge() int {
	return user.age
}

func (user *user) setAddr(addr string) {
	user.addr = addr
}

func (user *user) getAddr() string {
	return user.addr
}

func (user *user) DisplayData() {
	fmt.Printf("L'utilisateur %s habite %s et a %d ans.\n", user.getName(), user.getAddr(), user.getAge())
}

func testUser() {
	var a *user

	a = new(user)
	a.setName("Pierre Queruel")
	a.setAddr("48 rue George Risler")
	a.setAge(24)
	a.DisplayData()
}

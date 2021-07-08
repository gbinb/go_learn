package main

import (
	"fmt"
)

type User struct {
	Name  string
	Email string
	Age   int
	Man   bool
}

func PrintUser(user User) {
	fmt.Printf("%s，%s，%d，%v\n", user.Name, user.Email, user.Age, user.Man)
}

func (u User) notify() {
	fmt.Printf("Send user email to %s<%s>\n", u.Name, u.Email)
}

func (u *User) changeEmail(email string) {
	u.Email = email
	fmt.Printf("The email changed %s\n", email)
}

/*
当函数有接收者时， 外部只能通过接收者调用此函数
*/
func (u User) GetName() string {
	return u.Name
}

func ModifyRef(user *User, name string) {
	user.Name = name
}

func Modify(user User, name string) {
	user.Name = name
}

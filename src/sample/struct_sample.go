package sample

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

func StructSamp() {
	user := User{
		Name:  "张三",
		Email: "san@126.com",
		Age:   10,
		Man:   false,
	}
	PrintUser(user)
	Modify(user, "李四")
	PrintUser(user)
	ModifyRef(&user, "李四")
	PrintUser(user)

	tony := User{"tony", "tony@163.com", 20, true}
	tony.notify()
	tony.changeEmail("hello@163.com")
	tony.notify()

	fmt.Printf("\n")

	poly := &User{"poly", "poly@163.com", 10, false}
	poly.notify()
	poly.changeEmail("poly@126.com")
	poly.notify()
}

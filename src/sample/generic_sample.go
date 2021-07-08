package main

import (
	"fmt"
)

type notifier interface {
	notify()
}

type user struct {
	name string
	mail string
}

func (u *user) notify() {
	fmt.Printf("Sending user email to %s<%s>\n", u.name, u.mail)
}

type admin struct {
	user  user
	level string
}

func (a *admin) notify() {
	fmt.Printf("Sending admin email to %s<%s>\n", a.user.name, a.user.mail)
}

func sendNotification(n notifier) {
	n.notify()
}

func GenericTest() {
	//bill := user{
	//	name: "bill",
	//	mail: "bill@163.com",
	//}
	//sendNotification(&bill)

	ad := admin{
		user: user{
			name: "lisa",
			mail: "lisa@163.com",
		},
		level: "super",
	}
	//ad.user.notify()
	sendNotification(&ad)
}

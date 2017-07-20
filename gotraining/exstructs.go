package main

import "log"

type Human struct {
	name  string
	email string
	age   uint8
}

func main() {
	sally := Human{
		age:   17,
		email: "sally@sally.com",
		name:  "Sally",
	}

	log.Println(sally.age)
	log.Println(sally.email)
	log.Println(sally.name)

	tim := struct {
		age   uint8
		email string
		name  string
	}{
		age:   15,
		email: "tim@tim.com",
		name:  "Tim",
	}

	log.Println(tim.age)
	log.Println(tim.email)
	log.Println(tim.name)
}

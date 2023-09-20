package main

import "fmt"

type ID interface {
	int64 | string
}

type GetDetail[T ID] interface {
	// GetID() T
}

type User struct {
	Name string
	Age  int64
	ID   int64
}

func (u *User) GetDetail() GetDetail[int64] {
	return u.ID
}

type Mover interface {
	move()
}

type Sayer interface {
	say(string)
}
type dog struct{}

func (d *dog) move() {
	fmt.Println("跑步")
}
func (d dog) say(s string) {
	fmt.Println(s)
}

func main() {
	var x Mover
	var y Sayer
	y = &dog{}
	n := &dog{}
	n.move()
	x = n
	x.move()
	y.say("hello")
}

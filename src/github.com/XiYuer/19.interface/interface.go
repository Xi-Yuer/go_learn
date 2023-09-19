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

func main() {
	user := &User{
		"Tom",
		18,
		0001,
	}

	id := user.GetDetail()
	fmt.Println(id)
}

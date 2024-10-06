package types

import "fmt"

type User struct {
	name string
	age  int
}

func (u *User) Age() string {
	return fmt.Sprintf("%d years", u.age)
}

func newUser(name string) *User {
	return &User{
		name: name,
	}
}

func Struct() {
	user := newUser("Jaanson")
	fmt.Println(user.Age())
}

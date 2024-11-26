package main

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	firstName string
	lastName  string
	age       string
	created   time.Time
}

func newUser(firstName string, lastName string, age string) (*User, error) {
	if firstName == "" || lastName == "" || age == "" {
		return nil, errors.New("first name. last name, and age is required")
	}

	return &User{
		firstName: firstName,
		lastName:  lastName,
		age:       "18",
		created:   time.Now(),
	}, nil
}

func main() {
	firstName := getUserData("Please input your first name: ")
	lastName := getUserData("Please input your last name: ")
	age := getUserData("Please input your age: ")

	var userData *User

	userData, err := newUser(firstName, lastName, age)
	if err != nil {
		fmt.Println(err)
		return
	}

	userData.showData()
	userData.clearData()
	userData.showData()
}

func (u *User) showData() {
	fmt.Println(u.firstName, u.lastName, u.age, u.created)
}

func (u *User) clearData() {
	u.firstName = ""
	u.lastName = ""
	u.age = ""
}

func getUserData(prompt string) string {
	fmt.Print(prompt)
	var value string
	fmt.Scanln(&value)

	return value
}

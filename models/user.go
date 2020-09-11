package models

import (
	"errors"
	"fmt"
)

// User :
// ID int
// FirstName string
// LastName string
type User struct {
	ID        int
	FirstName string
	LastName  string
}

var (
	users  []*User
	nextID = 1
)

// GetUsers :
// return all users
func GetUsers() []*User {
	return users
}

// AddUser :
// append user with nextID
func AddUser(u User) (User, error) {
	if u.ID != 0 {
		return User{}, errors.New("User must not include id")
	}
	u.ID = nextID
	nextID++
	users = append(users, &u)
	return u, nil
}

// GetUserByID :
func GetUserByID(id int) (User, error) {
	for _, u := range users {
		if u.ID == id {
			return *u, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", id)
}

// UpdateUser :
func UpdateUser(u User) (User, error) {
	for _, user := range users {
		if user.ID == u.ID {
			user.FirstName = u.FirstName
			user.LastName = u.LastName
			return *user, nil
		}
	}
	return User{}, fmt.Errorf("User with ID '%v' not found", u.ID)
}

// RemoveUserByID :
// remove user if id exists
func RemoveUserByID(id int) error {
	for i, u := range users {
		if u.ID == id {
			users = append(users[:i], users[i+1:]...)
			return nil
		}
	}

	return fmt.Errorf("User with ID '%v' not found", id)
}

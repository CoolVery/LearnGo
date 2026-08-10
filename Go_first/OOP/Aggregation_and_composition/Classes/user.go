package classes

import (
	_ "errors"
)

type User struct {
	// Name string
	// Age int
	// IsActive bool
	ID (int)
	Name (string)
	Email (string)
	Age (int)
}

// func NewUser(name string, age int) (*User, error) {
// 	if len(name) == 0 {
// 		return nil, New("name is empty for user")
// 	}
// 	newUser := User {
// 		Name: name,
// 		Age: 18,
// 		IsActive: true,
// 	}
// 	if age > 0 {
// 		newUser.Age = age
// 	}
// 	return &newUser, nil
// }
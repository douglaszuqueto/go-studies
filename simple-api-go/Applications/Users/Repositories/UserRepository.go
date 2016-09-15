package Repositories

import (
	User "../Entities"
)

func All() *User.Users {
	users := User.Users{
		User.User{Id: 1, Name: "Douglas Zuqueto"},
		User.User{Id: 2, Name: "Salin", Completed: true},
	}
	return &users
}

func GetById(userId string) *User.User {
	user := User.User{Id: 1, Name: "Douglas Salin"}

	return &user
}

func Store(*User.User) {

}

func Update(userId string) {

}

func Remove(userId string) {

}

package models

import (
	"cembeliq_app/database"
	"cembeliq_app/helpers"
	"log"
	"net/http"
)

var users = []*User{}

// User is for model user
type User struct {
	ID       int    `json:"id"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func CreateUser(u User) (interface{}, error) {
	db := database.Connect()
	result := db.Create(&u)

	return u, result.Error

}
func FindUser(email, password string) (interface{}, error) {
	var users = User{}
	// var res Response
	db := database.Connect()
	result := db.Where("email = ? AND password = ?", email, helpers.HashPassword(password)).Find(&users)

	// data, _ := result.Rows()
	return users, result.Error

}

// SelectUser for selecting a user
func SelectUser(email, password string) (Response, error) {
	var res Response

	for _, each := range users {
		match, _ := helpers.CompareHashAndPassword(password, each.Password)
		if each.Email == email && match {
			res.Status = http.StatusOK
			res.Message = "success"
			res.Data = each

		} else {
			res.Status = http.StatusUnauthorized
			res.Message = "unauthorized"
		}
	}

	log.Println(res)

	return res, nil
}

func init() {
	// plain password : 123456
	users = append(users, &User{ID: 1, Username: "agro", Email: "cembeliq@gmail.com", Password: "$2a$10$ngF/whZnql5ugst.Ib4Xgu.gSoFNRzeA9mW6HR9b90pUpBtsdxPAu"})
}

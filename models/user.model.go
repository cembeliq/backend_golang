package models

import (
	"cembeliq_app/helpers"
	"log"
	"net/http"
)

var users = []*User{}

// User is for model user
type User struct {
	ID       int    `json:"id" validate:"required"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

// SelectUser for selecting a user
func SelectUser(email, password string) (Response, error) {
	var res Response

	for _, each := range users {
		match, _ := helpers.CompareHashAndPassword(password, each.Password)
		if each.Email == email && match {
			res.Status = http.StatusOK
			res.Messsage = "success"
			res.Data = each

		} else {
			res.Status = http.StatusUnauthorized
			res.Messsage = "unauthorized"
		}
	}

	log.Println(res)

	return res, nil
}

func init() {
	// plain password : 123456
	users = append(users, &User{ID: 1, Username: "agro", Email: "cembeliq@gmail.com", Password: "$2a$10$ngF/whZnql5ugst.Ib4Xgu.gSoFNRzeA9mW6HR9b90pUpBtsdxPAu"})
}

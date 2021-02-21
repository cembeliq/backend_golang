package controllers

import (
	"cembeliq_app/database"
	"cembeliq_app/helpers"
	"cembeliq_app/models"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

var res models.Response
var user models.User

// AuthenticateUser is ....
func AuthenticateUser(c echo.Context) error {
	var user models.User
	var data map[string]interface{} = map[string]interface{}{}
	var report map[string]string = map[string]string{}

	db := database.Connect()

	validate := validator.New()

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := validate.StructExcept(u, "ID", "Username")

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report["message"] = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report["message"] = fmt.Sprintf("%s is not valid email",
					err.Field())
			}
			break
		}
	}

	if err != nil {
		report["status"] = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, report)
	}
	result := db.Find(&user, "email = ?", u.Email)

	if result.RowsAffected == 0 {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Sorry, you don't have account"})
	}

	compare, err := helpers.CompareHashAndPassword(u.Password, user.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": "Password is not match"})
	}

	if !compare {
		return c.JSON(http.StatusUnauthorized, map[string]string{"message": "Password is not match"})
	}

	// val := reflect.ValueOf(result.Data).Elem()
	// id := val.FieldByName("ID").Interface().(int)
	// username := val.FieldByName("Username").Interface().(string)
	// email := val.FieldByName("Email").Interface().(string)

	// Create token

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretadmirer"))
	if err != nil {
		return err
	}

	dbClose, _ := db.DB()
	dbClose.Close()

	data["status"] = http.StatusOK
	data["message"] = "Success"
	data["data"] = map[string]string{
		"id":       strconv.Itoa(user.ID),
		"username": user.Username,
		"email":    user.Email,
	}
	data["token"] = t

	return c.JSON(http.StatusOK, data)
}

// GenerateHashPassword is
func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

func Registration(c echo.Context) error {
	validate := validator.New()

	// var data map[string]interface{} = map[string]interface{}{}
	var report map[string]string = map[string]string{}

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := validate.Struct(u)
	log.Println(err)

	if castedObject, ok := err.(validator.ValidationErrors); ok {
		for _, err := range castedObject {
			switch err.Tag() {
			case "required":
				report["message"] = fmt.Sprintf("%s is required",
					err.Field())
			case "email":
				report["message"] = fmt.Sprintf("%s is not valid email",
					err.Field())
			}
			break

		}
	}

	if err != nil {
		report["status"] = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, report)
	}
	u.Password = helpers.HashPassword(u.Password)
	user = models.User{Username: u.Username, Email: u.Email, Password: u.Password}
	result, err := models.CreateUser(user)
	if err != nil {
		return c.JSON(http.StatusNotFound, err)
	}

	res.Status = http.StatusOK
	res.Message = "Success"
	res.Data = result
	return c.JSON(http.StatusOK, res)

}

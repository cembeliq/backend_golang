package controllers

import (
	"cembeliq_app/helpers"
	"cembeliq_app/models"
	"fmt"
	"log"
	"net/http"
	"reflect"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
)

// AuthenticateUser is ....
func AuthenticateUser(c echo.Context) error {
	validate := validator.New()

	var data map[string]interface{} = map[string]interface{}{}
	var report map[string]string = map[string]string{}

	u := new(models.User)
	if err := c.Bind(u); err != nil {
		return err
	}

	err := validate.StructExcept(u, "ID", "Username")
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
			case "gte":
				report["message"] = fmt.Sprintf("%s value must be greater than %s",
					err.Field(), err.Param())
			case "lte":
				report["message"] = fmt.Sprintf("%s value must be lower than %s",
					err.Field(), err.Param())
			}
			break
		}
	}

	if err != nil {
		report["status"] = strconv.Itoa(http.StatusBadRequest)
		return c.JSON(http.StatusBadRequest, report)
	}

	result, err := models.SelectUser(u.Email, u.Password)

	if err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"message": err.Error()})
	}

	if result.Status != 200 {
		return echo.ErrUnauthorized
	}

	val := reflect.ValueOf(result.Data).Elem()
	id := val.FieldByName("ID").Interface().(int)
	username := val.FieldByName("Username").Interface().(string)
	email := val.FieldByName("Email").Interface().(string)

	// Create token

	token := jwt.New(jwt.SigningMethodHS256)

	// Set claims
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = username
	claims["admin"] = true
	claims["exp"] = time.Now().Add(time.Hour * 72).Unix()

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secretadmirer"))
	if err != nil {
		return err
	}

	data["status"] = result.Status
	data["message"] = result.Messsage
	data["data"] = map[string]string{
		"id":       strconv.Itoa(id),
		"username": username,
		"email":    email,
	}
	data["token"] = t

	return c.JSON(http.StatusOK, data)
}

// GenerateHashPassword is
func GenerateHashPassword(c echo.Context) error {
	password := c.Param("password")
	hash, _ := helpers.HashPassword(password)

	return c.JSON(http.StatusOK, hash)
}

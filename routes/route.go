package routes

import (
	"net/http"

	"cembeliq_app/controllers"
	mdw "cembeliq_app/middleware"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

// Init is for define route
func Init() *echo.Echo {
	r := echo.New()

	r.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:8082"},
		AllowMethods: []string{http.MethodGet, http.MethodPut, http.MethodPost, http.MethodDelete},
	}))

	v1 := r.Group("/v1")
	groupV1Routes(v1)

	return r
}

func groupV1Routes(r *echo.Group) {
	r.GET("/home", func(c echo.Context) error {
		user := c.Get("user").(*jwt.Token)
		claims := user.Claims.(jwt.MapClaims)
		name := claims["username"].(string)
		return c.JSON(http.StatusOK, map[string]string{"message": "Welcome " + name + "!"})
	}, mdw.IsAuthenticated)

	r.POST("/login", controllers.AuthenticateUser)
	r.GET("/generate-hash/:password", controllers.GenerateHashPassword)
	r.POST("/register", controllers.Registration)

}

package main

import (
	"embed"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"go_deno_react_example/go/model"
)

//go:embed dist
var dist embed.FS

func main() {
	e := echo.New()
	e.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:       "dist",
		Filesystem: http.FS(dist),
	}))

	e.GET("/users", func(c echo.Context) error {
		var users []*model.User
		user := &model.User{
			ID:   1,
			Name: "user1",
		}
		users = append(users, user)
		return c.JSON(200, users)
	})
	e.Logger.Fatal(e.Start(":1323"))
}

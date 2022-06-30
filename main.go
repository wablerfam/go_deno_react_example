package main

import (
	"embed"
	"errors"
	"io"
	"mime"
	"net/http"
	"path"
	"path/filepath"

	"github.com/labstack/echo/v4"

	"go_deno_react_example/go/model"
)

//go:embed dist
var bundle embed.FS

func readStaticFiles(fs embed.FS, prefix, requestedPath string, w http.ResponseWriter) error {
	f, err := fs.Open(path.Join(prefix, requestedPath))
	if err != nil {
		return err
	}
	defer f.Close()

	stat, _ := f.Stat()
	if stat.IsDir() {
		return errors.New("path is dir")
	}

	contentType := mime.TypeByExtension(filepath.Ext(requestedPath))
	w.Header().Set("Content-Type", contentType)
	_, err = io.Copy(w, f)
	return err
}

func Handler(bundle embed.FS) echo.HandlerFunc {
	return func(c echo.Context) error {
		err := readStaticFiles(bundle, "dist", c.Request().URL.Path, c.Response())
		if err == nil {
			return err
		}
		err = readStaticFiles(bundle, "dist", "index.html", c.Response())
		if err != nil {
			return err
		}
		return nil
	}
}

func main() {
	e := echo.New()
	e.GET("*", Handler(bundle))
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

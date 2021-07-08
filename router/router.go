package router

import (
	"context"
	"fmt"
	"log"
	"net/http"

	"github.com/bluayer/GoingGong/ent"
	"github.com/bluayer/GoingGong/ent/user"
	"github.com/bluayer/GoingGong/model"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func Create(client *ent.Client) (e *echo.Echo) {
	ctx := context.Background()
	app := echo.New()
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())

	app.GET("/", func(c echo.Context) error {
		users, err := client.User.Query().All(ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Server ERROR")
		}

		return c.JSON(http.StatusOK, echo.Map{
			"status": true,
			"users":  users,
		})
	})

	app.GET("/user/:name", func(c echo.Context) error {
		name := c.Param("name")
		foundUser, err := client.User.Query().Where(user.NameEQ(name)).First(ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "There's error at find user")
		}
		updatedUser, err := client.User.UpdateOneID(foundUser.ID).SetPingCnt(foundUser.PingCnt + 1).Save(ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "There's error at update user")
		}
		return c.JSON(http.StatusOK, echo.Map{
			"status": true,
			"user":   updatedUser,
		})
	})

	app.POST("/user", func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Bad Request")
		}
		fmt.Print(u.Name)
		_, err := client.User.Create().SetName(u.Name).SetPingCnt(0).Save(ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Server can't make user")
		}
		return c.String(http.StatusOK, "User Created!")
	})
	return app
}

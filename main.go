package main

import (
	"context"
	"goingong/ent"

	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	"github.com/bluayer/GoingGong/model"
)

func main() {
	client, err := ent.Open("sqlite3", "file:ent?mode=memory&cache=shared&_fk=1")
	ctx := context.Background()
	if err != nil {
		log.Fatalf("failed opening connection to sqlite: %v", err)
	}
	defer client.Close()
	// Run the auto migration tool.
	if err := client.Schema.Create(context.Background()); err != nil {
		log.Fatalf("failed creating schema resources: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", func(c echo.Context) error {
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

	e.POST("/user", func(c echo.Context) error {
		u := new(model.User)
		if err := c.Bind(u); err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Server ERROR")
		}
		_, err := client.User.Create().SetName(u.Name).Save(ctx)
		if err != nil {
			log.Println(err)
			return c.String(http.StatusBadRequest, "Server ERROR")
		}
		return c.String(http.StatusOK, "User Created!")
	})

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}

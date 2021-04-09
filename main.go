package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/mazharku/go-sqlc-echo-crud/schema"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	toml "github.com/pelletier/go-toml"
)

const SetMaxOpenConns = 20
const Offset = 0
const LIMIT = 25

var config *toml.Tree
var db *sql.DB
var database *schema.Queries

func main() {
	//Load config file
	var err error
	config, err = toml.LoadFile("config.ini")
	if err != nil {
		log.Fatalf("Toml Error: %v", err)
	}

	//DB connection
	dbconfig := config.Get("database").(*toml.Tree)
	db = dbconn(dbconfig)
	defer db.Close()
	database = schema.New(db)

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Routes
	e.GET("/", status)
	e.GET("/items", GetAll)
	e.POST("/items", CreateItem)
	// Start server
	address := fmt.Sprintf("%s:%s", config.Get("host").(string), config.Get("port").(string))
	e.Logger.Fatal(e.Start(address))

}

func status(c echo.Context) error {
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "api works!s"})
}

func GetAll(c echo.Context) error {
	context := context.Background()

	data, err := database.FindAll(context)
	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{"status": false, "message": err.Error()})
	}

	return c.JSON(http.StatusOK, data)
}

func CreateItem(c echo.Context) error {
	context := context.Background()
	data := &schema.CreateItemParams{}
	err := c.Bind(data)

	if err != nil {
		return c.JSON(http.StatusOK, echo.Map{"status": false, "message": err.Error()})
	}

	response := database.CreateItem(context, *data)
	if response != nil {
		return c.JSON(http.StatusOK, echo.Map{"status": false, "message": err.Error()})
	}
	return c.JSON(http.StatusOK, echo.Map{"status": true, "message": "Resource created"})
}

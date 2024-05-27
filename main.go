package main

import (
	"fmt"
	"log"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html/v2"
)

func main() {
	engine := html.New("./templates", ".html")
	app := fiber.New(fiber.Config{
		Views: engine,
	})
	app.Static("/css", "./static/css")
	app.Static("/js", "./static/js")
	app.Static("/images", "./static/images")
	app.Use("/favicon.ico", func(c *fiber.Ctx) error {
		return c.SendFile("./static/images/favicon.ico")
	})

	app.Get("/", Index)
	app.Post("/GetForm", GetForm)
	log.Fatal(app.Listen(":3000"))
}

func Index(c *fiber.Ctx) error {
	return c.Render("index", fiber.Map{})
}

func GetForm(c *fiber.Ctx) error {
	name := c.FormValue("name")
	// string name upper case
	nameUpper := strings.ToUpper(name)
	telefone := c.FormValue("telefone")
	age := c.FormValue("age")
	fmt.Println(nameUpper, telefone, age)
	data := map[string]interface{}{
		"Name": nameUpper,
		"Tel":  telefone,
		"Age":  age,
	}
	return c.Render("result", data)
}

package models

import "github.com/gofiber/fiber/v2"

type MethodEnum string

const (
	GET    MethodEnum = "GET"
	POST   MethodEnum = "POST"
	PUT    MethodEnum = "PUT"
	PATCH  MethodEnum = "PATCH"
	DELETE MethodEnum = "DELETE"
)

type Router struct {
	Method      MethodEnum
	Path        string
	RelatedFunc func(c *fiber.Ctx) error
}

type ResponseStructure struct {
	Data    interface{} `json:"data" xml:"data" form:"data"`
	Message string      `json:"message" xml:"message" form:"message"`
	Success bool        `json:"success" xml:"success" form:"success"`
}

type Person struct {
	Name string `json:"name" xml:"name" form:"name"`
	Pass string `json:"pass" xml:"pass" form:"pass"`
}

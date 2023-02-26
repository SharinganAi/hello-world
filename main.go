package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/:name", IndexHandler)
	router.GET("/", IndexHandlerXml)
	router.Run()
}

func IndexHandler(c *gin.Context) {
	name := c.Params.ByName("name")
	c.JSON(200, gin.H{
		"message": "hello " + name,
	})
}

func IndexHandlerXml(c *gin.Context) {
	c.XML(200, Person{
		FirstName: "Dilip",
		LastName:  "Singh",
	})
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstName,attr"`
	LastName  string   `xml:"lastName,attr"`
}

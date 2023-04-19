package main

import (
	"encoding/xml"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", IndexHandler)
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World"})
	// })
	router.Run(":8080")
}

type Person struct {
	XMLName   xml.Name `xml:"person"`
	FirstName string   `xml:"firstname,attr"`
	LastName  string   `xml:"lastname,attr"`
}

func IndexHandler(c *gin.Context) {

	c.XML(200, Person{FirstName: "John", LastName: "Appleseed"})
}

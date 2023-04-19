package main

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/xid"
)

type Recipe struct {
	ID           string    `json:"id"`
	Name         string    `json:"name"`
	Tags         []string  `json:"tags"`
	Ingredients  []string  `json:"ingredients"`
	Instructions []string  `json:"instructions"`
	PublisedAt   time.Time `json:"publishedAt"`
}

var recipes []Recipe

func init() {
	recipes = make([]Recipe, 0)
}

func main() {
	router := gin.Default()
	router.POST("/recipes", NewRecipeHandler)
	router.Run()
	// router.GET("/", IndexHandler)
	// router.GET("/", func(c *gin.Context) {
	// 	c.JSON(200, gin.H{
	// 		"message": "Hello World"})
	// })

}

func NewRecipeHandler(c *gin.Context) {
	var recipe Recipe
	if err := c.ShouldBindJSON(&recipe); err != nil {
		c.JSON(400, gin.H{"error": err.Error()})
		return
	}
	recipe.ID = xid.New().String()
	recipe.PublisedAt = time.Now()
	recipes = append(recipes, recipe)
	c.JSON(200, recipe)
}

// type Person struct {
// 	XMLName   xml.Name `xml:"person"`
// 	FirstName string   `xml:"firstname,attr"`
// 	LastName  string   `xml:"lastname,attr"`
// }

// func IndexHandler(c *gin.Context) {

// 	c.XML(200, Person{FirstName: "John", LastName: "Appleseed"})
// }

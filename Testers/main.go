package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type articleClothing struct {
	ID    string  `json:"id"`
	Tag   string  `json:"tag"`
	Brand string  `json:"brand"`
	Price float64 `json:"price"`
	Wears int32   `json:"wears"`
	CPW   float64 `json:"CPW"`
}

var clothes []articleClothing

func main() {

	clothes = append(clothes, articleClothing{ID: "001", Tag: "Type: T-Shirt Size: M", Price: 19.99, Wears: 1})
	clothes = append(clothes, articleClothing{ID: "002", Tag: "Type: Jeans Size: L", Price: 49.99, Wears: 5})
	clothes = append(clothes, articleClothing{ID: "003", Tag: "Type: Sneakers Size: 10", Price: 89.99, Wears: 25})

	for i := range clothes {
		clothes[i].CPW = clothes[i].Price / float64(clothes[i].Wears)
	}

	router := gin.Default()
	router.GET("/clothes", getClothes)
	router.GET("/clothes/:id", getClothesByID)
	router.POST("/clothes", postClothes)

	router.Run("localhost:8089")
}

func getClothes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, clothes)
}

func postClothes(c *gin.Context) {
	var newArticle articleClothing

	if err := c.BindJSON(&newArticle); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	clothes = append(clothes, newArticle)

	c.IndentedJSON(http.StatusCreated, newArticle)
}

func getClothesByID(c *gin.Context) {
	id := c.Param("id")

	for _, a := range clothes {
		if a.ID == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "article not found"})
}

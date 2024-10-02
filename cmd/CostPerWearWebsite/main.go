package main

import (
	// "github.com/Henry-Bodwell/CostPerWear/internal/app"

	// "net/http"
	"github.com/gin-gonic/gin"
	// "errors"
)

func getArticles(C *gin.Context) {
	// Go to database get articles of clothes

	// C.IndentedJSON(http.StatusOK, clothes)
}

func main() {
	router := gin.Default()

	router.Run("localhost:9090")
}

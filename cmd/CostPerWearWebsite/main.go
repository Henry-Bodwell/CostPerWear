package main

import (
	"fmt"

	"github.com/Henry-Bodwell/CostPerWear/internal/app"
)

func main() {
	myCloset := app.NewCloset("MainCloset")

	myTags := app.NewSet[string]()
	myTags.AddSlice([]string{"Grapic", "Short Sleeve", "Grey"})

	newShirt := app.NewClothes("Riki Shirt", "path/to/image", 24.99, "Cotton", "Comfort Colors", "Year Round", *myTags)

	myCloset.AddClothes(newShirt)

	fmt.Println(myCloset.GetTotalWears())
	myCloset.WearArticle(newShirt)
	fmt.Println(myCloset.GetTotalWears())
}

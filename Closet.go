package main

import (
	"strings"
)

type Closet struct {
	ID                  int          `json:"id"`
	Name                string       `json:"name"`
	AllClothes          []*Clothing  `json:"allClothes"`
	AllOutfits          []*Outfit    `json:"allOutfits"`
	UniqueTags          *Set[string] `json:"uniqueTags"`
	UniqueBrands        *Set[string] `json:"uniqueBrands"`
	UniqueMaterial      *Set[string] `json:"uniqueMaterial"`
	UniqueVibes         *Set[string] `json:"uniqueVibes"`
	UniqueClothingNames *Set[string] `json:"uniqueClothingNames"`

	OutfitLookup map[*Clothing][]*Outfit

	TotalWears    uint
	TotalArticles uint
	AvgWears      float32
	AvgCPW        float32
	TotalPrice    float32
}

// oldClosetImport: Imports existing clothes and outfits to new closet
func OldClosetImport(Name string, AllClothes []*Clothing, AllOutfits []*Outfit) *Closet {
	var myCloset = &Closet{
		AllClothes: AllClothes,
		AllOutfits: AllOutfits,
	}
	myCloset.Name = Name
	myCloset.OutfitLookup = make(map[*Clothing][]*Outfit)

	myCloset.UniqueBrands = NewSet[string]()
	myCloset.UniqueTags = NewSet[string]()
	myCloset.UniqueMaterial = NewSet[string]()
	for _, article := range AllClothes {
		myCloset.UniqueBrands.Add(article.Brand)
		myCloset.UniqueMaterial.Add(article.Material)
		myCloset.UniqueTags.AddAll(article.Tags)
		myCloset.UniqueClothingNames.Add(article.Name)
		myCloset.TotalWears += article.Wears
		myCloset.TotalPrice += article.Price
	}
	for _, fit := range AllOutfits {
		myCloset.updateOutfitMetrics(fit)
	}

	myCloset.updateTotalArticles()
	myCloset.updateAvgCPW()
	myCloset.updateAvgWears()

	return myCloset
}

// Default Constructor
func NewCloset(Name string) *Closet {
	var c = &Closet{}
	c.Name = Name
	c.OutfitLookup = make(map[*Clothing][]*Outfit)
	c.UniqueTags = NewSet[string]()
	c.UniqueBrands = NewSet[string]()
	c.UniqueMaterial = NewSet[string]()
	c.UniqueVibes = NewSet[string]()
	return c
}

// updateAvgCPW: total Price by total Wears
func (c *Closet) updateAvgCPW() {
	c.AvgCPW = c.TotalPrice / float32(c.TotalWears)
}

// updateTotalItems: updates total number of articles of clothes
func (c *Closet) updateTotalArticles() {
	c.TotalArticles = uint(len(c.AllClothes))
}

// updateAvgWears: total Wears by total Items
func (c *Closet) updateAvgWears() {
	c.AvgWears = float32(c.TotalWears) / float32(c.TotalArticles)
}

// addsClothes: Adds a new clothes Item to closet
func (c *Closet) AddClothes(article *Clothing) {
	if !c.UniqueClothingNames.Contains(article.Name) {
		c.AllClothes = append(c.AllClothes, article)
		c.UniqueTags.AddAll(article.Tags)
		c.UniqueBrands.Add(article.Brand)
		c.UniqueMaterial.Add(article.Material)

		c.TotalPrice += article.Price
		c.TotalWears += article.Wears
		c.TotalArticles++

		c.updateAvgCPW()
		c.updateAvgWears()
	} else {
		// Not sure how but maybe throw an error or something
	}
}

// addOutfit: add new outfit to closet
func (c *Closet) AddOutfit(fit *Outfit) {
	c.AllOutfits = append(c.AllOutfits, fit)
	c.updateOutfitMetrics(fit)

	c.TotalArticles += fit.numItems
	c.TotalPrice += fit.OutfitPrice
	c.TotalWears += fit.TotalWears

	c.updateAvgCPW()
	c.updateAvgWears()

}

// searchClothes: TODO search
func (c *Closet) SearchClothes(key string, Brand string, Material string, Tags []string) []Clothing {
	var result []Clothing

	// I think i want this to look maybe also return a slice of outfits the item is in??? would this be too slow?

	for _, clothing := range c.AllClothes {
		if articleMatches(clothing, key, Brand, Material, Tags) {
			result = append(result, *clothing)
		}
	}

	return result
}

// searchOutfits: TODO Search
func (c *Closet) SearchOutfits(key string, Vibe string, season string, Tags []string) []Outfit {
	var result []Outfit

	// I think i want this to look maybe also return a slice of outfits the item is in??? would this be too slow?

	for _, fit := range c.AllOutfits {
		if outfitMatches(fit, key, Vibe, season, Tags) {
			result = append(result, *fit)
		}
	}

	return result
}

// Maybe I need a map that has clothing article as a key and the set of outfits it is a part of as the values

// outfitMatches: Checks outfit against filters and returns true if it matches all search criteria
func outfitMatches(outfit *Outfit, key string, Vibe string, season string, Tags []string) bool {
	if key != "" && !containsIgnoreCase(outfit.Name, key) {
		return false
	}

	if Vibe != "" && !containsIgnoreCase(outfit.Vibe, Vibe) {
		return false
	}

	if season != "" && !containsIgnoreCase(outfit.Season, season) {
		return false
	}

	if len(Tags) > 0 && !hasAllTags(outfit.Tags, Tags) {
		return false
	}

	return true
}

// articleMatches: Checks clothing article against filters and returns true if it matches all search criteria
func articleMatches(article *Clothing, key string, Brand string, Material string, Tags []string) bool {
	if key != "" && !containsIgnoreCase(article.Name, key) {
		return false
	}

	if Brand != "" && !containsIgnoreCase(article.Brand, Brand) {
		return false
	}

	if Material != "" && !containsIgnoreCase(article.Material, Material) {
		return false
	}

	if len(Tags) > 0 && !hasAllTags(article.Tags, Tags) {
		return false
	}

	return true
}

// containsIgnoreCase: Modified version of strings.contains that compares the lowercase version of string and substring
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// hasAllTags: iterate through search Tags return false if one tag is not found in itemTags set
func hasAllTags(itemTags Set[string], searchTags []string) bool {
	for _, searchTag := range searchTags {
		if found := itemTags.Contains(searchTag); !found {
			return false
		}
	}

	return true
}

// updateOutfitMetrics: takes the outfit and updates the closet unique Vibes and Tags with outfits data.
// Also adds all clothuing articles to the outfit loopup map
func (c *Closet) updateOutfitMetrics(newFit *Outfit) {
	c.UniqueTags.AddAll(newFit.Tags)
	c.UniqueVibes.Add(newFit.Vibe)
	helper := []*Clothing{newFit.Top, newFit.Bottom, newFit.Shoes}
	for _, article := range helper {
		if article != nil {
			c.OutfitLookup[article] = append(c.OutfitLookup[article], newFit)
		}
	}
	if len(newFit.Accessories) > 0 {
		for _, accesory := range newFit.Accessories {
			c.OutfitLookup[accesory] = append(c.OutfitLookup[accesory], newFit)
		}
	}
}

// func (c *Closet) removeArticle(articleToRemove string) string {
// 	c.AllClothes = rem
// 	return ""
// }

// func (c *Closet) removeOutfit(fitToRemove string) string {

// 	return ""
// }

func (c *Closet) GetTotalWears() uint {
	return c.TotalWears
}

func (c *Closet) WearArticle(article *Clothing) {
	article.IncrementWears()
	c.TotalWears++
	c.updateAvgCPW()
	c.updateAvgWears()

}

func (c *Closet) WearOutfit(fit *Outfit) {
	fit.incrementWears()
	c.TotalWears += fit.numItems
}

func (c *Closet) GetUniqueTags() *Set[string] {
	return c.UniqueTags
}

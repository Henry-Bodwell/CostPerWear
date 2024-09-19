package app

import "strings"

type Closet struct {
	name           string
	allClothes     []*Clothing
	allOutfits     []*Outfit
	uniqueTags     *Set[string]
	uniqueBrands   *Set[string]
	uniqueMaterial *Set[string]
	uniqueVibes    *Set[string]

	outfitLookup map[*Clothing][]*Outfit

	totalWears uint
	totalItems uint
	avgWears   float32
	avgCPW     float32
	totalPrice float32
}

// oldClosetImport: Imports existing clothes and outfits to new closet
func oldClosetImport(name string, allClothes []*Clothing, allOutfits []*Outfit) *Closet {
	var myCloset = &Closet{
		allClothes: allClothes,
		allOutfits: allOutfits,
	}
	myCloset.name = name
	myCloset.outfitLookup = make(map[*Clothing][]*Outfit)

	myCloset.uniqueBrands = NewSet[string]()
	myCloset.uniqueTags = NewSet[string]()
	myCloset.uniqueMaterial = NewSet[string]()
	for _, article := range allClothes {
		myCloset.uniqueBrands.Add(article.brand)
		myCloset.uniqueMaterial.Add(article.material)
		myCloset.uniqueTags.AddAll(article.tags)
	}

	for _, fit := range allOutfits {
		myCloset.updateOutfitMetrics(fit)
	}

	return myCloset
}

// Default Constructor
func newCloset(name string) *Closet {
	var c = &Closet{}
	c.name = name
	c.outfitLookup = make(map[*Clothing][]*Outfit)
	c.uniqueTags = NewSet[string]()
	c.uniqueBrands = NewSet[string]()
	c.uniqueMaterial = NewSet[string]()
	c.uniqueVibes = NewSet[string]()
	return c
}

// UpdateTotalPrice: loops through all articles of clothing and sum price
func (c *Closet) updateTotalPrice() {
	var sum float32
	for _, article := range c.allClothes {
		sum += article.price
	}
	c.totalPrice = sum
}

// updateAvgCPW: total price by total wears
func (c *Closet) updateAvgCPW() {
	c.avgCPW = c.totalPrice / float32(c.totalWears)
}

// updateTotalItems: updates total number of items
func (c *Closet) updateTotalItems() {
	var sum uint
	for range c.allClothes {
		sum += 1
	}
	c.totalItems = sum
}

// updateTotalWears: Loop through articles and sum the total wears -- ineffiecient should probably add an increment de-incrment
func (c *Closet) updateTotalWears() {
	var sum uint
	for _, article := range c.allClothes {
		sum += article.wears
	}
	c.totalWears = sum
}

// updateAvgWears: total wears by total Items
func (c *Closet) updateAvgWears() {
	c.avgWears = float32(c.totalWears) / float32(c.totalItems)
}

// addsClothes: Adds a new clothes Item to closet
func (c *Closet) addClothes(article *Clothing) {
	// TODO: Add Clothes
	c.allClothes = append(c.allClothes, article)
	c.uniqueTags.AddAll(article.tags)
	c.uniqueBrands.Add(article.brand)
	c.uniqueMaterial.Add(article.material)
}

// addOutfit: add new outfit to closet
func (c *Closet) addOutfit(fit *Outfit) {
	c.allOutfits = append(c.allOutfits, fit)
	c.updateOutfitMetrics(fit)
}

// searchClothes: TODO search
func (c *Closet) searchClothes(key string, brand string, material string, tags []string) []Clothing {
	var result []Clothing

	// I think i want this to look maybe also return a slice of outfits the item is in??? would this be too slow?

	for _, clothing := range c.allClothes {
		if articleMatches(clothing, key, brand, material, tags) {
			result = append(result, *clothing)
		}
	}

	return result
}

// searchOutfits: TODO Search
func (c *Closet) searchOutfits(key string, vibe string, season string, tags []string) []Outfit {
	var result []Outfit

	// I think i want this to look maybe also return a slice of outfits the item is in??? would this be too slow?

	for _, fit := range c.allOutfits {
		if outfitMatches(fit, key, vibe, season, tags) {
			result = append(result, *fit)
		}
	}

	return result
}

// Maybe I need a map that has clothing article as a key and the set of outfits it is a part of as the values

// outfitMatches: Checks outfit against filters and returns true if it matches all search criteria
func outfitMatches(outfit *Outfit, key string, vibe string, season string, tags []string) bool {
	if key != "" && !containsIgnoreCase(outfit.name, key) {
		return false
	}

	if vibe != "" && !containsIgnoreCase(outfit.vibe, vibe) {
		return false
	}

	if season != "" && !containsIgnoreCase(outfit.season, season) {
		return false
	}

	if len(tags) > 0 && !hasAllTags(outfit.tags, tags) {
		return false
	}

	return true
}

// articleMatches: Checks clothing article against filters and returns true if it matches all search criteria
func articleMatches(article *Clothing, key string, brand string, material string, tags []string) bool {
	if key != "" && !containsIgnoreCase(article.name, key) {
		return false
	}

	if brand != "" && !containsIgnoreCase(article.brand, brand) {
		return false
	}

	if material != "" && !containsIgnoreCase(article.material, material) {
		return false
	}

	if len(tags) > 0 && !hasAllTags(article.tags, tags) {
		return false
	}

	return true
}

// containsIgnoreCase: Modified version of strings.contains that compares the lowercase version of string and substring
func containsIgnoreCase(str, substr string) bool {
	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}

// hasAllTags: iterate through search tags return false if one tag is not found in itemTags set
func hasAllTags(itemTags Set[string], searchTags []string) bool {
	for _, searchTag := range searchTags {
		if found := itemTags.Contains(searchTag); !found {
			return false
		}
	}

	return true
}

// updateOutfitMetrics: takes the outfit and updates the closet unique vibes and tags with outfits data.
// Also adds all clothuing articles to the outfit loopup map
func (c *Closet) updateOutfitMetrics(newFit *Outfit) {
	c.uniqueTags.AddAll(newFit.tags)
	c.uniqueVibes.Add(newFit.vibe)
	helper := []*Clothing{newFit.top, newFit.bottom, newFit.shoes}
	for _, article := range helper {
		if article != nil {
			c.outfitLookup[article] = append(c.outfitLookup[article], newFit)
		}
	}
	if len(newFit.accessories) > 0 {
		for _, accesory := range newFit.accessories {
			c.outfitLookup[accesory] = append(c.outfitLookup[accesory], newFit)
		}
	}
}

func (c *Closet) removeArticle(articleToRemove string) string {

	return ""
}

func (c *Closet) removeOutfit(fitToRemove string) string {

	return ""
}

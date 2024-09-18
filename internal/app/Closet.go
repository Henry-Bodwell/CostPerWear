package app

type Closet struct {
	name           string
	allClothes     []*Clothing
	allOutfits     []*Outfit
	uniqueTags     Set[string]
	uniqueBrands   Set[string]
	uniqueMaterial Set[string]
	uniqueVibes    Set[string]

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

	myCloset.uniqueBrands = *NewSet[string]()
	myCloset.uniqueTags = *NewSet[string]()
	myCloset.uniqueMaterial = *NewSet[string]()
	for _, article := range allClothes {
		myCloset.uniqueBrands.Add(article.brand)
		myCloset.uniqueMaterial.Add(article.material)
		myCloset.uniqueTags.AddAll(article.tags)
	}

	for _, fit := range allOutfits {
		myCloset.uniqueVibes.Add(fit.vibe)
	}

	return myCloset
}

// Default Constructor
func newCloset(name string) *Closet {
	var c = &Closet{}
	c.name = name

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
	c.uniqueTags.AddAll(fit.tags)
	c.uniqueVibes.Add(fit.vibe)
}

// searchClothes: TODO search
func (c *Closet) searchClothes(key string) {
	// I think i want this to look maybe also return a slice of outfits the item is in??? would this be too slow?
}

// searchOutfits: TODO Search
func (c *Closet) searchOutfits(key string) {

}

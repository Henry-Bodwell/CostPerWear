package app

type Closet struct {
	allClothes     []*Clothing
	allOutfits     []*Outfit
	uniqueTags     Set[string]
	uniqueBrands   Set[string]
	uniqueMaterial Set[string]

	totalWears uint
	totalItems uint
	avgWears   float32
	avgCPW     float32
	totalPrice float32
}

// TODO: Write Constructor
func (c *Closet) newCloset(allClothes []*Clothing, allOutfits []*Outfit) *Closet {
	var myCloset = &Closet{
		allClothes: allClothes,
		allOutfits: allOutfits,
	}
	c.uniqueBrands = *NewSet[string]()
	c.uniqueTags = *NewSet[string]()
	c.uniqueMaterial = *NewSet[string]()
	for _, article := range allClothes {
		c.uniqueBrands.Add(article.brand)
		c.uniqueMaterial.Add(article.material)
		c.uniqueTags.AddAll(article.tags)
	}

	return myCloset
}

// UpdateTotalPrice: loops through all articles of clothing and sum price
func (c *Closet) updateTotalPrice() {
	var sum float32
	for _, article := range c.allClothes {
		sum += article.price
	}
	c.totalPrice = sum
}

// updateAvgCPW, total price by total wears
func (c *Closet) updateAvgCPW() {
	c.avgCPW = c.totalPrice / float32(c.totalWears)
}

// get total number of items
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

// Creates a set
func (c *Closet) updateUniqueTags() {
	//TODO: make sets
}

func (c *Closet) updateUniqueBrands() {
	//TODO: make sets
}

func (c *Closet) updateUniqueMaterials() {
	//TODO: make sets
}

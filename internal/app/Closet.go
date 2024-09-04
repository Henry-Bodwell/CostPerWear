package app

type Closet struct {
	allClothes     []Clothing
	allOutfits     []Outfit
	uniqueTags     []string
	uniqueBrands   []string
	uniqueMaterial []string
	totalWears     uint
	totalItems     uint
	avgWears       float32
	avgCPW         float32
	totalPrice     float32
}

func (c *Closet) updateTotalPrice() {
	var sum float32
	for _, article := range c.allClothes {
		sum += article.price
	}
	c.totalPrice = sum
}

func (c *Closet) updateAvgCPW() {
	c.avgCPW = c.totalPrice / float32(c.totalWears)
}

func (c *Closet) updateTotalItems() {
	var sum uint
	for range c.allClothes {
		sum += 1
	}
	c.totalItems = sum
}

func (c *Closet) updateTotalWears() {
	var sum uint
	for _, article := range c.allClothes {
		sum += article.wears
	}
	c.totalWears = sum
}

func (c *Closet) updateAvgWears() {
	c.avgWears = float32(c.totalWears) / float32(c.totalItems)
}

func (c *Closet) updateUniqueTags() {
	//TODO: make sets
}

func (c *Closet) updateUniqueBrands() {
	//TODO: make sets
}

func (c *Closet) updateUniqueMaterials() {
	//TODO: make sets
}

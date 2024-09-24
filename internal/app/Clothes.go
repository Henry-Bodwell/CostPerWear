package app

type Clothing struct {
	name        string
	image       string
	price       float32
	wears       uint
	material    string
	brand       string
	season      string
	costPerWear float32
	tags        Set[string]
}

// Constructor
func NewClothes(name string, image string, price float32, material string, brand string, season string, tags Set[string]) *Clothing {
	return &Clothing{
		name:        name,
		image:       image,
		price:       price,
		material:    material,
		brand:       brand,
		season:      season,
		wears:       0,
		costPerWear: price,
		tags:        tags,
	}

}

// UpdateCPW: calculates the value of the cost per wear variable
// CPW = Price / Wears
func (c *Clothing) updateCPW() {
	c.costPerWear = c.price / float32(c.wears)
}

// incrementWears: does just that
func (c *Clothing) incrementWears() {
	c.wears++
	c.updateCPW()
}

// updateImage: Update the image path to string
func (c *Clothing) UpdateImage(path string) {
	c.image = path
}

// addTag: add a tag to the list of tags
func (c *Clothing) AddTag(tag string) {
	c.tags.Add(tag)
}

// removeTag: Takes string arg and if it exists in tags remove it return removeTag, if it is not in tags return ""
func (c *Clothing) RemoveTag(tagToRemove string) string {
	if c.tags.Contains(tagToRemove) {
		c.tags.Remove(tagToRemove)
		return tagToRemove
	} else {
		return ""
	}
}

package app

type Clothing struct {
	id          string
	image       string
	price       float32
	wears       uint
	material    string
	brand       string
	season      string
	costPerWear float32
	tags        []string
}
// UpdateCPW calculates the value of the cost per wear variable
// CPW = Price / Wears
func (c *Clothing) updateCPW() {
	c.costPerWear = c.price / float32(c.wears)
}

// incrementWears does just that
func (c *Clothing) incrementWears() {
	c.wears++
}

func (c *Clothing) updateImage(c.path String) {
	c.image = path
}

func (c *Clothing) addTag(tag String) {
	c.tag = append(c.tag, tag)
}
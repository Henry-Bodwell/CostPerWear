package app

type Clothing struct {
	ID           int         `json:"id"`
	Name         string      `json:"name"`
	Image        string      `json:"image"`
	Price        float32     `json:"price"`
	Wears        uint        `json:"wears"`
	Material     string      `json:"material"`
	Brand        string      `json:"brand"`
	Season       string      `json:"season"`
	CostPerWear  float32     `json:"costPerWear"`
	Tags         Set[string] `json:"tags"`
	ClothingType string      `json:"clothingType"`
}

// Constructor
func NewClothes(Name string, Image string, Price float32, Material string, Brand string, Season string, Tags Set[string], clothingType string) *Clothing {
	return &Clothing{
		Name:         Name,
		Image:        Image,
		Price:        Price,
		Material:     Material,
		Brand:        Brand,
		Season:       Season,
		Wears:        0,
		CostPerWear:  Price,
		Tags:         Tags,
		ClothingType: clothingType,
	}

}

// UpdateCPW: calculates the value of the cost per wear variable
// CPW = Price / Wears
func (c *Clothing) UpdateCPW() {
	if c.Wears != 0 {
		c.CostPerWear = c.Price / float32(c.Wears)
	}
}

// incrementWears: does just that
func (c *Clothing) IncrementWears() {
	c.Wears++
	c.UpdateCPW()
}

// updateImage: Update the image path to string
func (c *Clothing) UpdateImage(path string) {
	c.Image = path
}

// addTag: add a tag to the list of tags
func (c *Clothing) AddTag(tag string) {
	c.Tags.Add(tag)
}

// removeTag: Takes string arg and if it exists in tags remove it return removeTag, if it is not in tags return ""
func (c *Clothing) RemoveTag(tagToRemove string) string {
	if c.Tags.Contains(tagToRemove) {
		c.Tags.Remove(tagToRemove)
		return tagToRemove
	} else {
		return ""
	}
}

// GetName: returns string
func (c *Clothing) GetName() string {
	return c.Name
}

// GetWears: returns uint
func (c *Clothing) GetWears() uint {
	return c.Wears
}

// GetPrice: returns float32
func (c *Clothing) GetPrice() float32 {
	return c.Price
}

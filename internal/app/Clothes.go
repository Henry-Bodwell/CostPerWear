package app

import (
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"math/rand"
	"time"
)

type Clothing struct {
	id          string
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

func newClothes(name string, image string, price float32, material string, brand string, season string, tags Set[string]) *Clothing {
	id := generateID(name)
	return &Clothing{
		id:          id,
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

// Generates reasonably unqiue ID code
func generateID(name string) string {
	// Combine current time and a random number to generate a seed
	seed := fmt.Sprintf("%d%s%d", time.Now().UnixNano(), name, rand.Int())

	// Hash the seed using SHA-1
	hash := sha1.New()
	hash.Write([]byte(seed))

	// Convert the hash to a hex string
	return hex.EncodeToString(hash.Sum(nil))[:8] // Take the first 8 characters for a shorter ID
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

// Update the image path to string
func (c *Clothing) updateImage(path string) {

	c.image = path
}

// add a tag to the list of tags
func (c *Clothing) addTag(tag string) {
	c.tags.Add(tag)
}

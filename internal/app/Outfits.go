package app

type Outfit struct {
	top         Clothing
	bottom      Clothing
	shoes       Clothing
	accessories Clothing

	id         string
	formatilty string
	season     string
	tags       []string

	outfitPrice float32
	outfitWears uint
	totalWears  uint
	outfitCPW   float32
	avgCPW      float32
}

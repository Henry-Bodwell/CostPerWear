package app

type Outfit struct {
	top         *Clothing
	bottom      *Clothing
	shoes       *Clothing
	accessories []*Clothing

	id     string
	name   string
	vibe   string
	season string
	tags   Set[string]

	numItems    uint
	outfitPrice float32
	outfitWears uint
	totalWears  uint
	outfitCPW   float32
	avgCPW      float32
}

// newOutfit: Constructor
func newOutfit(top *Clothing, bottom *Clothing, shoes *Clothing, accessories []*Clothing, name string, vibe string, season string) *Outfit {
	var newFit = &Outfit{
		id:          generateID(name),
		top:         top,
		bottom:      bottom,
		shoes:       shoes,
		accessories: accessories,
		name:        name,
		vibe:        vibe,
		season:      season,
		tags:        *NewSet[string](),
	}

	newFit.tags.AddAll(top.tags)

	newFit.updateNumItems()
	newFit.calcPrice()
	newFit.calcTotalWears()
	newFit.calcAvgCPW()
	return newFit
}

// incrementWears: Increments the wears of outfit and all contained clothing
func (o *Outfit) incrementWears() {
	o.top.incrementWears()
	o.bottom.incrementWears()
	o.shoes.incrementWears()

	for _, item := range o.accessories {
		item.incrementWears()
	}

	o.outfitWears++
}

// calcPrice: saves the price of the outfit components
func (o *Outfit) calcPrice() {
	var accessoryPrice float32
	for _, article := range o.accessories {
		accessoryPrice += article.price
	}
	o.outfitPrice = o.top.price + o.bottom.price + o.shoes.price + accessoryPrice
}

// calcTotalWears: calc total wears of outfit components
func (o *Outfit) calcTotalWears() {
	var accessoryWears uint
	for _, article := range o.accessories {
		accessoryWears += article.wears
	}
	o.totalWears = o.top.wears + o.bottom.wears + o.shoes.wears + accessoryWears
}

// calcOutfit CPW:
// Update the cost per wear for the whole outfit
// outfitCPW = outfit price / outfit wears
func (o *Outfit) calcOutfitCPW() {
	o.outfitCPW = o.outfitPrice / float32(o.outfitWears)
}

// UpdateNumItems: designed to run at the creation or editing of an outfit
// updates variable numItems with count of all non-nil clothing pointers
func (o *Outfit) updateNumItems() {
	var articleCount uint
	list := []*Clothing{o.top, o.bottom, o.shoes}
	articleCount = uint(len(o.accessories))

	for _, article := range list {
		if article != nil {
			articleCount++
		}
	}
	o.numItems = articleCount
}

// calcAvgCPW: Updates avgCPW variable with the average cost per wear of the indiviual items
// Sum(Each items cpw) / num items
func (o *Outfit) calcAvgCPW() {
	var accessoryCPW float32
	for _, article := range o.accessories {
		accessoryCPW += article.costPerWear
	}
	o.avgCPW = (accessoryCPW + o.top.costPerWear + o.bottom.costPerWear + o.shoes.costPerWear) / float32(o.numItems)
}

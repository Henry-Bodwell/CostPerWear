package app

type Outfit struct {
	ID          int         `json:"id"`
	Top         *Clothing   `json:"top"`
	Bottom      *Clothing   `json:"bottom"`
	Shoes       *Clothing   `json:"shoes"`
	Accessories []*Clothing `json:"accessories"`

	Name   string      `json:"name"`
	Vibe   string      `json:"vibe"`
	Season string      `json:"season"`
	Tags   Set[string] `json:"tags"`

	numItems    uint
	OutfitPrice float32 `json:"outfitPrice"`
	OutfitWears uint    `json:"outfitWears"`
	TotalWears  uint    `json:"totalWears"`
	OutfitCPW   float32 `json:"outfitCPW"`
	AvgCPW      float32 `json:"avgCPW"`
}

// newOutfit: Constructor
func NewOutfit(Top *Clothing, Bottom *Clothing, Shoes *Clothing, Accessories []*Clothing, Name string, Vibe string, Season string) *Outfit {
	var newFit = &Outfit{
		Top:         Top,
		Bottom:      Bottom,
		Shoes:       Shoes,
		Accessories: Accessories,
		Name:        Name,
		Vibe:        Vibe,
		Season:      Season,
		Tags:        *NewSet[string](),
	}

	newFit.Tags.AddAll(Top.Tags)

	newFit.updateNumItems()
	newFit.calcPrice()
	newFit.calcTotalWears()
	newFit.calcAvgCPW()
	return newFit
}

// incrementWears: Increments the Wears of outfit and all contained clothing
func (o *Outfit) incrementWears() {
	o.Top.IncrementWears()
	o.Bottom.IncrementWears()
	o.Shoes.IncrementWears()

	for _, item := range o.Accessories {
		item.IncrementWears()
	}

	o.OutfitWears++
	o.calcAvgCPW()
	o.calcOutfitCPW()
}

// calcPrice: saves the Price of the outfit components
func (o *Outfit) calcPrice() {
	var accessoryPrice float32
	for _, article := range o.Accessories {
		accessoryPrice += article.Price
	}
	o.OutfitPrice = o.Top.Price + o.Bottom.Price + o.Shoes.Price + accessoryPrice
}

// calcTotalWears: calc total Wears of outfit components
func (o *Outfit) calcTotalWears() {
	var accessoryWears uint
	for _, article := range o.Accessories {
		accessoryWears += article.Wears
	}
	o.TotalWears = o.Top.Wears + o.Bottom.Wears + o.Shoes.Wears + accessoryWears
}

// calcOutfit CPW:
// Update the cost per wear for the whole outfit
// OutfitCPW = outfit Price / outfit Wears
func (o *Outfit) calcOutfitCPW() {
	o.OutfitCPW = o.OutfitPrice / float32(o.OutfitWears)
}

// UpdateNumItems: designed to run at the creation or editing of an outfit
// updates variable numItems with count of all non-nil clothing pointers
func (o *Outfit) updateNumItems() {
	var articleCount uint
	list := []*Clothing{o.Top, o.Bottom, o.Shoes}
	articleCount = uint(len(o.Accessories))

	for _, article := range list {
		if article != nil {
			articleCount++
		}
	}
	o.numItems = articleCount
}

// calcAvgCPW: Updates AvgCPW variable with the average cost per wear of the indiviual items
// Sum(Each items cpw) / num items
func (o *Outfit) calcAvgCPW() {
	var accessoryCPW float32
	for _, article := range o.Accessories {
		accessoryCPW += article.CostPerWear
	}
	o.AvgCPW = (accessoryCPW + o.Top.CostPerWear + o.Bottom.CostPerWear + o.Shoes.CostPerWear) / float32(o.numItems)
}

func (o *Outfit) GetCPW() float32 {
	return o.OutfitCPW
}

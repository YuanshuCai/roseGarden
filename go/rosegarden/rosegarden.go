package rosegarden

type Item struct {
	Name            string
	SellIn, Quality int
}

// 4 types of items, 1 normal and 3 specials based on the assumptions
type NormalItem struct {
	*Item
}

func NewNomralItem(item *Item) *NormalItem {
	return &NormalItem{
		Item: item,
	}
}

type AgedBrie struct {
	*Item
}

type Suluras struct {
	*Item
}

type BackstagePass struct {
	*Item
}

func NewAgedBrie(item *Item) *AgedBrie {
	return &AgedBrie{
		Item: item,
	}
}
func NewSuluras(item *Item) *Suluras {
	return &Suluras{
		Item: item,
	}
}
func NewBackstagePass(item *Item) *BackstagePass {
	return &BackstagePass{
		Item: item,
	}
}

// create an interface that all items can use
type Updatable interface {
	Update()
}
type UpdatableItemCreation func(item *Item) Updatable

func MapItem(createClosure map[string]UpdatableItemCreation, item *Item) Updatable {
	create, exists := createClosure[item.Name]
	if exists {
		return create(item)
	} else {
		return NewNomralItem(item)
	}
}

// Update functions for the items
func (item *NormalItem) Update() {
	var decreaseRate = -1
	item.SellIn += decreaseRate
	if item.SellIn > 0 {
		item.Quality += decreaseRate
	} else {
		item.Quality += decreaseRate * 2
	}
	item.itemQualityRangeAdjuster()
}

func (item *AgedBrie) Update() {
	var decreaseRate = -1
	item.SellIn += decreaseRate
	item.Quality -= decreaseRate
	if item.Quality > 50 {
		item.Quality = 50
	}
	item.itemQualityRangeAdjuster()
}
func (item *Suluras) Update() {}

func (item *BackstagePass) Update() {
	var decreaseRate = -1
	item.SellIn += decreaseRate
	s := item.SellIn
	switch {
	case s > 10:
		item.Quality -= decreaseRate
	case s <= 10 && s > 5:
		item.Quality -= decreaseRate * 2
	case s <= 5 && s >= 0:
		item.Quality -= decreaseRate * 3
	case s < 0:
		item.Quality = 0
	}
	item.itemQualityRangeAdjuster()
}

func (item *Item) itemQualityRangeAdjuster() {
	if item.Quality < 0 {
		item.Quality = 0
	} else if item.Quality > 50 {
		item.Quality = 50
	}

}

/*
The origional prompt solution is more in Python Style
Because it is trying to use a for loop with pointer in the UpdateQuality func
and Golang's variables in for loop are values,
this will cause passing issues in the future.
And there are 3 special items which don't follow the regular item rules,
I suggest to use switch case instead of if else to make the code more readable
*/
func UpdateQuality(items []*Item) {
	creationMap := map[string]UpdatableItemCreation{
		"Aged Brie": func(item *Item) Updatable {
			return NewNomralItem(item)
		},
		"Sulfuras, Hand of Ragnaros": func(item *Item) Updatable {
			return NewSuluras(item)
		},
		"Backstage passes to a TAFKAL80ETC concert": func(item *Item) Updatable {
			return NewBackstagePass(item)
		},
	}

	for i := 0; i < len(items); i++ {

		updatableItem := MapItem(creationMap, items[i])
		updatableItem.Update()
	}
}

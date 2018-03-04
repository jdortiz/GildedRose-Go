package main

import "fmt"

type Item struct {
	name            string
	sellIn, quality int
}

var items = []Item{
	Item{"+5 Dexterity Vest", 10, 20},
	Item{"Aged Brie", 2, 0},
	Item{"Elixir of the Mongoose", 5, 7},
	Item{"Sulfuras, Hand of Ragnaros", 0, 80},
	Item{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	Item{"Conjured Mana Cake", 3, 6},
}

type Updatable interface {
	Update()
}

type RegularItem struct {
	*Item
}

func NewRegularItem(item *Item) *RegularItem {
	return &RegularItem{
		Item: item,
	}
}

type AgedBrieItem struct {
	*Item
}

func NewAgedBrieItem(item *Item) *AgedBrieItem {
	return &AgedBrieItem{
		Item: item,
	}
}

type SulfurasItem struct {
	*Item
}

func NewSulfurasItem(item *Item) *SulfurasItem {
	return &SulfurasItem{
		Item: item,
	}
}

type BackstagePassesItem struct {
	*Item
}

func NewBackstagePassesItem(item *Item) *BackstagePassesItem {
	return &BackstagePassesItem{
		Item: item,
	}
}

type UpdatableItemCreation func(item *Item) Updatable

func UpdatableItemFactory(createClosure map[string]UpdatableItemCreation, item *Item) Updatable {

	create, exists := createClosure[item.name]
	if exists {
		return create(item)
	} else {
		return NewRegularItem(item)
	}
}

func (item *RegularItem) Update() {
	item.changeSellIn(-1)
	if item.sellIn < 0 {
		item.changeQuality(-2)
	} else {
		item.changeQuality(-1)
	}
	item.limitQualityToMin(0)
}

func (item *AgedBrieItem) Update() {
	item.changeSellIn(-1)
	if item.sellIn < 0 {
		item.changeQuality(+2)
	} else {
		item.changeQuality(+1)
	}
	item.limitQualityToMax(50)
}

func (item *SulfurasItem) Update() {
}

func (item *BackstagePassesItem) Update() {
	item.changeSellIn(-1)
	sellIn := item.sellIn
	switch {
	case sellIn >= 10:
		item.changeQuality(+1)
	case sellIn < 10 && sellIn >= 5:
		item.changeQuality(+2)
	case sellIn < 15 && sellIn >= 0:
		item.changeQuality(+3)
	case sellIn < 0:
		item.limitQualityToMax(0)
	}
	item.limitQualityToMax(50)
}

func main() {
	fmt.Println("# Before updating")
	fmt.Println(items)
	UpdateInventory(items)
	fmt.Println("OMGHAI!")
	fmt.Println("# After updating")
	fmt.Println(items)
}

func UpdateInventory(items []Item) {
	creationMap := map[string]UpdatableItemCreation {
		"Aged Brie": func (item *Item) Updatable {
			return NewAgedBrieItem(item)
		},
		"Sulfuras, Hand of Ragnaros": func (item *Item) Updatable {
			return NewSulfurasItem(item)
		},
		"Backstage passes to a TAFKAL80ETC concert": func (item *Item) Updatable {
			return NewBackstagePassesItem(item)
		},
	}

	for i := 0; i < len(items); i++ {

		updatableItem := UpdatableItemFactory(creationMap, &items[i])
		updatableItem.Update()
	}
}

func (item *Item) changeSellIn(increment int) {
	item.sellIn += increment
}

func (item *Item) changeQuality(increment int) {
	item.quality += increment
}

func (item *Item) limitQualityToMax(max int) {
	if item.quality > max {
		item.quality = max
	}
}

func (item *Item) limitQualityToMin(min int) {
	if item.quality < min {
		item.quality = min
	}
}

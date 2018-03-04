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

type AgedBrieItem struct {
	*Item
}

type SulfurasItem struct {
	*Item
}

type BackstagePassesItem struct {
	*Item
}

func UpdatableItemFactory(item *Item) Updatable {
	switch item.name {
	case "Aged Brie":
		return &AgedBrieItem{
			Item: item,
		}
	case "Sulfuras, Hand of Ragnaros":
		return &SulfurasItem{
			Item: item,
		}
	case "Backstage passes to a TAFKAL80ETC concert":
		return &BackstagePassesItem{
			Item: item,
		}
	default:
		return &RegularItem{
			Item: item,
		}
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
	for i := 0; i < len(items); i++ {

		updatableItem := UpdatableItemFactory(&items[i])
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

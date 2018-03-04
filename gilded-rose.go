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

		switch items[i].name {
		case "Aged Brie":
			items[i].changeSellIn(-1)
			if items[i].sellIn < 0 {
				items[i].changeQuality(+2)
			} else {
				items[i].changeQuality(+1)
			}
			items[i].limitQualityToMax(50)

		case "Sulfuras, Hand of Ragnaros":

		case "Backstage passes to a TAFKAL80ETC concert":
			items[i].changeSellIn(-1)
			sellIn := items[i].sellIn
			switch {
			case sellIn >= 10:
				items[i].changeQuality(+1)
			case sellIn < 10 && sellIn >= 5:
				items[i].changeQuality(+2)
			case sellIn < 15 && sellIn >= 0:
				items[i].changeQuality(+3)
			case sellIn < 0:
				items[i].limitQualityToMax(0)
			}
			items[i].limitQualityToMax(50)

		default:
			items[i].changeSellIn(-1)
			if items[i].sellIn < 0 {
				items[i].changeQuality(-2)
			} else {
				items[i].changeQuality(-1)
			}
			items[i].limitQualityToMin(0)
		}
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

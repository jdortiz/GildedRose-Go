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
		case "Anything":
			changeSellIn(&items[i].sellIn, -1)
			if items[i].sellIn < 0 {
				changeQuality(&items[i].quality, -2)
			} else {
				changeQuality(&items[i].quality, -1)
			}
			limitToMin(&items[i].quality, 0)
		case "Aged Brie":
			changeSellIn(&items[i].sellIn, -1)
			if items[i].sellIn < 0 {
				changeQuality(&items[i].quality, +2)
			} else {
				changeQuality(&items[i].quality, +1)
			}
			limitToMax(&items[i].quality, 50)

		default:
		if items[i].name != "Aged Brie" && items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
			if items[i].quality > 0 {
				if items[i].name != "Sulfuras, Hand of Ragnaros" {
					items[i].quality = items[i].quality - 1
				}
			}
		} else {
			if items[i].quality < 50 {
				items[i].quality = items[i].quality + 1
				if items[i].name == "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].sellIn < 11 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
					if items[i].sellIn < 6 {
						if items[i].quality < 50 {
							items[i].quality = items[i].quality + 1
						}
					}
				}
			}
		}

		if items[i].name != "Sulfuras, Hand of Ragnaros" {
			items[i].sellIn = items[i].sellIn - 1
		}

		if items[i].sellIn < 0 {
			if items[i].name != "Aged Brie" {
				if items[i].name != "Backstage passes to a TAFKAL80ETC concert" {
					if items[i].quality > 0 {
						if items[i].name != "Sulfuras, Hand of Ragnaros" {
							items[i].quality = items[i].quality - 1
						}
					}
				} else {
					items[i].quality = items[i].quality - items[i].quality
				}
			} else {
				if items[i].quality < 50 {
					items[i].quality = items[i].quality + 1
				}
			}
		}

		}
	}
}

func changeSellIn(sellIn *int, increment int) {
	*sellIn += increment
}

func changeQuality(quality *int, increment int) {
	*quality += increment
}

func limitToMax(value *int, max int) {
	if *value > max {
		*value = max
	}
}

func limitToMin(value *int, min int) {
	if *value < min {
		*value = min
	}
}

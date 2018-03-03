package main

import "testing"

func TestSellInDecreasesOneUnit(t *testing.T) {
	items := []Item{
		Item{"Anything", 6, 9},
	}

	UpdateInventory(items)

	if items[0].sellIn != 5 {
		t.Errorf("SellIn (expected: %d, actual: %d).", 5, items[0].sellIn)
	}
}

func TestQualityDecreasesOneUnit(t *testing.T) {
	items := []Item{
		Item{"Anything", 6, 9},
	}

	UpdateInventory(items)

	if items[0].quality != 8 {
		t.Errorf("Quality (expected: %d, actual: %d).", 8, items[0].quality)
	}
}

func TestQualityDecreasesTwoUnitsAfterSellDate(t *testing.T) {
	items := []Item{
		Item{"Anything", 0, 9},
	}

	UpdateInventory(items)

	if items[0].quality != 7 {
		t.Errorf("Quality (expected: %d, actual: %d).", 7, items[0].quality)
	}
}

func TestQualityIsNeverNegative(t *testing.T) {
	items := []Item{
		Item{"Anything", 5, 0},
	}

	UpdateInventory(items)

	if items[0].quality != 0 {
		t.Errorf("Quality (expected: %d, actual: %d).", 0, items[0].quality)
	}
}

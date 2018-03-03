package main

import "testing"

func TestSellInDecreasesOneUnit(t *testing.T) {
	testSellIn(t, "Anything", 6, 9, 5)
}

func TestQualityDecreasesOneUnit(t *testing.T) {
	testQuality(t, "Anything", 6, 9, 8)
}

func TestQualityDecreasesTwoUnitsAfterSellDate(t *testing.T) {
	testQuality(t, "Anything", 0, 9, 7)
}

func TestQualityIsNeverNegative(t *testing.T) {
	testQuality(t, "Anything", 5, 0, 0)
}

func testSellIn(t *testing.T, itemName string, initiallSellIn int, initialQuality int, expectedSellIn int) {
	items := []Item{
		Item{itemName, initiallSellIn, initialQuality},
	}

	UpdateInventory(items)

	if items[0].sellIn != expectedSellIn {
		t.Errorf("SellIn (expected: %d, actual: %d).", expectedSellIn, items[0].sellIn)
	}
}

func testQuality(t *testing.T, itemName string, initiallSellIn int, initialQuality int, expectedQuality int) {
	items := []Item{
		Item{itemName, initiallSellIn, initialQuality},
	}

	UpdateInventory(items)

	if items[0].quality != expectedQuality {
		t.Errorf("Quality (expected: %d, actual: %d).", expectedQuality, items[0].quality)
	}
}

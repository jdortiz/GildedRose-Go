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

func TestAgedBrieSellInDecreasesOneUnit(t *testing.T) {
	testSellIn(t, "Aged Brie", 5, 10, 4)
}

func TestAgedBrieQualityIncreasesOneUnit(t *testing.T) {
	testQuality(t, "Aged Brie", 5, 10, 11)
}

func TestAgedBrieQualityIncreasesTwoUnitsAfterSellDate(t *testing.T) {
	testQuality(t, "Aged Brie", 0, 10, 12)
}

func TestAgedBrieQualityIsNeverMoreThanFifty(t *testing.T) {
	testQuality(t, "Aged Brie", 5, 50, 50)
}

func TestSulfurasNeverHasToBeSold(t *testing.T) {
	testSellIn(t, "Sulfuras, Hand of Ragnaros", 0, 80, 0)
}

func TestSulfurasQualityNeverDecreases(t *testing.T) {
	testQuality(t, "Sulfuras, Hand of Ragnaros", 0, 80, 80)
}

func TestBackstageSellInDecreasesOneUnit(t *testing.T) {
	testSellIn(t, "Backstage passes to a TAFKAL80ETC concert", 6, 10, 5)
}

func TestBackstagePassesQualityIncreasesOneUnitIfSellInMoreThanTen(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 11, 20, 21)
}

func TestBackstagePassesQualityIncreasesTwoUnitsIfSellInLessThanEleven(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 10, 20, 22)
}

func TestBackstagePassesQualityIncreasesTwoUnitsIfSellInMoreThanFive(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 6, 20, 22)
}

func TestBackstagePassesQualityIncreasesThreeUnitsIfSellInLessThanSix(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 5, 20, 23)
}

func TestBackstagePassesQualityIsZeroAfterSellDate(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 0, 20, 0)
}

func TestBackstagePassesQualityIsNeverMoreThanFiftyWhenIncreasingOne(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 11, 50, 50)
}

func TestBackstagePassesQualityIsNeverMoreThanFiftyWhenIncreasingTwo(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 7, 49, 50)
}

func TestBackstagePassesQualityIsNeverMoreThanFiftyWhenIncreasingThree(t *testing.T) {
	testQuality(t, "Backstage passes to a TAFKAL80ETC concert", 1, 49, 50)
}

func TestConjuredSellInDecreasesOneUnit(t *testing.T) {
	testSellIn(t, "Conjured", 8, 12, 7)
}

func TestConjuredQualityDecreasesTwoUnits(t *testing.T) {
	testQuality(t, "Conjured", 8, 12, 10)
}

func TestConjuredQualityDecreasesFourUnitsAfterSellDate(t *testing.T) {
	testQuality(t, "Conjured", 0, 12, 8)
}

func TestConjuredQualityIsNeverNegative(t *testing.T) {
	testQuality(t, "Conjured", 8, 0, 0)
}

// Auxiliary functions

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

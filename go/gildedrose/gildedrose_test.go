package gildedrose

import (
	"testing"
)

func TestIncreaseQuality(t *testing.T) {
	item := &Item{Name: "Test Item", SellIn: 10, Quality: 40}

	item.increaseQuality(5)
	if item.Quality != 45 {
		t.Errorf("Expected quality to be 45, got %d", item.Quality)
	}

	item.increaseQuality(10)
	if item.Quality != 50 {
		t.Errorf("Expected quality to cap at 50, got %d", item.Quality)
	}

	item.Name = "Sulfuras, Hand of Ragnaros"
	item.increaseQuality(5)
	if item.Quality != 50 {
		t.Errorf("Expected quality to remain unchanged for Sulfuras, got %d", item.Quality)
	}
}

func TestDecreaseSellIn(t *testing.T) {
	item := &Item{Name: "Test Item", SellIn: 10, Quality: 40}

	item.decreaseSellIn()
	if item.SellIn != 9 {
		t.Errorf("Expected SellIn to decrease to 9, got %d", item.SellIn)
	}

	item.Name = "Sulfuras, Hand of Ragnaros"
	item.decreaseSellIn()
	if item.SellIn != 9 {
		t.Errorf("Expected SellIn to remain unchanged for Sulfuras, got %d", item.SellIn)
	}
}

func TestDecreaseQuality(t *testing.T) {
	item := &Item{Name: "Test Item", SellIn: 10, Quality: 10}

	item.decreaseQuality(5)
	if item.Quality != 5 {
		t.Errorf("Expected quality to decrease to 5, got %d", item.Quality)
	}

	item.decreaseQuality(10)
	if item.Quality != 0 {
		t.Errorf("Expected quality to cap at 0, got %d", item.Quality)
	}

	item.Name = "Sulfuras, Hand of Ragnaros"
	item.decreaseQuality(5)
	if item.Quality != 0 {
		t.Errorf("Expected quality to remain unchanged for Sulfuras, got %d", item.Quality)
	}
}

func TestCanBeUpdated(t *testing.T) {
	item := &Item{Name: "Test Item", SellIn: 10, Quality: 10}
	if !item.canBeUpdated() {
		t.Errorf("Expected item to be updatable")
	}

	item.Name = "Sulfuras, Hand of Ragnaros"
	if item.canBeUpdated() {
		t.Errorf("Expected Sulfuras to not be updatable")
	}
}

func TestIsConjured(t *testing.T) {
	item := &Item{Name: "Conjured Mana Cake", SellIn: 10, Quality: 10}
	if !item.isConjured() {
		t.Errorf("Expected item to be conjured")
	}

	item.Name = "Normal Item"
	if item.isConjured() {
		t.Errorf("Expected item to not be conjured")
	}
}

func TestIsBackstage(t *testing.T) {
	item := &Item{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 10, Quality: 10}
	if !item.isBackstage() {
		t.Errorf("Expected item to be backstage passes")
	}

	item.Name = "Normal Item"
	if item.isBackstage() {
		t.Errorf("Expected item to not be backstage passes")
	}
}

func TestIsAgedBrie(t *testing.T) {
	item := &Item{Name: "Aged Brie", SellIn: 10, Quality: 10}
	if !item.isAgedBrie() {
		t.Errorf("Expected item to be Aged Brie")
	}

	item.Name = "Normal Item"
	if item.isAgedBrie() {
		t.Errorf("Expected item to not be Aged Brie")
	}
}

func TestUpdateQuality(t *testing.T) {
	items := []*Item{
		{Name: "Aged Brie", SellIn: 2, Quality: 0},
		{Name: "Backstage passes to a TAFKAL80ETC concert", SellIn: 15, Quality: 20},
		{Name: "Conjured Mana Cake", SellIn: 3, Quality: 6},
		{Name: "Normal Item", SellIn: 5, Quality: 10},
		{Name: "Sulfuras, Hand of Ragnaros", SellIn: 0, Quality: 80},
	}

	UpdateQuality(items)

	if items[0].Quality != 1 {
		t.Errorf("Expected Aged Brie quality to increase, got %d", items[0].Quality)
	}

	if items[1].Quality != 21 {
		t.Errorf("Expected Backstage passes quality to increase by 1, got %d", items[1].Quality)
	}

	if items[2].Quality != 4 {
		t.Errorf("Expected Conjured Mana Cake quality to decrease by 2, got %d", items[2].Quality)
	}

	if items[3].Quality != 9 {
		t.Errorf("Expected Normal Item quality to decrease by 1, got %d", items[3].Quality)
	}

	if items[4].Quality != 80 {
		t.Errorf("Expected Sulfuras quality to remain unchanged, got %d", items[4].Quality)
	}
}

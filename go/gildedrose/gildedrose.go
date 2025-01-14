// Package gildedrose provides functionality for managing items with varying quality
// and sell-in properties, adhering to specific rules for quality updates.
package gildedrose

import (
	"strings"
)

// Item represents an item with a name, sell-in period, and quality.
type Item struct {
	Name            string
	SellIn, Quality int
}

// increaseQuality increases the quality of the item by a specified amount.
// Quality cannot exceed 50 or be updated if the item cannot be updated.
func (item *Item) increaseQuality(quality int) {
	if (!item.canBeUpdated() || quality <= 0) {
		return
	}

	item.Quality += quality

	if (item.Quality >= 50) {
		item.Quality = 50
		return
	}
}

// decreaseSellIn reduces the sell-in value of the item by 1 if the item can be updated.
func (item *Item) decreaseSellIn() {
	if (item.canBeUpdated()) {
		item.SellIn -= 1
	}
}

// decreaseQuality decreases the quality of the item by a specified amount.
// Quality cannot go below 0 or be updated if the item cannot be updated.
func (item *Item) decreaseQuality(quality int) {
	if (!item.canBeUpdated() || quality <= 0) {
		return
	}

	item.Quality -= quality

	if (item.Quality <= 0) {
		item.Quality = 0
		return
	}
}

// canBeUpdated checks if the item can have its quality or sell-in values updated.
// Items named "Sulfuras, Hand of Ragnaros" cannot be updated.
func (item *Item) canBeUpdated() bool {
	return item.Name != "Sulfuras, Hand of Ragnaros"
}

// isConjured checks if the item is a conjured item.
// Conjured items contain "Conjured" in their name.
func (item *Item) isConjured() bool {
	return strings.Contains(item.Name, "Conjured")
}

// isBackstage checks if the item is a backstage pass.
// Backstage passes contain "Backstage passes" in their name.
func (item *Item) isBackstage() bool {
	return strings.Contains(item.Name, "Backstage passes")
}

// isAgedBrie checks if the item is "Aged Brie".
func (item *Item) isAgedBrie() bool {
	return item.Name == "Aged Brie"
}

// UpdateQuality updates the quality of a slice of items according to their rules.
// It iterates over the items, updates their quality, and decreases their sell-in values.
func UpdateQuality(items []*Item) {
	for i := 0; i < len(items); i++ {
		updateQuality(items[i])
		items[i].decreaseSellIn()
	}
}

// updateQuality is a helper function that updates the quality of a single item.
// Specific rules are applied based on the item's type.
func updateQuality(item *Item) {
	if !item.canBeUpdated() {
		return
	}

	if item.isBackstage() {
		backstageQualityUpdate(item)
		return
	}

	if item.isAgedBrie() {
		if item.SellIn > 0 {
			item.increaseQuality(1)
		} else {
			item.increaseQuality(2)
		}
		return
	}

	classicQualityUpdate(item)
}

// backstageQualityUpdate applies quality update rules specific to backstage passes.
// Quality increases based on the sell-in value and drops to 0 after the sell-in period.
func backstageQualityUpdate(item *Item) {
	if (item.SellIn > 10) {
		item.increaseQuality(1)
		return
	}

	if item.SellIn > 5 {
		item.increaseQuality(2)
		return
	}

	if item.SellIn > 0 {
		item.increaseQuality(3)
		return
	}
	item.Quality = 0
}

// classicQualityUpdate applies the standard quality update rules for an item.
// Conjured items and items past their sell-in date degrade faster.
func classicQualityUpdate(item *Item) {
	var decrease = 1

	if item.SellIn <= 0 {
		decrease = decrease * 2
	}

	if (item.isConjured()) {
		decrease = decrease * 2
	}

	item.decreaseQuality(decrease)
}

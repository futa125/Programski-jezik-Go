package someRandomPackage

import (
	"testing"
)

func TestTooExpensive(t *testing.T) {
	tooExpensiveItem := ShopItem{Name: "Banana", Price: 10, Type: "Prehrana", Callory: 10, Quantity: 3}

	shopItems := []ShopItem{tooExpensiveItem}

	totalCost, err := TotalCost(shopItems)

	if (totalCost == 0 && err == errTooExpensive) {
		return
	}

	t.Error("Item wasn't too expensive!")
}

func TestWrongType(t *testing.T) {
	wrongTypeItem := ShopItem{Name: "Jabuka", Price: 10, Type: "Sok", Callory: 10, Quantity: 1}

	shopItems := []ShopItem{wrongTypeItem}

	totalCost, err := TotalCost(shopItems)

	if (totalCost == 0 && err == errWrongItem) {
		return
	}

	t.Error("Item wasn't wrong type!")
}

func TestTooCaloric(t *testing.T) {
	tooCaloricItem := ShopItem{Name: "Sladoled", Price: 10, Type: "Prehrana", Callory: 350, Quantity: 1}

	shopItems := []ShopItem{tooCaloricItem}

	totalCost, err := TotalCost(shopItems)

	if (totalCost == 0 && err == errTooCaloric) {
		return
	}

	t.Error("Item wasn't too caloric!")
}

func TestValidItem(t *testing.T) {
	validItem := ShopItem{Name: "Ananas", Price: 15, Type: "Prehrana", Callory: 50, Quantity: 1}

	shopItems := []ShopItem{validItem}

	totalCost, err := TotalCost(shopItems)

	if (totalCost != 0 && err == nil) {
		return
	}

	t.Error("Item wasn't valid!")
}
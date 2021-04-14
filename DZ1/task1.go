package main

import (
		"errors"
		"fmt"
)

var errNoData = errors.New("no data in list")

type ShopItem struct {
	Name string
	Price int
	Callory int64
	Quantity int
}

func main() {
	shopList := []ShopItem{
		{"Hosomaki Ginger", 54, 250, 1},
		{"Mlijecna cokolada", 8, 600, 3},
		{"Banana", 10, 150, 6},
	}

	tCal := totalCallory(shopList)
	fmt.Println(tCal)

	best, err := mostHealthy(shopList)
	fmt.Println(best)
	fmt.Println(err)
}

func totalCallory(shoppingList []ShopItem) (total int) {
	for _, item := range shoppingList {
		total += int(item.Callory) * item.Quantity
	}
	return total
}

func mostHealthy(shoppingList []ShopItem) (items []ShopItem, err error) {
	if len(shoppingList) == 0 {
		return items, errNoData
	}
	
	min := shoppingList[0].Callory

	for _, item := range shoppingList {
		if item.Callory == min {
			items = append(items, item)
		} else if item.Callory < min {
			items = []ShopItem{item}
			min = item.Callory
		}
	}

	return items, nil
}
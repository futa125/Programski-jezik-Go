package someRandomPackage

import (
	"errors"
)

type ShopItem struct {
	Name     string
	Price    int
	Type     string
	Callory  int64
	Quantity int
}

var (
	errTooExpensive = errors.New("too expensive item")
	errWrongItem   = errors.New("item not allowed")
	errTooCaloric   = errors.New("too much calories")
)

const (
	ToExpensiveLimit = 20
	AllowedType      = "Prehrana"
)

func TotalCost(items []ShopItem) (int, error) {
	total := 0
	for _, v := range items {
		if v.Price*v.Quantity > 20 {
			return 0, errTooExpensive
		}
		if v.Type != AllowedType {
			return 0, errWrongItem
		}
		if v.Callory > 300 {
			return 0, errTooCaloric
		}
		total += v.Price * v.Quantity
	}
	return total, nil
}
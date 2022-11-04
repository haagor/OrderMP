package model

import (
	"errors"
	"strconv"
	"strings"
)

type Product struct {
	ProductID string
	Name      string
	Price     float64
}

// string înput format : product,product_id,price
func StringToProduct(s string) (Product, error) {
	l := strings.Split(s, ",")
	if len(l) != 3 {
		return Product{}, errors.New("not valid product")
	}

	price, err := strconv.ParseFloat(strings.TrimSuffix(l[2], "\n"), 64)
	if err != nil {
		return Product{}, err
	}

	p := Product{ProductID: l[0], Name: l[1], Price: price}
	return p, nil
}

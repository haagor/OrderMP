package model

import (
	"errors"
	"strconv"
	"strings"
)

type Product struct {
	id    string
	name  string
	price float64
}

// string înput format : product,product_id,price
func StringToProduct(s string) (Product, error) {
	l := strings.Split(s, ",")
	if len(l) != 3 {
		return Product{}, errors.New("not a valid product")
	}

	price, err := strconv.ParseFloat(strings.TrimSuffix(l[2], "\n"), 64)
	if err != nil {
		return Product{}, err
	}

	p := Product{id: l[0], name: l[1], price: price}
	return p, nil
}

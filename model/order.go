package model

import (
	"bufio"
	"errors"
	"strconv"
	"strings"
)

type Order struct {
	OrderID    string
	VAT        float64
	TotalPrice float64
	Products   []Product
}

/* string înput format :
Order: 123456
VAT: 3.10
Total: 16.90

product,product_id,price
Formule(s) midi,aZde,14.90
Café,IZ8z,2
*/
func StringToOrder(s string) (Order, error) {
	l := strings.Split(s, "\n\nproduct,product_id,price\n")
	if len(l) != 2 {
		return Order{}, errors.New("not valid order")
	}
	header := l[0]
	products := l[1]

	var headerValue []string
	lines := strings.Split(header, "\n")
	if len(lines) != 3 {
		return Order{}, errors.New("not valid header order")
	}

	for _, line := range lines {
		parsedLine := strings.Split(line, " ")
		if len(parsedLine) != 2 {
			return Order{}, errors.New("not valid header order")
		}
		headerValue = append(headerValue, parsedLine[1])
	}

	vat, err := strconv.ParseFloat(headerValue[1], 64)
	if err != nil {
		return Order{}, err
	}

	totalPrice, err := strconv.ParseFloat(headerValue[2], 64)
	if err != nil {
		return Order{}, err
	}

	order := Order{OrderID: headerValue[0], VAT: vat, TotalPrice: totalPrice}

	scanner := bufio.NewScanner(strings.NewReader(products))
	for scanner.Scan() {
		p, err := StringToProduct(scanner.Text())
		if err != nil {
			return Order{}, err
		}
		order.Products = append(order.Products, p)
	}

	return order, nil
}

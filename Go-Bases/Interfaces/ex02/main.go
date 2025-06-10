package main

import "fmt"

type Product interface {
	Price() float64
}

type SmallProduct struct {
	value float64
}

type MediumProduct struct {
	value     float64
	extraCost int
}

type BigProduct struct {
	value        float64
	extraCost    int
	shippingCost float64
}

func (s *SmallProduct) Price() float64 {
	return s.value
}

func (m *MediumProduct) Price() float64 {
	return m.value + (m.value * float64(m.extraCost) / 100.0)
}

func (b *BigProduct) Price() float64 {
	b.value = b.value + (b.value * float64(b.extraCost) / 100.0)
	return b.value + b.shippingCost
}

func showPriceOfProduct(p Product) {
	value := p.Price()
	fmt.Printf("Preço total: R$%f\n", value)
}

func main() {
	s := SmallProduct{
		value: 1000,
	}

	m := MediumProduct{
		value:     2200,
		extraCost: 6,
	}

	b := BigProduct{
		value:        5000,
		extraCost:    6,
		shippingCost: 2500,
	}

	showPriceOfProduct(&s)
	showPriceOfProduct(&m)
	showPriceOfProduct(&b)
}

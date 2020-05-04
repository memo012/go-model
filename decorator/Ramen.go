package decorator

type Ramen struct {
	Name string
	Price float32
}

func (p *Ramen) GetPrice() float32 {
	return p.Price
}

func (p *Ramen) Description() string {
	return p.Name
}

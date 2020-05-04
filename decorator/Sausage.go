package decorator

type Sausage struct {
	Noddle Noddles
	Name string
	Price float32
}

func (s *Sausage) SetNoddles(noddles Noddles)  {
	s.Noddle = noddles
}

func (s *Sausage) Description() string {
	return s.Noddle.Description() + "+" + s.Name
}

func (s *Sausage) GetPrice() float32 {
	return s.Noddle.GetPrice() + s.Price
}

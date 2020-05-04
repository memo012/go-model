package decorator

type Egg struct {
	Noddle Noddles
	Name string
	Price float32
}

func (e *Egg) SetNoddles(noddles Noddles)  {
	e.Noddle = noddles
}

func (e *Egg) Description() string {
	return e.Noddle.Description() + "+" + e.Name
}
func (e *Egg) GetPrice() float32 {
	return e.Noddle.GetPrice() + e.Price
}

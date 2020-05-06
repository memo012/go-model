package trains

type WeightTrains struct {
	name string
	weight int64
	currentWeight int64
}

func (w *WeightTrains) SetWeightTrains(name string, weight int64, currentWeight int64)  {
	w.weight = weight
	w.name = name
	w.currentWeight = currentWeight
}

func (w *WeightTrains) SetCurrentWeight(currentWeight int64)  {
	w.currentWeight = currentWeight
}

func (w *WeightTrains) GetCurrentWeight() int64 {
	return w.currentWeight
}



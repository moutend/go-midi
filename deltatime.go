package midi

// DeltaTime represents delta time .
type DeltaTime struct {
	quantity *Quantity
}

func (d *DeltaTime) Quantity() *Quantity {
	if d.quantity == nil {
		d.quantity = &Quantity{}
	}

	return d.quantity
}

func parseDeltaTime(stream []byte) (*DeltaTime, error) {
	q, err := parseQuantity(stream)
	if err != nil {
		return nil, err
	}

	deltaTime := &DeltaTime{q}

	return deltaTime, nil
}

func NewDeltaTime(value uint32) (*DeltaTime, error) {
	d := &DeltaTime{}
	err := d.Quantity().SetUint32(value)
	if err != nil {
		return nil, err
	}

	return d, nil
}

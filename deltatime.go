package midi

type DeltaTime struct {
	quantity *Quantity
}

func (d *DeltaTime) Quantity() *Quantity {
	if d.quantity == nil {
		d.quantity = &Quantity{}
	}

	return d.quantity
}

func NewDeltaTime(value uint32) (*DeltaTime, error) {
	d := &DeltaTime{}
	err := d.Quantity().SetUint32(value)
	if err != nil {
		return nil, err
	}

	return d, nil
}

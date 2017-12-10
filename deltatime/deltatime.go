package deltatime

import "github.com/moutend/go-midi/quantity"

// DeltaTime represents delta time .
type DeltaTime struct {
	value *quantity.Quantity
}

// Quantity returns variable length quantity of delta time.
func (d *DeltaTime) Quantity() *quantity.Quantity {
	if d.value == nil {
		d.value = &quantity.Quantity{}
	}

	return d.value
}

// Parse parses data.
func Parse(data []byte) (*DeltaTime, error) {
	q, err := quantity.Parse(data)
	if err != nil {
		return nil, err
	}

	deltaTime := &DeltaTime{
		value: q,
	}
	return deltaTime, nil
}

// New creates DeltaTime with given value.
func New(value int) (*DeltaTime, error) {
	d := &DeltaTime{}
	err := d.Quantity().SetUint32(uint32(value))
	if err != nil {
		return nil, err
	}

	return d, nil
}

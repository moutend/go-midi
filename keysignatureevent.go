package midi

import "fmt"

// KeySignatureEvent corresponds to key signature meta event.
type KeySignatureEvent struct {
	deltaTime *DeltaTime
	key       int8
	scale     uint8
}

// DeltaTime returns delta time of key signature event.
func (e *KeySignatureEvent) DeltaTime() *DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &DeltaTime{}
	}
	return e.deltaTime
}

// String returns string representation of key signature event.
func (e *KeySignatureEvent) String() string {
	return fmt.Sprintf("&KeySignatureEvent{key: %v, scale: %v}", e.key, e.scale)
}

// Serialize serializes key signature event.
func (e *KeySignatureEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, e.DeltaTime().Quantity().Value()...)
	bs = append(bs, Meta, KeySignature)
	bs = append(bs, 0x02, byte(e.key), e.scale)

	return bs
}

// SetKey sets key.
func (e *KeySignatureEvent) SetKey(key int8) error {
	if key < -7 || 7 < key {
		return fmt.Errorf("midi: range of key is -7 to 7")
	}
	e.key = key

	return nil
}

// Key returns key.
func (e *KeySignatureEvent) Key() int8 {
	return e.key
}

// SetScale sets scale.
func (e *KeySignatureEvent) SetScale(scale uint8) error {
	if scale > 1 {
		return fmt.Errorf("midi: scale must be 0 (major) or 1 (minor)")
	}
	e.scale = scale

	return nil
}

// Scale returns scale.
func (e *KeySignatureEvent) Scale() uint8 {
	return e.scale
}

// NewKeySignatureEvent returns KeySignatureEvent with the given parameter.
func NewKeySignatureEvent(deltaTime *DeltaTime, key int8, scale uint8) (*KeySignatureEvent, error) {
	var err error

	event := &KeySignatureEvent{}
	event.deltaTime = deltaTime

	err = event.SetKey(key)
	if err != nil {
		return nil, err
	}
	err = event.SetScale(scale)
	if err != nil {
		return nil, err
	}
	return event, nil
}

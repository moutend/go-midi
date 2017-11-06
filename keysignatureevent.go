package midi

import "fmt"

type KeySignatureEvent struct {
	deltaTime *DeltaTime
	tempo     []byte
}

func (e *KeySignatureEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *KeySignatureEvent) String() string {
	return fmt.Sprintf("&KeySignatureEvent{}")
}

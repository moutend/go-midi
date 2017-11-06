package midi

import "fmt"

type TextEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *TextEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *TextEvent) Text() string {
	return string(e.text)
}

func (e *TextEvent) String() string {
	return fmt.Sprintf("&TextEvent{text: \"%v\"}", string(e.text))
}

package midi

import "fmt"

type CopyrightNoticeEvent struct {
	deltaTime *DeltaTime
	text      []byte
}

func (e *CopyrightNoticeEvent) DeltaTime() *DeltaTime {
	return e.deltaTime
}

func (e *CopyrightNoticeEvent) String() string {
	return fmt.Sprintf("&CopyrightNoticeEvent{text: \"%v\"}", string(e.text))
}

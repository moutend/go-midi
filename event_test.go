package midi

import "testing"

func TestParseEvent(t *testing.T) {
	stream := []byte{0x00, 0xff, 0x02, 0x12, 0x43, 0x6f, 0x70, 0x79, 0x72, 0x69, 0x67, 0x68, 0x74, 0x20, 0x28, 0x43, 0x29, 0x20, 0x32, 0x30, 0x31, 0x37}
	event, sizeOfEvent, err := parseEvent(stream)
	if err != nil {
		t.Fatal(err)
	}
	if sizeOfEvent != 22 {
		t.Fatalf("expected: size of event = 22, actual: size of event = %v", sizeOfEvent)
	}
	if event.DeltaTime().Quantity().Uint32() != 0 {
		t.Fatalf("expected: 0 actual: %v", event.DeltaTime().Quantity().Uint32())
	}
	switch event.(type) {
	case *CopyrightNoticeEvent:
		break
	default:
		t.Fatalf("type of event must be CopyrightNoticeEvent")
	}

	expectedText := stream[4:]
	actualText := event.(*CopyrightNoticeEvent).text

	if len(expectedText) != len(actualText) {
		t.Fatalf("expect: len(event.(*CopyrightNoticeEvent).text) = %v actual: len(event.(*CopyrightNoticeEvent).text) = %v", len(expectedText), len(actualText))
	}
	for i, v := range expectedText {
		if v != actualText[i] {
			t.Fatalf("expected: text[%v] = %v actual: text[%v] = %v", i, v, i, actualText[i])
		}
	}
}

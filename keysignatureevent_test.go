package midi

import "testing"

func TestKeySignatureEvent_DeltaTime(t *testing.T) {
	event := &KeySignatureEvent{}
	dt := event.DeltaTime()
	if dt == nil {
		t.Fatal("DeltaTime() don't return nil")
	}
}

func TestKeySignatureEvent_String(t *testing.T) {
	event, err := NewKeySignatureEvent(nil, -7, 1)
	if err != nil {
		t.Fatal(err)
	}

	expected := "&KeySignatureEvent{key: -7, scale: 1}"
	actual := event.String()
	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestKeySignatureEvent_Serialize(t *testing.T) {
	event, err := NewKeySignatureEvent(nil, -7, 1)
	if err != nil {
		t.Fatal(err)
	}

	expected := []byte{0x00, 0xff, 0x59, 0x02, 0xf9, 0x01}
	actual := event.Serialize()

	if len(expected) != len(actual) {
		t.Fatalf("expected: %v bytes actual: %v bytes", len(expected), len(actual))
	}
	for i, e := range expected {
		a := actual[i]
		if e != a {
			t.Fatalf("expected[%v] = 0x%x actual[%v] = 0x%x", i, e, i, a)
		}
	}
}

func TestKeySignatureEvent_SetKey(t *testing.T) {
	var err error

	event := &KeySignatureEvent{}

	err = event.SetKey(-8)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetKey(8)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetKey(-7)
	if err != nil {
		t.Fatal(err)
	}
	err = event.SetKey(7)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeySignatureEvent_Key(t *testing.T) {
	event := &KeySignatureEvent{key: -7}

	expected := int8(-7)
	actual := event.Key()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestKeySignatureEvent_SetScale(t *testing.T) {
	var err error

	event := &KeySignatureEvent{}

	err = event.SetScale(2)
	if err == nil {
		t.Fatalf("err must not be nil")
	}
	err = event.SetScale(1)
	if err != nil {
		t.Fatal(err)
	}
}

func TestKeySignatureEvent_Scale(t *testing.T) {
	event := &KeySignatureEvent{scale: 1}

	expected := uint8(1)
	actual := event.Scale()

	if expected != actual {
		t.Fatalf("expected: %v actual: %v", expected, actual)
	}
}

func TestNewKeySignatureEvent(t *testing.T) {
	_, err := NewKeySignatureEvent(nil, -8, 1)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	_, err = NewKeySignatureEvent(nil, -7, 255)
	if err == nil {
		t.Fatalf("err must not be nil")
	}

	event, err := NewKeySignatureEvent(nil, -7, 1)
	if err != nil {
		t.Fatal(err)
	}
	if event.key != -7 {
		t.Fatalf("expected: -7 actual: %v", event.key)
	}
	if event.scale != 1 {
		t.Fatalf("expected: 1 actual: %v", event.scale)
	}
}

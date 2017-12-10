package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// TimeSignatureEvent corresponds to time signature meta event.
type TimeSignatureEvent struct {
	deltaTime      *deltatime.DeltaTime
	runningStatus  bool
	numerator      uint8
	denominator    uint8
	metronomePulse uint8
	quarterNote    uint8
}

// deltatime.DeltaTime returns delta time of time signature event.
func (e *TimeSignatureEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes time signature event.
func (e *TimeSignatureEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.TimeSignature)
	bs = append(bs, 0x04, e.numerator, e.denominator, e.metronomePulse, e.quarterNote)

	return bs
}

// SetRunningStatus sets running status.
func (e *TimeSignatureEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *TimeSignatureEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetNumerator sets numerator.
func (e *TimeSignatureEvent) SetNumerator(numerator uint8) error {
	e.numerator = numerator

	return nil
}

// Numerator returns numerator.
func (e *TimeSignatureEvent) Numerator() uint8 {
	return e.numerator
}

// SetDenominator sets denominator.
func (e *TimeSignatureEvent) SetDenominator(denominator uint8) error {
	e.denominator = denominator

	return nil
}

// Denominator returns denominator.
func (e *TimeSignatureEvent) Denominator() uint8 {
	return e.denominator
}

// SetMetronomePulse sets metronomePulse.
func (e *TimeSignatureEvent) SetMetronomePulse(metronomePulse uint8) error {
	e.metronomePulse = metronomePulse

	return nil
}

// MetronomePulse returns metronomePulse.
func (e *TimeSignatureEvent) MetronomePulse() uint8 {
	return e.metronomePulse
}

// SetQuarterNote sets quarterNote.
func (e *TimeSignatureEvent) SetQuarterNote(quarterNote uint8) error {
	e.quarterNote = quarterNote

	return nil
}

// QuarterNote returns quarterNote.
func (e *TimeSignatureEvent) QuarterNote() uint8 {
	return e.quarterNote
}

// String returns string representation of time signature event.
func (e *TimeSignatureEvent) String() string {
	return fmt.Sprintf("&TimeSignatureEvent{numerator: %v, denominator: %v, metronomePulse: %v, quarterNote: %v}", e.numerator, e.denominator, e.metronomePulse, e.quarterNote)
}

// NewTimeSignatureEvent returns TimeSignatureEvent with the given parameter.
func NewTimeSignatureEvent(deltaTime *deltatime.DeltaTime, numerator, denominator, metronomePulse, quarterNote uint8) (*TimeSignatureEvent, error) {
	event := &TimeSignatureEvent{
		deltaTime:      deltaTime,
		numerator:      numerator,
		denominator:    denominator,
		metronomePulse: metronomePulse,
		quarterNote:    quarterNote,
	}
	return event, nil
}

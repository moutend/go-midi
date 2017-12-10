package event

import (
	"fmt"

	"github.com/moutend/go-midi/constant"
	"github.com/moutend/go-midi/deltatime"
)

// SMPTEOffsetEvent corresponds to SMPTE offset event.
type SMPTEOffsetEvent struct {
	deltaTime     *deltatime.DeltaTime
	runningStatus bool
	hour          uint8
	minute        uint8
	second        uint8
	frame         uint8
	subFrame      uint8
}

// deltatime.DeltaTime returns delta time of SMPTE offset event.
func (e *SMPTEOffsetEvent) DeltaTime() *deltatime.DeltaTime {
	if e.deltaTime == nil {
		e.deltaTime = &deltatime.DeltaTime{}
	}
	return e.deltaTime
}

// Serialize serializes SMPTE offset event.
func (e *SMPTEOffsetEvent) Serialize() []byte {
	bs := []byte{}
	bs = append(bs, constant.Meta, constant.SMPTEOffset)
	bs = append(bs, 0x05, e.hour, e.minute, e.second, e.frame, e.subFrame)

	return bs
}

// SetRunningStatus sets running status.
func (e *SMPTEOffsetEvent) SetRunningStatus(status bool) {
	e.runningStatus = status
}

// RunningStatus returns running status.
func (e *SMPTEOffsetEvent) RunningStatus() bool {
	return e.runningStatus
}

// SetHour sets hour.
func (e *SMPTEOffsetEvent) SetHour(hour uint8) error {
	if hour > 23 {
		return fmt.Errorf("midi: hour is 0 to 23")
	}
	e.hour = hour

	return nil
}

// Hour returns hour.
func (e *SMPTEOffsetEvent) Hour() uint8 {
	return e.hour
}

// SetMinute sets minute.
func (e *SMPTEOffsetEvent) SetMinute(minute uint8) error {
	if minute > 59 {
		return fmt.Errorf("midi: minute is 0 to 59")
	}
	e.minute = minute

	return nil
}

// Minute returns minute.
func (e *SMPTEOffsetEvent) Minute() uint8 {
	return e.minute
}

// SetSecond sets second.
func (e *SMPTEOffsetEvent) SetSecond(second uint8) error {
	if second > 59 {
		return fmt.Errorf("midi: second is 0 to 59")
	}
	e.second = second

	return nil
}

// Second returns second.
func (e *SMPTEOffsetEvent) Second() uint8 {
	return e.second
}

// SetFrame sets frame.
func (e *SMPTEOffsetEvent) SetFrame(frame uint8) error {
	if frame > 30 {
		return fmt.Errorf("midi: frame is 0 to 30")
	}
	e.frame = frame

	return nil
}

// Frame returns frame.
func (e *SMPTEOffsetEvent) Frame() uint8 {
	return e.frame
}

// SetSubFrame sets sub frame.
func (e *SMPTEOffsetEvent) SetSubFrame(subFrame uint8) error {
	if subFrame > 99 {
		return fmt.Errorf("midi: subFrame is 0 to 99")
	}
	e.subFrame = subFrame

	return nil
}

// SubFrame returns sub frame.
func (e *SMPTEOffsetEvent) SubFrame() uint8 {
	return e.subFrame
}

// String returns string representation of SMPTE offset event.
func (e *SMPTEOffsetEvent) String() string {
	return fmt.Sprintf("&SMPTEOffsetEvent{hour: %v, minute: %v, second: %v, frame: %v, subFrame: %v}", e.hour, e.minute, e.second, e.frame, e.subFrame)
}

// NewSMPTEOffsetEvent returns SMPTEOffsetEvent with the given parameter.
func NewSMPTEOffsetEvent(deltaTime *deltatime.DeltaTime, hour, minute, second, frame, subFrame uint8) (*SMPTEOffsetEvent, error) {
	var err error

	event := &SMPTEOffsetEvent{}
	event.deltaTime = deltaTime
	err = event.SetHour(hour)
	if err != nil {
		return nil, err
	}
	err = event.SetMinute(minute)
	if err != nil {
		return nil, err
	}
	err = event.SetSecond(second)
	if err != nil {
		return nil, err
	}
	err = event.SetFrame(frame)
	if err != nil {
		return nil, err
	}
	err = event.SetSubFrame(subFrame)
	if err != nil {
		return nil, err
	}
	return event, nil
}

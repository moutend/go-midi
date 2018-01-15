//go:generate stringer -type=Control -output=control_string.go

package constant

import (
	"fmt"
	"strconv"
	"strings"
)

// Control represents MIDI controls such as main volume, modulation and etc.
type Control uint8

const (
	BankSelect                      Control = 0x00
	Modulation                      Control = 0x01
	BreathController                Control = 0x02
	FootController                  Control = 0x04
	PortamentoTime                  Control = 0x05
	DataEntry                       Control = 0x06
	MainVolume                      Control = 0x07
	Balance                         Control = 0x08
	Pan                             Control = 0x0a
	Expression                      Control = 0x0b
	EffectControl1                  Control = 0x0c
	EffectControl2                  Control = 0x0d
	GeneralPurposeController1       Control = 0x10
	GeneralPurposeController2       Control = 0x11
	GeneralPurposeController3       Control = 0x12
	GeneralPurposeController4       Control = 0x13
	BankSelectLSB                   Control = 0x20
	ModulationLSB                   Control = 0x21
	BreathControllerLSB             Control = 0x22
	FootControllerLSB               Control = 0x24
	PortamentoTimeLSB               Control = 0x25
	DataEntryLSB                    Control = 0x26
	MainVolumeLSB                   Control = 0x27
	BalanceLSB                      Control = 0x28
	PanLSB                          Control = 0x2a
	ExpressionLSB                   Control = 0x2b
	EffectControl1LSB               Control = 0x2c
	EffectControl2LSB               Control = 0x2d
	GeneralPurposeController1LSB    Control = 0x30
	GeneralPurposeController2LSB    Control = 0x31
	GeneralPurposeController3LSB    Control = 0x32
	GeneralPurposeController4LSB    Control = 0x33
	Hold1                           Control = 0x40
	PortamentoOnOff                 Control = 0x41
	Sostenuto                       Control = 0x42
	SoftPedal                       Control = 0x43
	LegatoFootswitch                Control = 0x44
	Hold2                           Control = 0x45
	SoundVariation                  Control = 0x46
	HarmonicIntensity               Control = 0x47
	ReleaseTime                     Control = 0x48
	AttackTime                      Control = 0x49
	Brightness                      Control = 0x4a
	DecayTime                       Control = 0x4b
	VibratoRate                     Control = 0x4c
	VibratoDepth                    Control = 0x4d
	VibratoDelay                    Control = 0x4e
	UndefinedSoundController        Control = 0x4f
	GeneralPurposeController5       Control = 0x50
	GeneralPurposeController6       Control = 0x51
	GeneralPurposeController7       Control = 0x52
	GeneralPurposeController8       Control = 0x53
	PortamentoControl               Control = 0x54
	ReverbSendLevel                 Control = 0x5b
	TremoloDepth                    Control = 0x5c
	ChorusSendLevel                 Control = 0x5d
	CelesteDepth                    Control = 0x5e
	PhaserDepth                     Control = 0x5f
	DataIncrement                   Control = 0x60
	DataDecrement                   Control = 0x61
	NonRegisteredParameterNumberLSB Control = 0x62
	NonRegisteredParameterNumberMSB Control = 0x63
	RegisteredParameterNumberLSB    Control = 0x64
	RegisteredParameterNumberMSB    Control = 0x65
)

func ParseControlName(s string) (Control, error) {
	s = strings.ToLower(s)
	i, err := strconv.Atoi(s)
	if err == nil {
		return Control(i), nil
	}
	return 0, fmt.Errorf("invalid")
}

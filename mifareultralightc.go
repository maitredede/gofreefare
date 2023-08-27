package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareUltralightC struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareUltralightC)(nil)

func (t *MifareUltralightC) Device() gonfc.Device {
	return t.device
}

func (MifareUltralightC) Type() TagType {
	return MIFARE_ULTRALIGHT_C
}

func mifareUltralightCTaste(device gonfc.Device, target gonfc.Target) (*MifareUltralightC, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x00 {
		return nil, false
	}

	tag := &MifareUltralightC{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

func isMifareUltralightcOnReader(device gonfc.Device, target *gonfc.ISO14443aTarget) bool {
	mod := gonfc.Modulation{
		Type:     gonfc.ISO14443a,
		BaudRate: gonfc.Nbr106,
	}
	initData := target.UID[:]
	if _, err := device.InitiatorSelectPassiveTarget(mod, initData); err != nil {
		panic(err)
	}
	if err := device.SetPropertyBool(gonfc.EasyFraming, false); err != nil {
		panic(err)
	}
	cmdStep1 := []byte{0x1A, 0x00}
	resStep1 := make([]byte, 9)
	ret, err := device.InitiatorTransceiveBytes(cmdStep1, resStep1, 0)
	if err != nil {
		panic(err)
	}
	if err := device.SetPropertyBool(gonfc.EasyFraming, true); err != nil {
		panic(err)
	}
	if err := device.InitiatorDeselectTarget(); err != nil {
		panic(err)
	}
	return ret > 0
}

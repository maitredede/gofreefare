package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareUltralight struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareUltralight)(nil)

func (t *MifareUltralight) Device() gonfc.Device {
	return t.device
}

func (MifareUltralight) Type() TagType {
	return MIFARE_ULTRALIGHT
}

func mifareUltralightTaste(device gonfc.Device, target gonfc.Target) (*MifareUltralight, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x00 {
		return nil, false
	}

	tag := &MifareUltralight{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

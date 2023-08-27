package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareClassic4k struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareClassic4k)(nil)

func (t *MifareClassic4k) Device() gonfc.Device {
	return t.device
}

func (MifareClassic4k) Type() TagType {
	return MIFARE_CLASSIC_4K
}

func mifareClassic4kTaste(device gonfc.Device, target gonfc.Target) (*MifareClassic4k, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x18 && mf.Sak != 0x38 {
		return nil, false

	}

	tag := &MifareClassic4k{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

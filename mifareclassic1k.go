package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareClassic1k struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareClassic1k)(nil)

func (t *MifareClassic1k) Device() gonfc.Device {
	return t.device
}

func (MifareClassic1k) Type() TagType {
	return MIFARE_CLASSIC_1K
}

func mifareClassic1kTaste(device gonfc.Device, target gonfc.Target) (*MifareClassic1k, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x08 && mf.Sak != 0x28 && mf.Sak != 0x68 && mf.Sak != 0x88 {
		return nil, false

	}

	tag := &MifareClassic1k{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

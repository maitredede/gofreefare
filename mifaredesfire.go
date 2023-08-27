package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareDesfire struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareDesfire)(nil)

func (t *MifareDesfire) Device() gonfc.Device {
	return t.device
}

func (MifareDesfire) Type() TagType {
	return MIFARE_DESFIRE
}

func mifareDesfireTaste(device gonfc.Device, target gonfc.Target) (*MifareDesfire, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x18 && mf.Sak != 0x38 {
		return nil, false

	}

	tag := &MifareDesfire{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

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
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareClassic4k)(nil)

func (t *MifareClassic4k) Device() gonfc.Device {
	return t.device
}

func (MifareClassic4k) Type() TagType {
	return MIFARE_CLASSIC_4K
}

func mifareClassic4kTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareClassic4k, bool) {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return nil, false
	}
	btSak := target.NTI.NAI().BtSak
	if btSak != 0x18 && btSak != 0x38 {
		return nil, false

	}

	tag := &MifareClassic4k{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

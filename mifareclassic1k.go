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
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareClassic1k)(nil)

func (t *MifareClassic1k) Device() gonfc.Device {
	return t.device
}

func (MifareClassic1k) Type() TagType {
	return MIFARE_CLASSIC_1K
}

func mifareClassic1kTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareClassic1k, bool) {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return nil, false
	}
	btSak := target.NTI.NAI().BtSak
	if btSak != 0x08 && btSak != 0x28 && btSak != 0x68 && btSak != 0x88 {
		return nil, false

	}

	tag := &MifareClassic1k{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

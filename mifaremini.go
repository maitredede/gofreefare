package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareMini struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareMini)(nil)

func (t *MifareMini) Device() gonfc.Device {
	return t.device
}

func (MifareMini) Type() TagType {
	return MIFARE_MINI
}

func mifareMiniTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareMini, bool) {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return nil, false
	}
	if target.NTI.NAI().BtSak != 0x09 {
		return nil, false
	}

	tag := &MifareMini{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

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
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareUltralight)(nil)

func (t *MifareUltralight) Device() gonfc.Device {
	return t.device
}

func (MifareUltralight) Type() TagType {
	return MIFARE_ULTRALIGHT
}

func mifareUltralightTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareUltralight, bool) {
	if !taste(target) {
		return nil, false
	}
	isC, err := isMifareUltralightcOnReader(device, target.NTI.NAI())
	if err != nil {
		return nil, false
	}
	if isC {
		return nil, false
	}

	tag := &MifareUltralight{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

func taste(target *gonfc.NfcTarget) bool {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return false
	}
	btSak := target.NTI.NAI().BtSak
	return btSak == 0x00
}

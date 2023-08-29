package gofreefare

import (
	"bytes"
	"time"

	"github.com/maitredede/gonfc"
)

type MifareDesfire struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareDesfire)(nil)

func (t *MifareDesfire) Device() gonfc.Device {
	return t.device
}

func (MifareDesfire) Type() TagType {
	return MIFARE_DESFIRE
}

func mifareDesfireTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareDesfire, bool) {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return nil, false
	}
	btSak := target.NTI.NAI().BtSak
	if btSak != 0x20 {
		return nil, false
	}

	// We have three different ATS prefixes to
	// check for, standalone, JCOP and JCOP3.
	STANDALONE_DESFIRE := []byte{0x75, 0x77, 0x81, 0x02}
	JCOP_DESFIRE := []byte{0x75, 0xf7, 0xb1, 0x02}
	JCOP3_DESFIRE := []byte{0x78, 0x77, 0x71, 0x02}

	ats := target.NTI.NAI().ATS()
	ok := false
	if bytes.Equal(ats, STANDALONE_DESFIRE) {
		ok = true
	}
	if bytes.Equal(ats, JCOP_DESFIRE) {
		ok = true
	}
	if bytes.Equal(ats, JCOP3_DESFIRE) {
		ok = true
	}
	if !ok {
		return nil, false
	}

	tag := &MifareDesfire{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

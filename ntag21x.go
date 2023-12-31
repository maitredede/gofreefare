package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type NTag21x struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*NTag21x)(nil)

func (t *NTag21x) Device() gonfc.Device {
	return t.device
}

func (NTag21x) Type() TagType {
	return NTAG_21x
}

func nTag21xTaste(device gonfc.Device, target *gonfc.NfcTarget) (*NTag21x, bool) {
	if target.NM.Type != gonfc.NMT_ISO14443A {
		return nil, false
	}
	btSak := target.NTI.NAI().BtSak
	if btSak != 0x00 {
		return nil, false
	}

	tag := &NTag21x{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

func nTag21xIsAuthSupported(device gonfc.Device, nai *gonfc.NfcIso14443aInfo) bool {
	mod := gonfc.Modulation{
		Type:     gonfc.ISO14443a,
		BaudRate: gonfc.Nbr106,
	}
	initData := nai.UID()
	_, err := device.InitiatorSelectPassiveTarget(mod, initData)
	if err != nil {
		panic(err)
	}
	if err := device.SetPropertyBool(gonfc.EasyFraming, false); err != nil {
		panic(err)
	}
	cmdStep1 := []byte{0x60}
	resStep1 := make([]byte, 8)
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

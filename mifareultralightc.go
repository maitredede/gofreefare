package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type MifareUltralightC struct {
	device gonfc.Device
	// 	target  gonfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*MifareUltralightC)(nil)

func (t *MifareUltralightC) Device() gonfc.Device {
	return t.device
}

func (MifareUltralightC) Type() TagType {
	return MIFARE_ULTRALIGHT_C
}

func mifareUltralightCTaste(device gonfc.Device, target *gonfc.NfcTarget) (*MifareUltralightC, bool) {
	if !taste(target) {
		return nil, false
	}
	isC, err := isMifareUltralightcOnReader(device, target.NTI.NAI())
	if err != nil {
		return nil, false
	}
	if !isC {
		return nil, false
	}

	tag := &MifareUltralightC{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

func isMifareUltralightcOnReader(device gonfc.Device, nai *gonfc.NfcIso14443aInfo) (bool, error) {
	mod := gonfc.Modulation{
		Type:     gonfc.ISO14443a,
		BaudRate: gonfc.Nbr106,
	}
	initData := nai.UID()
	if _, err := device.InitiatorSelectPassiveTarget(mod, initData); err != nil {
		return false, err
	}
	if err := device.SetPropertyBool(gonfc.EasyFraming, false); err != nil {
		return false, err
	}
	cmdStep1 := []byte{0x1A, 0x00}
	resStep1 := make([]byte, 9)
	ret, err := device.InitiatorTransceiveBytes(cmdStep1, resStep1, 0)
	if err != nil {
		return false, err
	}
	if err := device.SetPropertyBool(gonfc.EasyFraming, true); err != nil {
		return false, err
	}
	if err := device.InitiatorDeselectTarget(); err != nil {
		return false, err
	}
	return ret > 0, nil
}

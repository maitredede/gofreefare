package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

type FelicaTag struct {
	device gonfc.Device
	// 	target  libnfc.Target
	// 	tagtype int
	active int
	timeout/*int*/ time.Duration
	target *gonfc.NfcTarget
}

var _ FreefareTag = (*FelicaTag)(nil)

func (FelicaTag) Type() TagType {
	return FELICA
}

func (t *FelicaTag) Device() gonfc.Device {
	return t.device
}

func felicaTaste(device gonfc.Device, target *gonfc.NfcTarget) (*FelicaTag, bool) {
	if target.NM.Type != gonfc.Felica {
		return nil, false
	}

	tag := &FelicaTag{
		device:  device,
		target:  target,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

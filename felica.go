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
	target *gonfc.FelicaTarget
}

var _ FreefareTag = (*FelicaTag)(nil)

func (FelicaTag) Type() TagType {
	return FELICA
}

func (t *FelicaTag) Device() gonfc.Device {
	return t.device
}

func felicaTaste(device gonfc.Device, target gonfc.Target) (*FelicaTag, bool) {
	fc, ok := target.(*gonfc.FelicaTarget)
	if !ok {
		return nil, false
	}

	tag := &FelicaTag{
		device:  device,
		target:  fc,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

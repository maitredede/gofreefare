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
	target *gonfc.ISO14443aTarget
}

var _ FreefareTag = (*MifareMini)(nil)

func (t *MifareMini) Device() gonfc.Device {
	return t.device
}

func (MifareMini) Type() TagType {
	return MIFARE_MINI
}

func mifareMiniTaste(device gonfc.Device, target gonfc.Target) (*MifareMini, bool) {
	mf, ok := target.(*gonfc.ISO14443aTarget)
	if !ok {
		return nil, false
	}
	if mf.Sak != 0x09 {
		return nil, false
	}

	tag := &MifareMini{
		device:  device,
		target:  mf,
		active:  0,
		timeout: MIFARE_DEFAULT_TIMEOUT,
	}
	return tag, true
}

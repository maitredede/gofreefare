package gofreefare

import (
	"time"

	"github.com/maitredede/gonfc"
)

const (
	MIFARE_DEFAULT_TIMEOUT = 2000 * time.Millisecond
)

func freefareTagNew(device gonfc.Device, target gonfc.Target) FreefareTag {

	if fc, ok := felicaTaste(device, target); ok {
		return fc
	}
	if mm, ok := mifareMiniTaste(device, target); ok {
		return mm
	}
	if mc1, ok := mifareClassic1kTaste(device, target); ok {
		return mc1
	}
	if mc4, ok := mifareClassic4kTaste(device, target); ok {
		return mc4
	}
	if md, ok := mifareDesfireTaste(device, target); ok {
		return md
	}
	if nt, ok := nTag21xTaste(device, target); ok {
		return nt
	}
	if muc, ok := mifareUltralightCTaste(device, target); ok {
		return muc
	}
	if mu, ok := mifareUltralightTaste(device, target); ok {
		return mu
	}
	return nil
}

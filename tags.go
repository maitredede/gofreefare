package gofreefare

import "github.com/maitredede/gonfc"

type FreefareTag interface {
	Device() gonfc.Device
	Type() TagType
}

type TagType int

const (
	FELICA TagType = iota
	MIFARE_MINI
	MIFARE_CLASSIC_1K
	MIFARE_CLASSIC_4K
	MIFARE_DESFIRE
	//    MIFARE_PLUS_S2K
	//    MIFARE_PLUS_S4K
	//    MIFARE_PLUS_X2K
	//    MIFARE_PLUS_X4K
	MIFARE_ULTRALIGHT
	MIFARE_ULTRALIGHT_C
	NTAG_21x
)

// type FreefareTag struct {
// 	device  gonfc.Device
// 	target  gonfc.Target
// 	tagtype int
// 	active  int
// 	timeout/*int*/ time.Duration
// }

// type MifareClassic struct {
// 	FreefareTag
// }

// type MifareDesfire struct {
// 	FreefareTag
// }

// type MifareUltralight struct {
// 	FreefareTag
// }

// type Ntag21x struct {
// 	FreefareTag
// }

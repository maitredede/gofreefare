package gofreefare

import (
	"fmt"

	"github.com/maitredede/gonfc"
)

func GetTags(device gonfc.Device) ([]FreefareTag, error) {
	var tags []FreefareTag

	if err := device.InitiatorInit(); err != nil {
		return nil, fmt.Errorf("initiator_init: %w", err)
	}

	// Drop the field for a while
	if err := device.SetPropertyBool(gonfc.ActivateField, false); err != nil {
		return nil, fmt.Errorf("field=false: %w", err)
	}

	// Configure the CRC and Parity settings
	if err := device.SetPropertyBool(gonfc.HandleCRC, true); err != nil {
		return nil, fmt.Errorf("crc=true: %w", err)
	}
	if err := device.SetPropertyBool(gonfc.HandleParity, true); err != nil {
		return nil, fmt.Errorf("parity=true: %w", err)
	}
	if err := device.SetPropertyBool(gonfc.AutoISO14443_4, true); err != nil {
		return nil, fmt.Errorf("autoISO=true: %w", err)
	}

	// Enable field so more power consuming cards can power themselves up
	if err := device.SetPropertyBool(gonfc.ActivateField, true); err != nil {
		return nil, fmt.Errorf("field=true: %w", err)
	}

	// Poll for a ISO14443A (MIFARE) tag
	modulation := gonfc.Modulation{
		Type:     gonfc.ISO14443a,
		BaudRate: gonfc.Nbr106,
	}
	candidates, err := device.InitiatorListPassiveTargets(modulation)
	if err != nil {
		return tags, fmt.Errorf("list passive targets: %w", err)
	}
	for _, c := range candidates {
		t := freefareTagNew(device, c)
		if t != nil {
			tags = append(tags, t)
		}
	}

	// Poll for a FELICA tag
	modulation = gonfc.Modulation{
		Type:     gonfc.Felica,
		BaudRate: gonfc.Nbr424,
	}
	candidates, err = device.InitiatorListPassiveTargets(modulation)
	if err != nil {
		return nil, fmt.Errorf("list passive targets: %w", err)
	}
	for _, c := range candidates {
		t := freefareTagNew(device, c)
		tags = append(tags, t)
	}
	return tags, nil
}
